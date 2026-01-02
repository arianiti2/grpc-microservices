package service

import (
	"context"

	pb "github.com/arianiti2/grpc-microservices/gen/go/api/v1"
)


type HealthService struct {
	pb.UnimplementedMyServiceServer 
}


func (s *HealthService) HealthCheck(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{
		Status: "Service " + req.ServiceName + " is running in 2026!",
	}, nil
}
