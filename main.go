package main

import (
	"log"
	"net"

	"github.com/atharv-bhadange/grpc-user-service/db"
	pb "github.com/atharv-bhadange/grpc-user-service/go-proto"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	userServiceServer := db.GetUserServiceServer()

	pb.RegisterUserServiceServer(grpcServer, userServiceServer)

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server listening on port 8080")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
