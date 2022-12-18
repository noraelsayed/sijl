package sijl

import (
	"database/sql"
	"log"

	"github.com/CSC354/sijl/pkg/validator"
	proto "github.com/CSC354/sijl/psijl"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateRegister(req *proto.NewUserRequest, db *sql.DB) proto.Err {
	if !validator.IsValidUsername(req.Username) {
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

	if !validator.IsValidPassword(req.Password) {
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
