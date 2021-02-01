package restapi

import (
	"context"
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"group-management-api/domain/model"
	"net/http"
)

// Universal payload tester for Http endpoints. Throws validation errors or decoding errors.
func validatePayload(next http.HandlerFunc, payload validation.Validatable) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// We will try to decode our request.body with our Validatable struct.
		decodingError := json.NewDecoder(request.Body).Decode(&payload)

		if decodingError != nil {
			// Call our defined handler function for Http BadRequest status code.
			badRequestResponse(writer, decodingError)
			return
		}

		defer request.Body.Close()

		// Validate the payload.
		if validationErrors := payload.Validate(); validationErrors != nil {
			// Bad request response.
			badRequestResponse(writer, validationErrors)
			return
		}

		// We will add a new filed to the context of our request.
		newPayloadContext := context.WithValue(request.Context(), "payload", payload)
		// Successful payload decoding, call the next handler.
		next.ServeHTTP(writer, request.WithContext(newPayloadContext))
	}
}

func notFoundResponse(writer http.ResponseWriter, error error) {
	response := map[string]string{
		"error": error.Error(),
	}
	jsonResponse(writer, response, http.StatusNotFound)
}

func createdResponse(writer http.ResponseWriter, responseData interface{}) {
	jsonResponse(writer, responseData, http.StatusCreated)
}

func okResponse(writer http.ResponseWriter, responseData interface{}) {
	jsonResponse(writer, responseData, http.StatusOK)
}

func badRequestResponse(writer http.ResponseWriter, error error) {
	response := map[string]string{
		"error": error.Error(),
	}
	jsonResponse(writer, response, http.StatusBadRequest)
}

func unauthorizedResponse(writer http.ResponseWriter, error error) {
	response := map[string]string{
		"error": error.Error(),
	}
	jsonResponse(writer, response, http.StatusUnauthorized)
}

func successfulDeleteResponse(writer http.ResponseWriter) {
	// Our response, what we reply back. Using the map[string]string we can
	// define json properties and their values.
	response := map[string]string{}
	jsonResponse(writer, response, http.StatusNoContent)
}

func internalServerErrorResponse(writer http.ResponseWriter, error error) {
	// Our response, what we reply back. Using the map[string]string we can define json properties and their values.
	response := map[string]string{
		"error": error.Error(),
	}
	jsonResponse(writer, response, http.StatusInternalServerError)
}


func jsonResponse(writer http.ResponseWriter, responseData interface{}, httpStatusCode int) {
	// Let's respond with a error corresponding to httpStatusCode.
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpStatusCode)

	// Our response, what we reply back. Using the map[string]string we can
	// define json properties and their values.
	// If data is empty, then send an empty object.
	if responseData == nil {
		responseData = map[string]string{}
	}

	// Encode the response and if something goes wrong with the encoding then
	// return a internal server error response.
	if encodingError := json.NewEncoder(writer).Encode(responseData); encodingError != nil {
		internalServerErrorResponse(writer, encodingError)
		return
	}
	return
}

func currentUserFromCtx(r *http.Request) *model.User {
	return r.Context().Value(contextCurrentUserKey).(*model.User)
}

func userFromCtx(r *http.Request) *model.User {
	return r.Context().Value(contextUserKey).(*model.User)
}

func groupFromCtx(r *http.Request) *model.Group {
	return r.Context().Value(contextGroupKey).(*model.Group)
}
