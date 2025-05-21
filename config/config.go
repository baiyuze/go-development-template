package config

import (
	"app/internal/dto"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/dig"
)

func InitConfig() (*dto.Config, error) {

	envName := os.Getenv("APP_ENV")
	var c dto.Config

	if envName == "" {
		envName = "dev"
	}
	viper.SetConfigType("yaml")
	viper.SetConfigName("config." + envName)
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		return &c, fmt.Errorf("读取配置失败: %w", err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		return &c, fmt.Errorf("解析配置失败: %w", err)
	}

	viper.WatchConfig()

	return &c, nil

}

func ProvideConfig(contanier *dig.Container) {
	contanier.Provide(InitConfig)
}
