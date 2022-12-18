package validator

import (
	"regexp"
	"unicode"

	validation "github.com/go-ozzo/ozzo-validation"
)

const UsernameConvention string = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"

func IsValidUsername(username string) bool {
	re := regexp.MustCompile(UsernameConvention)

	err := validation.Validate(username, validation.Required, validation.Match(re))
	if err != nil {
		return false
	}
	return true
}

func IsValidPassword(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
