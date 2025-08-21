package main

import (
	"context"
	"log"
	"time"

	pb "github.com/harshsharma2112/Go-Grpc-Project/proto"
)

// we passing this interface becoz interfce had those methods
// alias can also be used but it is better to use interface
// In this case, pb.GreetServiceClient is an interface, and stream is a value that implements this interface.
// we can not directly point to interface thats why not using *
func CallSayHelloBidirectionalStreaming(stream pb.GreetServiceClient, names *pb.Nameslist) {
	log.Printf("client stream started")
	//calling that fucntion
	ClientStream, err := stream.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send the name %v", err)
	}

	for _, name := range names.Names {
		//creating the new instance of struct on heap & allocate it to heap
		req := &pb.HelloRequest{
			Name: name,
		}
		err := ClientStream.Send(req)
		if err != nil {
			log.Fatalf("not able to send %v", err)
		}
		log.Printf("sending names %s", names)
		time.Sleep(2 * time.Second)
	}
	err = ClientStream.CloseSend()
	if err != nil {
		log.Fatalf("could not close stream %v", err)
	}
	log.Printf("succesfully send the names")
}
