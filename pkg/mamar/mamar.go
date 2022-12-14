package mamar

import (
	"context"
	"database/sql"
	proto "github.com/CSC354/sijl/pmamar"
	_ "github.com/microsoft/go-mssqldb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewMamarStub() (stub proto.MamarClient, conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial("mamar:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	stub = proto.NewMamarClient(conn)
	return
}

// MamarGetPort creates a new Mamar instance, get a port and then close the connection
func MamarGetPort(service string) (port *proto.Port, err error) {
	mamr, conn, err := NewMamarStub()
	if err != nil {
		return
	}
	defer conn.Close()
	port, err = mamr.GetPort(context.Background(), &proto.Service{Name: "sijl_db"})
	if err != nil {
		return
	}
	return
}

// ConnectDB returns a SQL connection instance to a given db from Mamar
func ConnectDB(dbName string) (db *sql.DB, err error) {
	connectionStr, err := MamarGetPort(dbName)
	if err != nil {
		return
	}
	db, err = sql.Open("sqlserver", connectionStr.Address)
	if err != nil {
		return
	}
	err = db.PingContext(context.Background())
	if err != nil {
		return
	}
	return
}
