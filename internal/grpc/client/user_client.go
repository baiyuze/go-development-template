package client

import (
	"app/internal/container"
	pb "app/internal/grpc/proto"
	"context"
	"log"
	"time"
)

func SayHello(token string) (*pb.HelloResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req := &pb.HelloRequest{Name: "吃啥"}
	resp, err := container.Deps.Context.UserClient.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	// 4. 打印结果
	log.Printf("服务端响应: %s", resp.Greeting)
	return resp, err
}
