package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

type bar struct {
	UnimplementedBarServer
}

func (b *bar) Process(ctx context.Context, in *FooRequest) (*FooResponse, error) {
	return &FooResponse{Message: "message"}, nil
}

func serverMain() {
	log.Print("Server Main")
	lis, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatalf("Fail to Listen %v", err)
	}
	s := grpc.NewServer()

	RegisterBarServer(s, &bar{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fail to Serve %v", err)
	}
}

func clientMain() {
	conn, _ := grpc.Dial("localhost:8099", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	c := NewBarClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Process(ctx, &FooRequest{Message: "FooRequest"})
	if err != nil {
		log.Fatalf("error %v", err)
	}
	log.Printf("Response %v", r)
}

func unimplementedMain() {
	log.Fatalf("UnImplemented Parameter")
}

func main() {
	args := os.Args[1:]

	if args[0] == "server" {
		serverMain()
	} else if args[0] == "client" {
		clientMain()
	} else {
		unimplementedMain()
	}

}
