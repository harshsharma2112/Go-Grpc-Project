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
	pb.GreetServiceServer// pb.UnimplementedGreetServiceServer it is empty auotgegenerated struct which have all methods 
	// so basically it help to implememnt those methods which we need only.
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
	//why hellos erver struct becasue You can’t pass individual methods — Go requires a type that implements the full interface.
	pb.RegisterGreetServiceServer(grpcServer, &helloserver{})
	log.Printf("Server failed to start %v", lis.Addr())

	//server listen to port
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start the server %v", err)
	}

}
