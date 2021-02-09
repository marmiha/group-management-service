package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"group-management-api/app"
	"group-management-api/app/config/impl"
	"group-management-api/app/container"
	"group-management-api/app/logger"
	"group-management-api/domain"
	"group-management-api/domain/payload"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var cont *container.Container

func TestMain(m *testing.M) {
	var err error
	cont, err = app.InitApp()

	if err != nil {
		logger.Log.WithField("err", err).Error("App initialization failed.")
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func TestAdapter(t *testing.T) {
	adapterImpl := cont.AppConfig.AdapterConfig.Impl
	switch adapterImpl {
	case impl.RestAdapter:
		// Currently only REST API is implemented. More switch cases can be if there are more adapter implementations.
		t.Log("Testing REST API adapter.")
		RestImplTests(t)
	default:
		t.Errorf("Unknown adapter implementation %v.", adapterImpl)
	}
}

func RestImplTests(t *testing.T) {
	tt := []SubTest {
		{"UserTests", UserTests},
		{"GroupTests", GroupTests},
		{"AuthorizationTests", AuthorizationTests},
	}

	for _, st := range tt {
		st.Run(t)
	}
}

func UserTests(t *testing.T) {

	// SubTests
	tt := []SubTest {
		{"CreateUser", userCreationTest},
		{"LoginUser", userLoginTest},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}

var userMain = &User{
	Info: payload.RegisterUserPayload{
		Email:    "user@email.com",
		Password: "user",
		Name:     "User",
	},
	Token: "",
}

var userDelete = &User{
	Info: payload.RegisterUserPayload{
		Email: "delete@email.com",
		Password: "delete",
		Name: "",
	},
	Token: "",
}

func userCreationTest(t *testing.T) {

	tt := []RestTestCase{
		NewRestBuilder("UserDoesNotExist").
			Path("/users/1").
			Get().
			ExpectCode(http.StatusNotFound).
			Build(),

		NewRestBuilder("RegisterUserWithoutEmailFail").
			Path("/users").
			Post().
			ExpectCode(http.StatusBadRequest).
			WithBody(payload.RegisterUserPayload{
				Email:    "",
				Name:     userMain.GetName(),
				Password: userMain.GetPassword(),
			}).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("RegisterUserWithInvalidEmailFail").
			Path("/users").
			Post().
			ExpectCode(http.StatusBadRequest).
			WithBody(payload.RegisterUserPayload{
				Email:    "xxxx",
				Name:     userMain.GetName(),
				Password: userMain.GetPassword(),
			}).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("RegisterUserWithInvalidPasswordFail").
			Path("/users").
			Post().
			ExpectCode(http.StatusBadRequest).
			WithBody(payload.RegisterUserPayload{
				Email:    userMain.GetEmail(),
				Name:     userMain.GetName(),
				Password: "",
			}).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("RegisterUserWithCapsEmailFail").
			Path("/users").
			Post().
			ExpectCode(http.StatusBadRequest).
			WithBody(payload.RegisterUserPayload{
				Email:    "EMAIL@gmail.com",
				Name:     userMain.GetName(),
				Password: userMain.GetPassword(),
			}).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("RegisterUserSuccess").
			Path("/users").
			Post().
			ExpectCode(http.StatusCreated).
			WithBody(payload.RegisterUserPayload{
				Email:    userMain.GetEmail(),
				Name:     userMain.GetName(),
				Password: userMain.GetPassword(),
			}).
			ExpectBody(map[string]interface{}{
				"token": isPresent,
				"user": map[string]interface{}{
					"email": userMain.GetEmail(),
					"name":  userMain.GetName(),
				},
			}).
			Build(),

		NewRestBuilder("RegisterUserWithoutNameSuccess").
			Path("/users").
			Post().
			ExpectCode(http.StatusCreated).
			WithBody(payload.RegisterUserPayload{
				Email:    userDelete.GetEmail(),
				Name:     "",
				Password: userMain.GetPassword(),
			}).
			ExpectBody(map[string]interface{}{
				"token": isPresent,
				"user": map[string]interface{}{
					"email": userDelete.GetEmail(),
					"name":  "",
				},
			}).
			Build(),

		NewRestBuilder("UserDoesExist").
			Path("/users/1").
			Get().
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"email": userMain.GetEmail(),
				"name":  userMain.GetName(),
			}).
			Build(),

		NewRestBuilder("RegisterUserWithSameEmailFails").
			Path("/users").
			Post().
			ExpectCode(http.StatusBadRequest).
			WithBody(payload.RegisterUserPayload{
				Email:    userMain.GetEmail(),
				Name:     userMain.GetName(),
				Password: userMain.GetPassword(),
			}).
			ExpectBody(map[string]interface{}{
				"err": domain.ErrUserWithEmailAlreadyExists.Error(),
			}).
			Build(),
	}

	runApiRestTestCases(tt, t)
}

func userLoginTest(t *testing.T) {
	tt := []RestTestCase{
		NewRestBuilder("WrongPasswordLoginFail").
			Path("/login").
			Post().
			ExpectCode(http.StatusBadRequest).
			WithBody(payload.CredentialsUserPayload{
				Email:    userMain.GetEmail(),
				Password: "wrongpassword",
			}).
			ExpectBody(map[string]interface{}{
				"err": nil,
			}).
			Build(),

		NewRestBuilder("RightPasswordTokenResponseSuccess").
			Path("/login").
			Post().
			ExpectCode(http.StatusOK).
			WithBody(payload.CredentialsUserPayload{
				Email:    userMain.GetEmail(),
				Password: userMain.GetPassword(),
			}).
			ExpectBody(map[string]interface{}{
				"token": saveAuthToken(userMain),
			}).
			Build(),

		NewRestBuilder("AccessCurrentUserProfileSuccess").
			Path("/users/current").
			Get().
			WithAuth(userMain).
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"name":  userMain.GetName(),
				"email": userMain.GetEmail(),
			}).
			Build(),

		NewRestBuilder("PasswordFieldNotInResponseBody").
			Path("/users/current").
			Get().
			WithAuth(userMain).
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"password":      isNotPresent,
				"password_hash": isNotPresent,
			}).
			Build(),
	}
	runApiRestTestCases(tt, t)
}

func GroupTests(t *testing.T) {

}

func AuthorizationTests(t *testing.T) {

}

// For executing multiple test cases on REST API endpoints.
func runApiRestTestCases(tc []RestTestCase, t *testing.T) {
	apiUrl := "/api/v1"
	for _, tc := range tc {
		t.Run(tc.Name, func(t *testing.T) {

			var err error
			var buf []byte = nil

			// If theres a request body.
			if tc.BodyReq != nil {
				// Marshall it.
				buf, err = json.Marshal(tc.BodyReq)
				if err != nil {
					t.Fatalf("Unable to marshall request body %v ", err)
				}
			}

			// Create the request.
			req, _ := http.NewRequest(tc.Method, apiUrl+tc.Url, bytes.NewBuffer(buf))

			// Set the request headers.
			if tc.Auth != nil {
				req.Header.Set("Authorization", "Bearer "+tc.Auth.GetToken())
			}

			// Execute the request.
			res := execRequest(req)

			// Compare the response code to the one that we expect.
			checkResponseCode(t, tc.RespCode, res.Code)

			// Do we have to check the values in body?
			if tc.BodyRes != nil {

				// Unmarshall the response body.
				var bodyRes map[string]interface{}

				if err = json.Unmarshal(res.Body.Bytes(), &bodyRes); err != nil {
					t.Errorf("Unable to unmarshall response body: %v", err)
				}

				// Check if the response body includes the wanted fields.
				firstIncludedInOther(tc.BodyRes, bodyRes, t)
			}
		})
	}
}

// Used for checking if information of the first is included in the second.
// Mainly used for comparing if the values of the API response bodies are what we expect.
// Can also check if the value is present when req["key"] == nil
// Can also pass values of res fields to functions when req["key"] == func(interface{}) error
// I also consider this function as my brain child 🐤.
func firstIncludedInOther(req map[string]interface{}, res map[string]interface{}, t *testing.T) {
	for key, element := range req {
		if val, ok := res[key]; ok == false {
			// If the key is not set.
			if element == nil {
				// If we just wanted it to be set.
				t.Errorf("Expected to have \"%v\" but it was unset in response.", key)
				continue
			}

			if fun, ok := element.(func(interface{}) error); ok {
				// Pass the nil variable to the function
				if err := fun(nil); err != nil {
					t.Errorf("Functional call of \"%v\" returned an error: %v", key, err)
				}
				continue
			}

			t.Errorf("Expected to have \"%v\"=%v but it was unset in response.", key, element)
		} else if _, isMap := element.(map[string]interface{}); isMap {
			// If the req element is a map.
			if _, isMap := val.(map[string]interface{}); isMap {
				// Recursive.
				firstIncludedInOther(element.(map[string]interface{}), val.(map[string]interface{}), t)
			}
		} else if fun, ok := element.(func(interface{}) error); ok {
			// Is element a function that accepts a string.
			if str, ok := val.(interface{}); ok == true && str != nil {
				// Pass the function the variable.
				err := fun(str)
				if err != nil {
					t.Errorf("Calling the function on \"%v\" returned an error: %v", key, err)
				}
			} else {
				t.Errorf("Expected to get value of \"%v\" but it was unset in response.", key)
			}
		} else if element == nil {
			// If its nil, then we just want to make sure it is set.
			continue
		} else if val != element {
			t.Errorf("Expected to have \"%v\"=%v but the response had: %v.", key, element, val)
		}
	}
}

// For executing requests on our Router.
func execRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	cont.Adapter.(*chi.Mux).ServeHTTP(rr, req)
	return rr
}

// For comparing HTTP response codes.
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

/******* Field Functions ********/

// Used for logging the fields of the responses.
func logField(key string, t *testing.T) func(interface{}) error {
	return func(val interface{}) error {
		t.Logf("Logged \"%v\" = %v", key, val)
		return nil
	}
}

func isPresent(val interface{}) error {
	if val == nil {
		return errors.New("field should be present")
	}
	return nil
}

func isNotPresent(val interface{}) error {
	if val != nil {
		return errors.New("field should not be present")
	}
	return nil
}

func saveAuthToken(u *User) func(interface{}) error {
	return func(str interface{}) error {
		tmpToken, ok := str.(string)
		if !ok || str == nil {
			return errors.New("token in response is not a string")
		}
		// Set the auth token to the value found in body.
		u.SetToken(tmpToken)
		return nil
	}
}

// Simple User Struct for easier testing.
type User struct {
	Info  payload.RegisterUserPayload
	Token string
}

func (u *User) GetName() string {
	return u.Info.Name
}

func (u *User) GetEmail() string {
	return u.Info.Email
}

func (u *User) GetPassword() string {
	return u.Info.Password
}

func (u *User) GetToken() string {
	return u.Token
}

func (u *User) SetName(name string) {
	u.Info.Name = name
}

func (u *User) SetEmail(email string) {
	u.Info.Email = email
}

func (u *User) SetPassword(pass string) {
	u.Info.Password = pass
}

func (u *User) SetToken(token string) {
	u.Token = token
}

// SubTest class.
type SubTest struct {
	Name string
	Fun func(t *testing.T)
}

func (st SubTest) Run(t *testing.T)  {
	t.Run(st.Name, func(t *testing.T) {
		st.Fun(t)
	})
}

// Simple Test Case For Rest.
type RestTestCase struct {
	Name     string
	Method   string
	Url      string
	Auth     *User
	RespCode int
	BodyReq  interface{}
	BodyRes  map[string]interface{}
}

// RestTestCaseBuilder for easier test reading and writing.
type RestTestCaseBuilder struct {
	tc RestTestCase
}

func NewRestBuilder(name string) *RestTestCaseBuilder {
	return &RestTestCaseBuilder{
		RestTestCase{
			Name:     name,
			Method:   http.MethodGet,
			Url:      "",
			Auth:     nil,
			RespCode: http.StatusOK,
			BodyReq:  nil,
			BodyRes:  nil,
		}}
}

func (r *RestTestCaseBuilder) Get() *RestTestCaseBuilder {
	r.tc.Method = http.MethodGet
	return r
}

func (r *RestTestCaseBuilder) Post() *RestTestCaseBuilder {
	r.tc.Method = http.MethodPost
	return r
}

func (r *RestTestCaseBuilder) Put() *RestTestCaseBuilder {
	r.tc.Method = http.MethodPut
	return r
}

func (r *RestTestCaseBuilder) Patch() *RestTestCaseBuilder {
	r.tc.Method = http.MethodPatch
	return r
}

func (r *RestTestCaseBuilder) Delete() *RestTestCaseBuilder {
	r.tc.Method = http.MethodDelete
	return r
}

func (r *RestTestCaseBuilder) ExpectCode(code int) *RestTestCaseBuilder {
	r.tc.RespCode = code
	return r
}

func (r *RestTestCaseBuilder) Path(path string) *RestTestCaseBuilder {
	r.tc.Url = path
	return r
}

func (r *RestTestCaseBuilder) WithBody(body interface{}) *RestTestCaseBuilder {
	r.tc.BodyReq = body
	return r
}

func (r *RestTestCaseBuilder) ExpectBody(body map[string]interface{}) *RestTestCaseBuilder {
	r.tc.BodyRes = body
	return r
}

func (r *RestTestCaseBuilder) WithAuth(user *User) *RestTestCaseBuilder {
	r.tc.Auth = user
	return r
}

func (r *RestTestCaseBuilder) Build() RestTestCase {
	return r.tc
}
