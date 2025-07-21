package main

import (
	"context"

	pb "github.com/harshsharma2112/Go-Grpc-Project/proto"
)
//server side it is a method whereas on client side it is a funtion
// client will call this method
// we passed context as a parameter because call is coming form client side.
func (s *helloserver) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "HELLO",
	}, nil
}
