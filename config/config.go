package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBName     string
	DBPort     string

	Addr string
}

func LoadConfig() *Config {
	viper.AutomaticEnv()

	config := &Config{
		DBUsername: viper.GetString("DB_USERNAME"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBHost:     viper.GetString("DB_HOST"),
		DBName:     viper.GetString("DB_NAME"),
		DBPort:     viper.GetString("DB_PORT"),
	}

	return config
}

func (cfg *Config) GetAddr() string {
	if cfg.Addr == "" {
		return ":8080"
	}
	return cfg.Addr
}
