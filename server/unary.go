package main

import (
	"context"
	"go-grpc/proto"
	"log"
)

func (s *helloServer) SayHello(
	ctx context.Context, req *proto.NoParam) (*proto.HelloResponse, error) {
	log.Printf("handling request...")
	return &proto.HelloResponse{
		Message: "Hello",
	}, nil
}
