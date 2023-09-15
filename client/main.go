package main

import (
	pb "github.com/LittleMikle/testGRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{
			"Mikle",
			"Jopa",
			"Negra",
		},
	}

	//callSayHello(client)
	//callSayHelloServerStreaming(client, names)
	//callSayHelloClientStream(client, names)
	callHelloBidirectionalStream(client, names)
}
