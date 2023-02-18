package main

import (
	"fmt"
	"go-grpc/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(
	in *proto.NamesList, stream proto.GreetService_SayHelloServerStreamingServer) error {

	log.Printf("received names: %v", in.Names)
	for _, name := range in.Names {
		res := &proto.HelloResponse{
			Message: fmt.Sprintf("Hello, %v", name),
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}
