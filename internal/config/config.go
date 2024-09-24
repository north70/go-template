package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Environment string

const (
	Local Environment = "local"
	Dev   Environment = "dev"
	Prod  Environment = "prod"
)

type AppConfig struct {
	ServiceName string      `mapstructure:"service_name"`
	HTTPPort    int         `mapstructure:"http_port"`
	GRPCPort    int         `mapstructure:"grpc_port"`
	LogLevel    string      `mapstructure:"log_level"`
	Environment Environment `mapstructure:"environment"`
	MetricsPort int         `mapstructure:"metrics_port"`
}

type DatabaseConfig struct {
	PostgresDSN   string `mapstructure:"postgres_dsn" env:"POSTGRES_DSN"`
	RedisAddr     string `mapstructure:"redis_addr" env:"REDIS_ADDR"`
	RedisPassword string `mapstructure:"redis_password" env:"REDIS_PASSWORD"`
}

type ClientConfig struct {
	ExternalServiceEndpoint string `mapstructure:"external_service_endpoint"`
}

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Client   ClientConfig   `mapstructure:"client"`
}

func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()

	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	v.AutomaticEnv()

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}
	parseEnv(&config)

	return &config, nil
}

func parseEnv(cfg *Config) {
	if cfg.App.Environment != Local {
		cfg.Database.PostgresDSN = os.Getenv("POSTGRES_DSN")
		cfg.Database.RedisAddr = os.Getenv("REDIS_ADDR")
		cfg.Database.RedisPassword = os.Getenv("REDIS_PASSWORD")
	}
}
