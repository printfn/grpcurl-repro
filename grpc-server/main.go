package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/printfn/grpcurl-repro/protobuf-service/servicepb"
	"github.com/printfn/grpcurl-repro/protobuf-shared/sharedpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	servicepb.UnimplementedTestServiceServer
}

func (s *server) GetHelloWorld(ctx context.Context, in *emptypb.Empty) (*sharedpb.HelloWorldResponse, error) {
	log.Printf("Received message")
	return &sharedpb.HelloWorldResponse{
		Result: fmt.Sprintf("Hello World at %s", time.Now().Format(time.DateTime)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	servicepb.RegisterTestServiceServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
