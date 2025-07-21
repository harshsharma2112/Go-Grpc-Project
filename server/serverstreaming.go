package main

import (
	"log"
	"time"

	pb "github.com/harshsharma2112/Go-Grpc-Project/proto"
)

func (s *helloserver) SayHelloServerStreaming(req *pb.Nameslist, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Received names:%v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		// send the response back to the client
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

// SayHelloServerStreaming is a function that will be called by the client
