package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"github.com/arianiti2/grpc-microservices/internal/service"
	pb "github.com/arianiti2/grpc-microservices/gen/go/api/v1"


)

func main() {
	
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	
	grpcServer := grpc.NewServer()


	pb.RegisterMyServiceServer(grpcServer, &service.HealthService{})

	
	go func() {
		log.Printf("Server is running on %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()


	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh

	log.Println("Stopping the server...")
	grpcServer.GracefulStop() 
	log.Println("server stopped. Goodbye!")
}
