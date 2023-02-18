package main

import (
	"context"
	"go-grpc/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Printf("streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names %v", err)
	}

	for {
		message, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}

	log.Println("streaming finished")
}
