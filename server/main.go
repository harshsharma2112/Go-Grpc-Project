package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb.
)

const (
	port = ":8080"
)


type helloserver struct{
	pb.GreetServiceServer
}

func main() {
	//using package for listen
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server %v", err)
	} //now i have the port

	//now i will create the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloserver{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start the server %v", err)
	}

}
