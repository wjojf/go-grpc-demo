package main

import (
	"fmt"
	"go-grpc/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream proto.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&proto.MessagesList{Messages: messages})
			}
			return err
		}
		log.Printf("Received request with name: %v", req.Name)
		messages = append(messages, fmt.Sprintf("Hello, %v", req.Name))
	}
}
