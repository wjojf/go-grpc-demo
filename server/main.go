package main

import (
	"go-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

type helloServer struct {
	proto.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	
	grpcServer := grpc.NewServer()
	proto.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("server started at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
