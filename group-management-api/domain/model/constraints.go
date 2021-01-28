package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Model fields constraints.
var (
	GroupNameMinLength = 2
	GroupNameMaxLength = 10
	UserNameMinLength = 3
	UserNameMaxLength = 12
)

/* Custom common rules for payload and entities fields verification. */

// Group.Name is required.
var GroupNameRule = []validation.Rule {
	validation.Required,
	validation.Length(GroupNameMinLength, GroupNameMaxLength),
}

// User.Name is not required.
var UserNameRule = []validation.Rule {
	validation.Length(UserNameMinLength, UserNameMaxLength),
}

// User.Email is required.
var UserEmailRule = []validation.Rule {
	validation.Required,
	// Change this to is.Email if you want to check MX records for domain validation.
	is.EmailFormat,
}
