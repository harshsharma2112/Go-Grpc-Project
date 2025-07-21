package main

import (
	"context"
	"io"
	"log"

	pb "github.com/harshsharma2112/Go-Grpc-Project/proto"
)

func SayHelloServerStream(client pb.GreetServiceClient, names *pb.Nameslist) {
	log.Printf("streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names:%v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while receivng stream message:%v", err)
		}
		log.Println(message)
	}
	log.Printf("stream ended")
}
