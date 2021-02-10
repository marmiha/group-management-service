package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	"strconv"
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

/* REST ADAPTER IMPLEMENTATION TESTING */
func RestImplTests(t *testing.T) {
	tt := []SubTest{
		{"User", UserTests},
		{"Group", GroupTests},
		{"Authorization", AuthorizationTests},
	}

	for _, st := range tt {
		st.Run(t)
	}
}

/* User Tests */
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
		Email:    "delete@email.com",
		Password: "delete",
		Name:     "",
	},
	Token: "",
}

func UserTests(t *testing.T) {
	tt := []SubTest{
		{"Register", userCreationTest},
		{"Login", userLoginTest},
		{"Modify", userModifyTest},
		{"Unregister", userUnregisterTest},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}

func userCreationTest(t *testing.T) {

	tt := []RestTestCase{
		NewRestBuilder("UserDoesNotExist").
			Path("/users/1").
			Get().
			ExpectCode(http.StatusNotFound).
			Build(),

		NewRestBuilder("WithoutEmailFails").
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

		NewRestBuilder("WithInvalidEmailFails").
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

		NewRestBuilder("WithInvalidPasswordFails").
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

		NewRestBuilder("WithCapsEmailFails").
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

		NewRestBuilder("CorrectPasswordAndEmailSuccess").
			Path("/users").
			Post().
			ExpectCode(http.StatusCreated).
			WithBody(payload.RegisterUserPayload{
				Email:    userMain.GetEmail(),
				Name:     userMain.GetName(),
				Password: userMain.GetPassword(),
			}).
			ExpectBody(map[string]interface{}{
				"token": saveAuthToken(userMain),
				"user": map[string]interface{}{
					"email": userMain.GetEmail(),
					"name":  userMain.GetName(),
					"id":    setID(userMain),
				},
			}).
			Build(),

		NewRestBuilder("WithoutNameSuccess").
			Path("/users").
			Post().
			ExpectCode(http.StatusCreated).
			WithBody(payload.RegisterUserPayload{
				Email:    userDelete.GetEmail(),
				Name:     "",
				Password: userDelete.GetPassword(),
			}).
			ExpectBody(map[string]interface{}{
				"token": saveAuthToken(userDelete),
				"user": map[string]interface{}{
					"email": userDelete.GetEmail(),
					"name":  "",
					"id":    setID(userDelete),
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

		NewRestBuilder("WithTakenEmailFails").
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

func userUnregisterTest(t *testing.T) {
	tt := []RestTestCase{
		NewRestBuilder("DeleteUserWrongPasswordFail").
			Path("/users/current").
			Delete().
			WithAuth(userDelete).
			WithBody(payload.UnregisterUserPayload{
				Password: "wrongpassword",
			}).
			ExpectCode(http.StatusBadRequest).
			Build(),

		NewRestBuilder("DeleteUserSuccess").
			Path("/users/current").
			Delete().
			ExpectCode(http.StatusNoContent).
			WithAuth(userDelete).
			WithBody(payload.UnregisterUserPayload{
				Password: userDelete.GetPassword(),
			}).
			Build(),

		NewRestBuilder("DeletedUserNotFoundAnymore").
			Path("/users/" + strconv.Itoa(userDelete.GetID())).
			Get().
			ExpectCode(http.StatusNotFound).
			Build(),
	}
	runApiRestTestCases(tt, t)
}

func userModifyTest(t *testing.T) {
	firstName := "first new name"
	secondName := "second new name"
	firstEmail := "nottaken@first.com"
	secondEmail := "nottaken@second.com"
	newPassword := "new_password"

	tt := []RestTestCase{
		NewRestBuilder("WithEmptyPayload").
			Path("/users/current").
			Patch().
			WithAuth(userMain).
			WithBody(payload.ModifyUserPayload{}).
			ExpectCode(http.StatusBadRequest).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("OnlyNameSuccess").
			Path("/users/current").
			Patch().
			WithAuth(userMain).
			WithBody(payload.ModifyUserPayload{
				Name: firstName,
			}).
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"name": firstName,
			}).
			After(func() {
				userMain.SetName(firstName)
			}).
			Build(),

		NewRestBuilder("TakenEmailFails").
			Path("/users/current").
			Patch().
			WithAuth(userMain).
			WithBody(payload.ModifyUserPayload{
				Email: userMain.GetEmail(),
			}).
			ExpectCode(http.StatusBadRequest).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("EmailWithCapsFails").
			Path("/users/current").
			Patch().
			WithAuth(userMain).
			WithBody(payload.ModifyUserPayload{
				Email: "eMaIl@CaPs.com",
			}).
			ExpectCode(http.StatusBadRequest).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("NotTakenEmailSuccess").
			Path("/users/current").
			Patch().
			WithAuth(userMain).
			WithBody(payload.ModifyUserPayload{
				Email: firstEmail,
			}).
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"email": firstEmail,
			}).
			After(func() {
				userMain.SetEmail(firstEmail)
			}).
			Build(),

		NewRestBuilder("NotTakenEmailAndNameSuccess").
			Path("/users/current").
			Patch().
			WithAuth(userMain).
			WithBody(payload.ModifyUserPayload{
				Name:  secondName,
				Email: secondEmail,
			}).
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"name":  secondName,
				"email": secondEmail,
			}).
			After(func() {
				userMain.SetEmail(secondEmail)
				userMain.SetName(secondName)
			}).
			Build(),

		NewRestBuilder("ChangePasswordWrongCurrentPasswordFails").
			Path("/users/current/attributes/password").
			Put().
			WithAuth(userMain).
			WithBody(payload.ChangePasswordPayload{
				CurrentPassword: "wrong_password",
				NewPassword:     newPassword,
			}).
			ExpectCode(http.StatusBadRequest).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("PasswordDidNotChange").
			Path("/login").
			Post().
			ExpectCode(http.StatusBadRequest).
			WithBody(payload.CredentialsUserPayload{
				Email:    secondEmail,
				Password: newPassword,
			}).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("ChangePasswordEmptyNewFail").
			Path("/users/current/attributes/password").
			Put().
			WithAuth(userMain).
			WithBody(payload.ChangePasswordPayload{
				CurrentPassword: userMain.GetPassword(),
				NewPassword:     "f",
			}).
			ExpectCode(http.StatusBadRequest).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("PasswordDidNotChange").
			Path("/login").
			Post().
			ExpectCode(http.StatusBadRequest).
			WithBody(payload.CredentialsUserPayload{
				Email:    secondEmail,
				Password: newPassword,
			}).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("PasswordChangeSuccess").
			Path("/users/current/attributes/password").
			Put().
			WithAuth(userMain).
			WithBody(payload.ChangePasswordPayload{
				CurrentPassword: userMain.GetPassword(),
				NewPassword:     newPassword,
			}).
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"id": float64(userMain.GetID()),
			}).
			After(func() {
				userMain.SetPassword(newPassword)
			}).
			Build(),

		NewRestBuilder("PasswordDidChange").
			Path("/login").
			Post().
			ExpectCode(http.StatusOK).
			WithBody(payload.CredentialsUserPayload{
				Email:    secondEmail,
				Password: newPassword,
			}).
			ExpectBody(map[string]interface{}{
				"token": isPresent,
			}).
			Build(),
	}
	runApiRestTestCases(tt, t)
}

/* Group Tests */
var groupMain = &Group{
	Info: payload.CreateGroupPayload{
		Name: "main",
	},
}

var groupDelete = &Group{
	Info: payload.CreateGroupPayload{
		Name: "delete",
	},
}

func GroupTests(t *testing.T) {
	tt := []SubTest{
		{"Create", groupCreationTest},
		{"Modify", groupModifyTest},
		{"Delete", groupDeleteTest},
	}

	for _, tc := range tt {
		tc.Run(t)
	}
}

func groupCreationTest(t *testing.T) {
	tt := []RestTestCase{
		NewRestBuilder("NameCapsFails").
			Path("/groups").
			Post().
			WithBody(payload.CreateGroupPayload{
				Name: "cApS",
			}).
			ExpectCode(http.StatusBadRequest).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("GroupNotCreated").
			Path("/groups/1").
			Get().
			ExpectCode(http.StatusNotFound).
			Build(),

		NewRestBuilder("CorrectNameSuccess").
			Path("/groups").
			Post().
			WithBody(payload.CreateGroupPayload{
				Name: groupMain.GetName(),
			}).
			ExpectCode(http.StatusCreated).
			ExpectBody(map[string]interface{}{
				"id":   setIDgroup(groupMain),
				"name": groupMain.GetName(),
			}).
			Build(),

		NewRestBuilder("GroupCreatedCheck").
			Path("/groups/1").
			Get().
			ExpectCode(http.StatusOK).
			Build(),

		NewRestBuilder("TakenNameFails").
			Path("/groups").
			Post().
			WithBody(payload.CreateGroupPayload{
				Name: groupMain.GetName(),
			}).
			ExpectCode(http.StatusBadRequest).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("GroupNotCreatedCheck").
			Path("/groups/2").
			Get().
			ExpectCode(http.StatusNotFound).
			Build(),

		NewRestBuilder("NotTakenNameSuccess").
			Path("/groups").
			Post().
			WithBody(payload.CreateGroupPayload{
				Name: groupDelete.GetName(),
			}).
			ExpectCode(http.StatusCreated).
			ExpectBody(map[string]interface{}{
				"id":   setIDgroup(groupDelete),
				"name": groupDelete.GetName(),
			}).
			Build(),

		NewRestBuilder("TwoGroupsCreatedCheck").
			Path("/groups/2").
			Get().
			ExpectCode(http.StatusOK).
			Build(),
	}
	runApiRestTestCases(tt, t)
}

func groupModifyTest(t *testing.T) {
	newName := "new group name"

	tt := []RestTestCase{
		NewRestBuilder("NameCapsFails").
			Path(fmt.Sprintf("/groups/%v", groupMain.GetID())).
			Patch().
			WithBody(payload.ModifyGroupPayload{
				Name: "cApS",
			}).
			ExpectCode(http.StatusBadRequest).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("NameDidNotChange").
			Path(fmt.Sprintf("/groups/%v", groupMain.GetID())).
			Get().
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"id": float64(groupMain.ID),
				"name": groupMain.GetName(),
			}).
			Build(),

		NewRestBuilder("NameTakenFails").
			Path(fmt.Sprintf("/groups/%v", groupMain.GetID())).
			Patch().
			WithBody(payload.ModifyGroupPayload{
				Name: groupDelete.GetName(),
			}).
			ExpectCode(http.StatusBadRequest).
			ExpectBody(map[string]interface{}{
				"err": isPresent,
			}).
			Build(),

		NewRestBuilder("NameTakenDidNotChange").
			Path(fmt.Sprintf("/groups/%v", groupMain.GetID())).
			Get().
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"id": float64(groupMain.ID),
				"name": groupMain.GetName(),
			}).
			Build(),

		NewRestBuilder("NotTakenAndCorrectNameSuccess").
			Path(fmt.Sprintf("/groups/%v", groupMain.GetID())).
			Patch().
			WithBody(payload.ModifyGroupPayload{
				Name: newName,
			}).
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"id": float64(groupMain.ID),
				"name": newName,
			}).
			After(func() {
				groupMain.SetName(newName)
			}).
			Build(),

		NewRestBuilder("NameDidChange").
			Path(fmt.Sprintf("/groups/%v", groupMain.GetID())).
			Get().
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"id": float64(groupMain.ID),
				"name": newName,
			}).
			Build(),
	}
	runApiRestTestCases(tt, t)
}

func groupDeleteTest(t *testing.T) {
	tt := []RestTestCase{
		NewRestBuilder("GroupExists").
			Path(fmt.Sprintf("/groups/%v", groupDelete.GetID())).
			Get().
			ExpectCode(http.StatusOK).
			ExpectBody(map[string]interface{}{
				"id": float64(groupDelete.GetID()),
			}).
			Build(),

		NewRestBuilder("SuccessfulDelete").
			Path(fmt.Sprintf("/groups/%v", groupDelete.GetID())).
			Delete().
			ExpectCode(http.StatusNoContent).
			Build(),

		NewRestBuilder("GroupDoesNotExist").
			Path(fmt.Sprintf("/groups/%v", groupDelete.GetID())).
			Get().
			ExpectCode(http.StatusNotFound).
			Build(),
	}
	runApiRestTestCases(tt, t)
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

			// Callbacks based on success
			if t.Failed() {
				if tc.FFail != nil {
					tc.FFail()
				}
			} else {
				if tc.FSucc != nil {
					tc.FSucc()
				}
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

func setID(user *User) func(interface{}) error {
	return func(val interface{}) error {
		id, ok := val.(float64)
		if !ok {
			return errors.New("value is not a number")
		}
		user.SetID(int(id))
		return nil
	}
}

func setIDgroup(group *Group) func(interface{}) error {
	return func(val interface{}) error {
		id, ok := val.(float64)
		if !ok {
			return errors.New("value is not a number")
		}
		group.SetID(int(id))
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

// Simple Group Struct for easier testing.
type Group struct {
	ID   int
	Info payload.CreateGroupPayload
}

func (g *Group) GetID() int {
	return g.ID
}

func (g *Group) SetName(name string) {
	g.Info.Name = name
}

func (g *Group) GetName() string {
	return g.Info.Name
}

func (g *Group) SetID(id int) {
	g.ID = id
}

// Simple User Struct for easier testing.
type User struct {
	ID    int
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

func (u *User) GetID() int {
	return u.ID
}

func (u *User) SetID(id int) {
	u.ID = id
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
	Fun  func(t *testing.T)
}

func (st SubTest) Run(t *testing.T) {
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
	FFail    func()
	FSucc    func()
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
			FSucc:    nil,
			FFail:    nil,
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

func (r *RestTestCaseBuilder) After(fun func()) *RestTestCaseBuilder {
	r.tc.FSucc = fun
	return r
}

func (r *RestTestCaseBuilder) Catch(fun func()) *RestTestCaseBuilder {
	r.tc.FSucc = fun
	return r
}

func (r *RestTestCaseBuilder) Build() RestTestCase {
	return r.tc
}
