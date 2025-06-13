package container

import (
	pb "app/internal/grpc/proto"
	"app/utils"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/consul/api"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

type Clients struct {
	UserClient pb.HelloServiceClient
	UserConn   *grpc.ClientConn
}

type ServiceDiscovery struct {
	client *api.Client
}

// NewServiceDiscovery 初始化发现器
func NewServiceDiscovery() *ServiceDiscovery {

	config := api.DefaultConfig()
	addr := os.Getenv("ADDR")
	if len(addr) != 0 {
		config.Address = addr
	}
	// Grpc 等待时间不能过长，不然不能访问
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("创建 Consul 客户端失败: %v", err)
	}
	return &ServiceDiscovery{client: client}
}

// GetServiceAddress 获取服务地址（只查一次）
func (d *ServiceDiscovery) GetServiceAddress(serviceName string) (string, error) {
	services, err := d.client.Agent().Services()
	if err != nil {
		return "", err
	}

	for _, service := range services {
		if service.Service == serviceName {
			return fmt.Sprintf("%s:%d", service.Address, service.Port), nil
		}
	}

	return "", fmt.Errorf("服务 %s 未找到", serviceName)
}

// func
func newClient[T any](serverName string, constructor func(grpc.ClientConnInterface) T) (T, *grpc.ClientConn) {
	discover := NewServiceDiscovery()
	target, err := discover.GetServiceAddress(serverName)
	if err != nil {
		log.Fatalf("获取TargetName失败: %v", err)
	}
	client, conn := utils.GrpcFactory[T](target, constructor)
	return client, conn
}

func InitClients() *Clients {
	client, conn := newClient("user-service", pb.NewHelloServiceClient)
	return &Clients{
		UserClient: client,
		UserConn:   conn,
	}
}

func NewProvideClients(container *dig.Container) {
	container.Provide(InitClients)
}
