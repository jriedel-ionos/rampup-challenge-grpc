package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

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
	port := flag.Int("port", 8080, "port for the backend")
	flag.Parse()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if err != nil {
		panic(err)
	}

	log.Printf("Backend server started on port %v", *port)

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterEnvVariableServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
