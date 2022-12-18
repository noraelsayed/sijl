package sijl

import (
	"context"
	"database/sql"
	"log"
	"os/exec"

	"github.com/CSC354/sijl/psijl"

	"github.com/CSC354/sijl/pkg/stmts"
	"github.com/CSC354/sijl/pwathiq"
	_ "github.com/microsoft/go-mssqldb"
)

type Sijl struct {
	psijl.UnimplementedSijlServer
	*sql.DB
	pwathiq.WathiqClient
}

// Login implements psijl.SijlServer
func (s *Sijl) Login(ctx context.Context, req *psijl.LoginRequest) (*psijl.LoginResponse, error) {
	res := psijl.LoginResponse{}

	userStmt, err := s.DB.Prepare(stmts.CheckUser)
	if err != nil {
		log.Fatal(err)
	}
	defer userStmt.Close()
	var c int
	err = userStmt.QueryRow(sql.Named("username", req.Username)).Scan(&c)
	if err != nil {
		log.Fatal(err)
	}
	if c == 0 {
		res.Error = psijl.Err_WrongUsername
		return &res, err
	}

	passStmt, err := s.DB.Prepare(`
SELECT COUNT(*)
FROM SIJL.USERS
WHERE username = @username AND hash = HASHBYTES('SHA2_512', @password)
`)
	if err != nil {
		log.Fatal(err)
	}
	defer passStmt.Close()

	var valid int
	err = passStmt.QueryRow(sql.Named("username", req.Username), sql.Named("password", req.Password)).Scan(&valid)
	if err != nil {
		log.Fatal(err)
	}
	// generate token
	if valid == 1 {
		tkn, err := s.WathiqClient.GetToken(ctx, &pwathiq.TokenRequest{Username: req.Username})
		if err != nil {
			log.Fatal(err)
		}
		res.Token = tkn.Token
		return &res, nil
	}
	res.Error = psijl.Err_WrongPassword
	return &res, err

}

// Register implements psijl.SijlServer
func (s *Sijl) Register(ctx context.Context, req *psijl.NewUserRequest) (*psijl.LoginResponse, error) {
	res := psijl.LoginResponse{}
	res.Error = ValidateRegister(req, s.DB)
	if res.Error != psijl.Err_Ok {
		return &res, nil
	}
	stmt, err := s.DB.Prepare(`INSERT INTO SIJL.USERS(hash, first_name, last_name,
email,username,age) VALUES (HashBytes('SHA2_512', @hash ), @first_name , @last_name,
@email , @username , @age )`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	q, err := stmt.Exec(sql.Named("hash", req.Password), sql.Named("first_name", req.FirstName),
		sql.Named("last_name", req.LastName), sql.Named("email", req.Email),
		sql.Named("username", req.Username), sql.Named("age", req.Age))
	if err != nil {
		log.Println(q)
		log.Fatal(err)
	}
	tkn, err := s.WathiqClient.GetToken(ctx, &pwathiq.TokenRequest{Username: req.Username})
	if err != nil {
		log.Fatal(err)
	}
	res.Token = tkn.Token

	cmd := exec.Command("python3", "/usr/src/sijl/imgs/gen.py", req.Username)
	cmd.Dir = "/usr/src/sijl/imgs/"
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return &res, err
}

// TODO Handle used emails
// TODO organize prepared statements
// TODO is it a good idea to generate hashes in the database side instead of here?
// TODO Replace fatal errors
