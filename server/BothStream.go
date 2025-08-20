package main

import (
	"io"
	"log"
	"time"

	pb "github.com/harshsharma2112/Go-Grpc-Project/proto"
)

func (s *helloserver) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received names %v", req.Name)
		resp := &pb.HelloResponse{
			//creta the map
			Message: "Hello" + req.Name,
		}

		if err := stream.Send(resp); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}

}
