package dto

type Config struct {
	Service struct {
		Name   string
		Env    string
		Port   string
		Region string
	}

	Consul struct {
		Port string
		Host string
	}
}
