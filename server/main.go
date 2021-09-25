package main

import (
	"log"
	"net"

	"github.com/hokita/grpc_practice/service"
	pb "github.com/hokita/grpc_practice/user"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagerServer(s, &service.UserService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
