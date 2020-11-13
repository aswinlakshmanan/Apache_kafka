package main

import (
	"context"
	"log"
	"net"

	"./proto"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterDetailsserviceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}

func (*server) Info(ctx context.Context, in *proto.Detailsrequest) (*proto.Detailsresponse, error) {
	result := "Hello " + in.GetUsername().GetName()
	res := proto.detailsresponse{
		Result: result,
	}
	return &res, nil
}
