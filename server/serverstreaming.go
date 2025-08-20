package main

import (
	"log"
	"time"

	pb "github.com/harshsharma2112/Go-Grpc-Project/proto"
)

//req *pb.NamesList is a pointer to a NamesList message, allowing efficient memory management and access to message fields. Using pointers to message types is a common pattern in gRPC and protocol buffer programming in Go.

// the stream object is passed as a parameter to the function because it represents the existing connection to the client, and it's used to send messages to the client.
func (s *helloserver) SayHelloServerStreaming(req *pb.Nameslist, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Received names:%v", req.Names)
	for _, name := range req.Names {
		//When we write res := &pb.HelloResponse, we are creating a pointer to the address of the HelloResponse struct.
		//The & operator returns the memory address of the HelloResponse struct, and the res variable is assigned to hold that memory address.
		res := &pb.HelloResponse{
			//map
			Message: "Hello " + name,
		}
		// send the response back to the client
		//fucntion send return error 
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

// SayHelloServerStreaming is a function that will be called by the client
