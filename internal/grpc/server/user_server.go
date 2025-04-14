package main

import (
	pb "app/api/proto"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()
	message := fmt.Sprintf("你好, %s!", name)
	return &pb.HelloResponse{Greeting: message}, nil
}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"

// 	pb "example.com/helloworld/helloworldpb"
// 	"google.golang.org/grpc"
// )

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Greeting: "test"}, nil
}
func registerToConsul() {
	config := api.DefaultConfig()
	client, _ := api.NewClient(config)

	reg := &api.AgentServiceRegistration{
		ID:      "user-service",
		Name:    "user-service",
		Port:    50051,
		Address: "localhost",
		Check: &api.AgentServiceCheck{
			GRPC:     "localhost:50051",
			Interval: "10s",
		},
	}

	client.Agent().ServiceRegister(reg)
}
func main() {
	go registerToConsul()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	fmt.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
