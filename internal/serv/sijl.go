package serv

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/CSC354/sijl/pkg/mamar"
	"github.com/CSC354/sijl/pkg/wathiq"
	"github.com/CSC354/sijl/psijl"

	"github.com/CSC354/sijl/pwathiq"
	_ "github.com/microsoft/go-mssqldb"
	"google.golang.org/grpc"
)

type Sijl struct {
	psijl.UnimplementedSijlServer
	*sql.DB
	pwathiq.WathiqClient
}

// Login implements psijl.SijlServer
func (s *Sijl) Login(ctx context.Context, req *psijl.LoginRequest) (*psijl.LoginResponse, error) {
	res := psijl.LoginResponse{}

	userStmt, err := s.DB.Prepare(`SELECT COUNT(*)
FROM SIJL.USERS
WHERE username = @username`)
	if err != nil {
		log.Fatal(err)
	}
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
	return &res, err
}

func StartSijlServer() error {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	qaida, err := mamar.ConnectDB("SIJL")
	if err != nil {
		log.Fatal(err)
	}
	wathq, conn, err := wathiq.NewWathiqStub()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	sijl := &Sijl{DB: qaida, WathiqClient: wathq}
	psijl.RegisterSijlServer(grpcServer, sijl)
	err = grpcServer.Serve(lis)
	return err
}

// TODO organize prepared statements
// TODO is it a good idea to generate hashes in the database side instead of here?
