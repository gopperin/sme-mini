package grpc

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	"gateway/config"
)

// ClientConn ClientConn
var ClientConn *grpc.ClientConn

// Init Init
func Init() error {
	_conn, err := grpc.Dial(config.Server.GrpcURI, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	ClientConn = _conn
	return nil
}

// Close Close
func Close() error {
	err := ClientConn.Close()
	if err != nil {
		fmt.Println("grpc close error", err.Error())
		return err
	}
	fmt.Println("grpc close")
	return nil
}
