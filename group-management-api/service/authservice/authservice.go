// Implements the common interface the domain implementation will use for getting access tokens and handling the user
// password credentials. Different implementations will handle password hashing and token generation differently.
package authservice

import "group-management-api/domain/model"

type AuthServiceInterface interface {
	// Access token generation.
	GenerateAccessToken(user *model.User) (string, error)
	// Parse token from request.
	ParseToken()
}
