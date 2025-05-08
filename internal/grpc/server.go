package server

import (
	ctr "app/internal/grpc/container"
	"fmt"
	"log"
	"net"

	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func IntServer(container *dig.Container) {
	go RegisterToConsul()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 注册健康检查服务
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	ctr.InitContanier(s, container)
	fmt.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
