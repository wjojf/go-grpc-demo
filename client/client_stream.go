package main

import (
	"context"
	"go-grpc/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &proto.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response %v", err)
	}
	log.Printf("Received messages: %v", res.Messages)
}
