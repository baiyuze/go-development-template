package container

import (
	"app/internal/container"
	"app/internal/grpc/handler"
	pb "app/internal/grpc/proto"

	"google.golang.org/grpc"
)

func InitContanier(s *grpc.Server, Deps *container.AppDependency) {
	// pb.RegisterUserServiceServer(s grpc.ServiceRegistrar, srv pb.UserServiceServer)
	pb.RegisterHelloServiceServer(s, &handler.HelloServer{
		UserService: Deps.UserService,
	})
}
