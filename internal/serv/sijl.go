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
func (s *Sijl) Login(ctx context.Context, req *psijl.LoginRequest) (res *psijl.LoginResponse, err error) {
	stmt, err := s.DB.Prepare(`IF HASHBYTES('SHA2_512', ?) = (
SELECT hash
FROM SIJL.USERS
WHERE username = ?
)
BEGIN
   Select 1
END
ELSE
BEGIN
   Select 0
END
`)
	var valid int
	err = stmt.QueryRow(req.Username, req.Password).Scan(&valid)
	if err != nil {
		log.Fatal(err)
	}
	if valid == 1 {
		tkn, err := s.WathiqClient.GetToken(ctx, &pwathiq.TokenRequest{Username: req.Username})
		if err != nil {
			log.Fatal(err)
		}
		res.Token = tkn.Token
		return res, nil
	}
	res.Error = psijl.Err_WrongPassword
	return

}

// Register implements psijl.SijlServer
func (s *Sijl) Register(ctx context.Context, req *psijl.NewUserRequest) (res *psijl.LoginResponse, err error) {
	res.Error = ValidateRegister(req, s.DB)
	if res.Error != psijl.Err_Ok {
		return
	}
	stmt, err := s.DB.Prepare(`INSERT INTO SIJL.USERS(hash, first_name, last_name, email,username,age)
VALUES (HashBytes('SHA2_512', ?),?,?,?,?,?)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(req.Password, req.FirstName, req.LastName, req.Email, req.Username, req.Age)
	if err != nil {
		log.Fatal(err)
	}
	tkn, err := s.WathiqClient.GetToken(ctx, &pwathiq.TokenRequest{Username: req.Username})
	if err != nil {
		log.Fatal(err)
	}
	res.Token = tkn.Token
	return
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