package AppContext

import (
	"app/config"
	"app/internal/dto"
	pb "app/internal/grpc/proto"
	"app/utils"
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AppContext struct {
	Config *dto.Config
	Logger *zap.Logger

	UserClient pb.HelloServiceClient
	UserConn   *grpc.ClientConn
}

type ServiceDiscovery struct {
	client *api.Client
}

// var AppContext *AppContext

// 初始化发现器
func NewServiceDiscovery() *ServiceDiscovery {

	config := api.DefaultConfig()
	config.Address = "consul.sanyang.life:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("创建 Consul 客户端失败: %v", err)
	}
	return &ServiceDiscovery{client: client}
}

// 获取服务地址（只查一次）
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

func InitClient(logger *zap.Logger) *AppContext {

	client, conn := newClient("user-service", pb.NewHelloServiceClient)
	return &AppContext{
		Config:     config.Cfg,
		UserClient: client,
		UserConn:   conn,
		Logger:     logger,
	}
	// return Context
}
