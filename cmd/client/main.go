package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/arianiti2/grpc-microservices/gen/go/api/v1"


)

func main() {
	
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	
	client := pb.NewMyServiceClient(conn)

	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	resp, err := client.HealthCheck(ctx, &pb.HealthRequest{ServiceName: "UserModule"})
	if err != nil {
		log.Fatalf("HealthCheck failed: %v", err)
	}

	log.Printf("Status: %s", resp.GetStatus())
}
