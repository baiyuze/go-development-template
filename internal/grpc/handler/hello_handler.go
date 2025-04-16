package handler

import (
	pb "app/internal/grpc/proto"
	"app/internal/service"
	"context"
	"encoding/json"
	"fmt"
)

type HelloServer struct {
	pb.UnimplementedHelloServiceServer
	UserService service.UserService
}

func (s *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	data, err := s.UserService.GetUserOne()
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	message := fmt.Sprintf("你好, %s!", string(jsonData))
	return &pb.HelloResponse{Greeting: message}, nil
}
