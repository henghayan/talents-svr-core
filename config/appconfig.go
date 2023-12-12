package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	DBPassword string
	// 其他配置项
}

func LoadConfig() (*AppConfig, error) {
	viper.AutomaticEnv() // 自动读取环境变量

	var config AppConfig
	err := viper.Unmarshal(&config)
	return &config, err
}
