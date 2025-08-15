package main

import (
	"context"
	"log"
	"time"

	pb "github.com/harshsharma2112/Go-Grpc-Project/proto"
)

// // . is the pointer to the struct but not directly it is baiscally in laymen term is pointing to the methods of interface
// func callSayHelloClientStream(stream pb.GreetServiceClient, names *pb.Nameslist) {
// 	log.Printf("streaming strarted")
// 	stream, err := client.callSayHelloClientStream(context.Background())
// 	if err != nil {
// 		log.Fatalf("could not send names:%v", err)
// 	}

// 	for _, names := range names.Names {
// 		req := &pb.HelloRequest{
// 			Name: name,
// 		}
// 		if err := stream.Send(req); err != nil {
// 			log.Fatalf("error while sending %v", err)
// 		}
// 		log.Printf("send the request with name %s", name)
// 		time.Sleep(2 * time.Second)
// 	}

// 	res, err := stream.CloseAndRecv()
// 	log.Printf("client streaming finished")
// 	if err != nil {
// 		log.Fatalf("error while receiving %v", err)
// 	}
// 	log.Printf("%v", res.Message)

// }
func callSayHelloClientStream(stream pb.GreetServiceClient, names *pb.Nameslist) {
	log.Printf("streaming started")
	clientStream, err := stream.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names:%v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := clientStream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		log.Printf("send the request with name %s", name)
		time.Sleep(2 * time.Second)
	}

	if err := clientStream.CloseSend(); err != nil {
		log.Fatalf("could not close stream:%v", err)
	}
	log.Printf("client streaming finished")
}
