package main

import (
	"log"
	"net"

	pb "github.com/gkpodder/grpc-go/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloserver struct {
	pb.GreetServiceServer
}

func main() {
	// Start the server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloserver{})
	log.Printf("Server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
