package template

var (
	MainSRV = `package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "{{.Dir}}/proto"
)

type srv struct {
	pb.Unimplemented{{title .Alias}}Server
}

// SayHello implements {{.Dir}}.{{title .Alias}}Server
func (s *srv) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.Register{{title .Alias}}Server(s, &srv{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
`
)
