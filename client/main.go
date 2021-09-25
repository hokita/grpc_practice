package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/hokita/grpc_practice/user"
	"google.golang.org/grpc"
)

const (
	address   = "localhost:50051"
	defaultID = 1
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagerClient(conn)

	id := defaultID
	if len(os.Args) > 1 {
		id, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("argumant error: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &pb.UserRequest{Id: uint32(id)})
	if err != nil {
		log.Fatalf("getting user failed: %v", err)
	}
	log.Printf("name: %s", r)
}
