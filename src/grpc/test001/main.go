package main

import (
	"context"
	"fmt"
	"github.com/sahara-gopher/tour/src/grpc/test001/proto/greet"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloServer struct {
}

func (s *HelloServer) Say(ctx context.Context, in *greet.SayRequest) (*greet.SayResponse, error) {
	return &greet.SayResponse{
		Reply: "reply content",
	}, nil
}

func main() {
	fmt.Println("start ...")
	run()
}

func run() {
	lis, err := net.Listen("tcp", ":8999")
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}
	server := grpc.NewServer()
	greet.RegisterHelloServiceServer(server, &HelloServer{})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
