package main

import (
	"context"
	pb "github.com/LittleMikle/testGRPC/proto"
	"io"
	"log"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("failed to send names: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed while streaming: %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
