package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// this is bidirectional that why we need to defien port othereiwse if it is unary no need to defien the port
const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("failed to conenct %v",err)
	}
	defer conn.Close()
}
