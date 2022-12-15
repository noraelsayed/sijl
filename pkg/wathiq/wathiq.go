package wathiq

import (
	"log"

	"github.com/CSC354/sijl/pwathiq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewWathiqStub() (stub pwathiq.WathiqClient, conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial("wathiq:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print(err)
		return
	}
	stub = pwathiq.NewWathiqClient(conn)
	return
}
