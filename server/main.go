package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"rampup-challenge/server/pb"
)

type server struct {
	pb.UnimplementedEnvVariableServer
}

func (s *server) GetEnvironmentVariable(ctx context.Context, in *pb.GetEnvironmentVariableRequest) (*pb.GetEnvironmentVariableResponse, error) {
	value := os.Getenv(in.VariableName)
	fmt.Println(value)
	return &pb.GetEnvironmentVariableResponse{
		Value: value,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterEnvVariableServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
