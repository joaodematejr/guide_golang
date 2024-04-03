package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"

	pb "your_package_path/service" // Atualize com o caminho correto
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	a, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalf("Invalid argument: %v", err)
	}

	b, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		log.Fatalf("Invalid argument: %v", err)
	}

	req := &pb.AddRequest{A: int32(a), B: int32(b)}
	resp, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call service: %v", err)
	}
	log.Printf("Result: %d\n", resp.Result)
}
