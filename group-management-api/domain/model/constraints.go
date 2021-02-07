package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Model fields constraints. These are also used for payload verification.
var (
	GroupNameMinLength    = 3
	GroupNameMaxLength    = 40

	UserNameMinLength     = 3
	UserNameMaxLength     = 40

	UserPasswordMinLength = 4
	UserPasswordMaxLength = 120
)

/* Custom common rules for payload and entities fields verification. */

var GroupNameRule = []validation.Rule{
	is.LowerCase,
	validation.Length(GroupNameMinLength, GroupNameMaxLength),
}
var GroupNameRequiredRule = append(GroupNameRule, validation.Required)

var UserNameRule = []validation.Rule{
	validation.Length(UserNameMinLength, UserNameMaxLength),
}
var UserNameRequiredRule = append(UserNameRule, validation.Required)

var UserEmailRule = []validation.Rule{
	// Change this to is.Email if you want to check MX records for domain validation.
	is.LowerCase,
	is.EmailFormat,
}
var UserEmailRequiredRule = append(UserEmailRule, validation.Required)

// User password length rule (this is not for Hashed password), it's meant for authentication payloads verification.
var UserPasswordRule = []validation.Rule{
	validation.Required,
	validation.Length(UserPasswordMinLength, UserPasswordMaxLength),
}
