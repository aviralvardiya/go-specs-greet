package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"go-specs-greet/adapters/grpcserver"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &grpcserver.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
