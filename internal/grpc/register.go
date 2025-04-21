package server

import (
	"github.com/hashicorp/consul/api"
)

func RegisterToConsul() {

	config := api.DefaultConfig()
	client, _ := api.NewClient(config)

	reg := &api.AgentServiceRegistration{
		ID:      "user-service",
		Name:    "user-service",
		Port:    50051,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			GRPC:     "127.0.0.1:50051",
			Interval: "10s",
		},
	}

	client.Agent().ServiceRegister(reg)
}
