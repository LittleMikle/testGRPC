package main

import (
	"context"
	pb "github.com/LittleMikle/testGRPC/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("failed with say")
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("failed while sending %v", err)
		}
		log.Printf("Sent the req with name: %s", name)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finish")
	if err != nil {
		log.Fatalf("failed while receiving: %v", err)
	}
	log.Printf("%v", res.Messages)
}
