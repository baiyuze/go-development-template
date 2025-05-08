package container

import (
	"app/internal/grpc/handler"
	pb "app/internal/grpc/proto"
	"app/internal/service"

	"go.uber.org/dig"
	"google.golang.org/grpc"
)

func InitContanier(s *grpc.Server, container *dig.Container) {
	container.Invoke(func(userService service.UserService) {
		// pb.RegisterUserServiceServer(s grpc.ServiceRegistrar, srv pb.UserServiceServer)
		pb.RegisterHelloServiceServer(s, &handler.HelloServer{
			UserService: userService,
		})
	})

}
