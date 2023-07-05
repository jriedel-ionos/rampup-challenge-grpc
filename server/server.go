package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jriedel-ionos/rampup-challenge-grpc/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedEnvVariableServer
}

func (s *server) GetEnvironmentVariable(
	_ context.Context,
	in *pb.GetEnvironmentVariableRequest,
) (
	*pb.GetEnvironmentVariableResponse,
	error,
) {
	value := os.Getenv(in.VariableName)
	fmt.Println(value)
	return &pb.GetEnvironmentVariableResponse{
		Value: value,
	}, nil
}

func main() {
	const Port = "8080"

	listener, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		panic(err)
	}

	log.Println("Backend server started on port " + Port)

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterEnvVariableServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
