package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "your_package_path/service" // Atualize com o caminho correto
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.A + req.B
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	log.Println("Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
