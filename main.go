package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatalf("Fail to Listen %v", err)
	}
	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fail to Serve %v", err)
	}
}
