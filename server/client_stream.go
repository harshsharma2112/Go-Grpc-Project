package main

import (
	"io"
	"log"

	pb "github.com/harshsharma2112/Go-Grpc-Project/proto"
)

func (s *helloserver) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var names []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Name: names}) // Use the correct field name as per your proto definition
		}
		if err != nil {
			return err
		}
		log.Printf("got req with this %v", req.Name)
		names = append(names, "Hello "+req.Name)
	}
}
