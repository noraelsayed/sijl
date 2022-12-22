package serv

import (
	"log"
	"net"

	"github.com/CSC354/sijl/internal/profile"
	"github.com/CSC354/sijl/internal/sijl"
	"github.com/CSC354/sijl/pkg/mamar"
	"github.com/CSC354/sijl/pkg/wathiq"
	"github.com/CSC354/sijl/psijl"
	"google.golang.org/grpc"
)

func StartSijlServer() error {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	sijlDb, err := mamar.ConnectDB("QAIDA")
	if err != nil {
		log.Fatal(err)
	}
	wathq, conn, err := wathiq.NewWathiqStub()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	sijl := &sijl.Sijl{DB: sijlDb, WathiqClient: wathq}
	psijl.RegisterSijlServer(grpcServer, sijl)

	profile := &profile.Profile{DB: sijlDb, WathiqClient: wathq}

	psijl.RegisterProfileServer(grpcServer, profile)

	err = grpcServer.Serve(lis)
	return err
}
