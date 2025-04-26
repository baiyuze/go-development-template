package server

import (
	"fmt"
	"os"

	"github.com/hashicorp/consul/api"
)

func RegisterToConsul() {

	config := api.DefaultConfig()
	addr := os.Getenv("ADDR")
	fmt.Println(addr, "addraddraddr")
	config.Address = addr
	client, _ := api.NewClient(config)

	reg := &api.AgentServiceRegistration{
		ID:      "user-service",
		Name:    "user-service",
		Port:    50051,
		Address: "192.168.2.136",
		Check: &api.AgentServiceCheck{
			GRPC:     "192.168.2.136:50051",
			Interval: "10s",
		},
	}

	client.Agent().ServiceRegister(reg)
}
