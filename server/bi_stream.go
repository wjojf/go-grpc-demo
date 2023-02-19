package main

import (
	"fmt"
	"go-grpc/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloBiDirectionalStreaming(
	stream proto.GreetService_SayHelloBiDirectionalStreamingServer) error {

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		log.Printf("Received request with name %v", req.Name)
		res := &proto.HelloResponse{
			Message: fmt.Sprintf("Hello, %v", req.Name),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
