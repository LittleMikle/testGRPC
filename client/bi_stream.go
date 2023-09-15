package main

import (
	"context"
	pb "github.com/LittleMikle/testGRPC/proto"
	"io"
	"log"
	"time"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("failed with say bidirectional")
	}

	waitCH := make(chan struct{})

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("failed with streaming: %v", err)
			}
			log.Println(msg)
		}
		close(waitCH)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("failed to send: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitCH
	log.Printf("BIDIRECTIONAL STREAMING FINISH")
}
