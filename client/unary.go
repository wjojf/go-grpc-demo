package main

import (
	"context"
	"go-grpc/proto"
	"log"
	"time"
)

func callSeyHello(client proto.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &proto.NoParam{})
	if err != nil {
		log.Fatalf("could not greet %v", err)
	}
	log.Printf("received: %v", res.Message)
}
