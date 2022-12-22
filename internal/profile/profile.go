package profile

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	. "github.com/CSC354/sijl/perrors"
	"github.com/CSC354/sijl/pkg/stmts"
	. "github.com/CSC354/sijl/psijl"
	"github.com/CSC354/sijl/pwathiq"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Profile struct {
	UnimplementedProfileServer
	*sql.DB
	pwathiq.WathiqClient
}

// Get implements psijl.ProfileServer
func (p *Profile) Get(ctx context.Context, req *GetProfileRequest) (*ProfileResponse, error) {
	res := ProfileResponse{Details: &Details{}}
	checkStmt, err := p.DB.Prepare(stmts.CheckUser)
	if err != nil {
		log.Fatal(err)
	}
	defer checkStmt.Close()
	var c int
	err = checkStmt.QueryRow(sql.Named("username", req.Username)).Scan(&c)
	if err != nil {
		log.Fatal(err)
	}
	if c == 0 {
		res.Error = int32(Errors_NotFound)
		return &res, err
	}
	userStmt, err := p.DB.Prepare(`SELECT
  first_name,
  last_name,
  username,
  github,
  home,
  about,
  twitter,
  date_joined,
  age
FROM
  SIJL.USERS
WHERE
  username = @username
`)
	if err != nil {
		log.Fatal(err)
	}
	var date time.Time
	defer userStmt.Close()

	var (
		github  sql.NullString
		home    sql.NullString
		about   sql.NullString
		twitter sql.NullString
	)

	err = userStmt.QueryRow(sql.Named("username", req.Username)).
		Scan(&res.Details.FirstName, &res.Details.LastName, &res.Details.Username,
			&github, &home, &about, &twitter, &date, &res.Details.Age)
	if err != nil {
		log.Fatal(err)
	}
	res.Details.Github, res.Details.Home, res.Details.About, res.Details.Twitter =
		github.String, home.String, about.String, twitter.String

	res.Details.Joined = date.Unix()
	res.Details.Avatar, err = ioutil.ReadFile(fmt.Sprintf("/usr/src/sijl/imgs/%s.png", req.Username))
	if err != nil {
		log.Fatal(fmt.Sprintf("/usr/src/sijl/imgs/%s.png", req.Username), err)
	}

	return &res, err
}

// Update implements psijl.ProfileServer
func (p *Profile) Update(ctx context.Context, req *UpdateProfileRequest) (*ProfileResponse, error) {
	res := ProfileResponse{}
	a, err := p.WathiqClient.Validate(context.Background(), &pwathiq.ValidateRequest{Token: req.Token})
	if err != nil {
		log.Fatal(err)
	}

	if a.Id != req.Profile.Username {
		res.Error = int32(Errors_SomethingWrong)
		return &res, err
	}

	stmt, err := p.DB.Prepare(`UPDATE
  SOMEONE
SET
  SOMEONE.first_name = @firstname,
  SOMEONE.last_name = @lastname,
  SOMEONE.github = @github,
  SOMEONE.home = @home,
  SOMEONE.about = @about,
  SOMEONE.twitter = @twitter,
  SOMEONE.age = @age
FROM
  SIJL.USERS as SOMEONE
WHERE
  SOMEONE.username = 'saleh'
`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		sql.Named("firstname", req.Profile.FirstName),
		sql.Named("lastname", req.Profile.LastName),
		sql.Named("github", req.Profile.Github),
		sql.Named("home", req.Profile.Home),
		sql.Named("about", req.Profile.About),
		sql.Named("twitter", req.Profile.Twitter),
		sql.Named("age", req.Profile.Age),
	)

	if err != nil {
		log.Fatal(err)
	}
	res.Details = req.Profile
	res.Error = int32(Errors_Ok)
	return &res, err
}

func ValidateUpdateRequest(req *Details) Errors {
	if req.Age < 15 {
		return Errors_InvalidAge
	}
	err := validation.Validate(req.FirstName, validation.Required)
	if err != nil {
		return Errors_InvalidFirstName
	}

	err = validation.Validate(req.LastName, validation.Required)
	if err != nil {
		return Errors_InvalidLastName
	}
	return Errors_Ok
}
