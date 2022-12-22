package sijl

import (
	"database/sql"
	"log"

	. "github.com/CSC354/sijl/perrors"
	"github.com/CSC354/sijl/pkg/validator"
	proto "github.com/CSC354/sijl/psijl"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateRegister(req *proto.NewUserRequest, db *sql.DB) Errors {
	if !validator.IsValidUsername(req.Username) {
		return Errors_InvalidUsername
	}

	stmt, err := db.Prepare("SELECT COUNT(*) FROM SIJL.USERS WHERE username = @username")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var c int
	err = stmt.QueryRow(sql.Named("username", req.Username)).Scan(&c)

	if c > 0 {
		return Errors_AlreadyUsedUsername
	}

	c = 0
	emailstmt, err := db.Prepare("SELECT COUNT(*) FROM SIJL.USERS WHERE email = @email")
	if err != nil {
		log.Fatal(err)
	}
	defer emailstmt.Close()

	err = emailstmt.QueryRow(sql.Named("email", req.Email)).Scan(&c)

	if c > 0 {
		return Errors_AlreadyUsedEmail
	}

	if !validator.IsValidPassword(req.Password) {
		return Errors_InvalidPassword
	}

	err = validation.Validate(req.FirstName, validation.Required)
	if err != nil {
		return Errors_InvalidFirstName
	}

	err = validation.Validate(req.LastName, validation.Required)
	if err != nil {
		return Errors_InvalidLastName
	}

	err = validation.Validate(req.Email, validation.Required, is.Email)
	if err != nil {
		return Errors_InvalidEmail
	}
	if req.Age <= 15 {
		return Errors_InvalidAge
	}

	return Errors_Ok
}
