package main

import (
	"context"
	"go-grpc/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBiDirectionalStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Println("Bidirectional streaming started...")

	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names %v", err)
	}

	// Receive
	waitChanel := make(chan struct{})
	go func() {
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
		close(waitChanel)
	}()

	// Send
	for _, name := range names.Names {
		req := &proto.HelloRequest{Name: name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitChanel
	log.Printf("Streaming finished")
}
