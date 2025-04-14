package client

import (
	pb "app/api/proto"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

type UserClient struct {
	conn   *grpc.ClientConn
	client pb.HelloServiceClient
}

func NewHelloClient() *UserClient {
	target := discoverUserService()
	if target == "" {
		log.Fatal("user-service not found in consul")
	}

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("无法连接用户服务: %v", err)
	}
	client := pb.NewHelloServiceClient(conn)
	return &UserClient{conn: conn, client: client}
}

func discoverUserService() string {
	config := api.DefaultConfig()
	client, _ := api.NewClient(config)
	services, _ := client.Agent().Services()

	for _, srv := range services {
		if srv.Service == "user-service" {
			return fmt.Sprintf("%s:%d", srv.Address, srv.Port)
		}
	}
	return ""
}

func (u *UserClient) SayHello(token string) (*pb.HelloResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req := &pb.HelloRequest{Name: "Alice"}
	resp, err := u.client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	// 4. 打印结果
	log.Printf("服务端响应: %s", resp.Greeting)
	return resp, err
}
