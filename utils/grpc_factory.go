package utils

import (
	"fmt"

	"google.golang.org/grpc"
)

// GrpcFactory grpc 创建连接工厂函数
func GrpcFactory[T any](target string, constructor func(cc grpc.ClientConnInterface) T) (T, *grpc.ClientConn) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败:", err)
	}
	client := constructor(conn)
	return client, conn
}
