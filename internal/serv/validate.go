package serv

import (
	"database/sql"
	"log"
	"regexp"
	"unicode"

	proto "github.com/CSC354/sijl/psijl"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateRegister(req *proto.NewUserRequest, db *sql.DB) proto.Err {
	usernameConvention := "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"
	re := regexp.MustCompile(usernameConvention)

	err := validation.Validate(req.Username, validation.Required, validation.Match(re))
	if err != nil {
		return proto.Err_InvalidUsername
	}

	stmt, err := db.Prepare("SELECT COUNT(*) FROM SIJL.USERS WHERE username = @username")
	if err != nil {
		log.Fatal(err)
	}
	var c int
	err = stmt.QueryRow(sql.Named("username", req.Username)).Scan(&c)

	if c > 0 {
		return proto.Err_AlreadyUsedUsername
	}

	if !isValidPassword(req.Password) {
		return proto.Err_InvalidPassword
	}

	err = validation.Validate(req.FirstName, validation.Required)
	if err != nil {
		return proto.Err_InvalidFirstName
	}

	err = validation.Validate(req.LastName, validation.Required)
	if err != nil {
		return proto.Err_InvalidLastName
	}

	err = validation.Validate(req.Email, validation.Required, is.Email)
	if err != nil {
		return proto.Err_InvalidEmail
	}
	if req.Age <= 15 {
		return proto.Err_InvalidAge
	}

	return proto.Err_Ok
}

func isValidPassword(s string) bool {
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
