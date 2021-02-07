package restapi

import "group-management-api/domain/model"

// Error Response
// swagger:model ErrorResponse
type ErrorResponse struct {
	// the description of this error
	//
	// required: true
	// min: 1
	// example: ErrNotFound
	ErrorString string `json:"err"`
}

// Register Response
// swagger:model RegisterResponse
type RegisterResponse struct {
	User  model.User `json:"user"`

	// the jwt authentication token.
	//
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJzdGFuZGFyZF9jbGFpbXMiOnsiZXhwIjoxNjEyNTc3NTQxLCJqdGkiOiIxIiwiaWF0IjoxNjEyNTQxNTQxLCJpc3MiOiJHcm91cE1hbmFnZW1lbnRBcHAifX0.skb_BHRkLz86btb9JG20Xu7p9zDUhbqBLoZHIdM2PV0
	Token string     `json:"token"`
}

// Login Response
// swagger:model LoginResponse
type LoginResponse struct {
	// the jwt authentication token.
	//
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJzdGFuZGFyZF9jbGFpbXMiOnsiZXhwIjoxNjEyNTc3NTQxLCJqdGkiOiIxIiwiaWF0IjoxNjEyNTQxNTQxLCJpc3MiOiJHcm91cE1hbmFnZW1lbnRBcHAifX0.skb_BHRkLz86btb9JG20Xu7p9zDUhbqBLoZHIdM2PV0
	Token string `json:"token"`
}