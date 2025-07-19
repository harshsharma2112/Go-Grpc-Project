package main

import (
	"log"
	"net"

	pb "github.com/harshsharma2112/Go-Grpc-Project/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

// we cretaed the struct becaus interface is there in greet_grpc.pb.go
type helloserver struct {
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
	//registering service
	pb.RegisterGreetServiceServer(grpcServer, &helloserver{})
	log.Printf("Server failed to start %v", lis.Addr())

	//server listen to port
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start the server %v", err)
	}

}
