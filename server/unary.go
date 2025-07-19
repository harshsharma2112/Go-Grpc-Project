package main

import (
	"context"

	pb "github.com/harshsharma2112/Go-Grpc-Project/proto"
)

func (s *helloserver) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "HELLO",
	}, nil
}
