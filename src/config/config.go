package config

import (
	"context"
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

var MainConfig *Config

type Config struct {
	PostgresConfig
	TransportConfig
}

type PostgresConfig struct {
	Host         string `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port         string `envconfig:"POSTGRES_PORT" default:"5432"`
	DatabaseName string `envconfig:"POSTGRES_NAME" default:"market"`
	Username     string `envconfig:"POSTGRES_USER" default:"market"`
	Password     string `envconfig:"POSTGRES_PASS" default:"market"`
	MaxConns     int    `envconfig:"POSTGRES_MAX_CONNS" default:"200"`
	Schema       string `envconfig:"POSTGRES_SCHEMA" default:"public"`
}

type TransportConfig struct {
	HTTPListenAddr string `envconfig:"http_listen_addr" required:"true" default:":8080"`
}

func InitConfigs(ctx context.Context) error {
	var cfg Config

	err := envconfig.Process("APP_CONFIG", &cfg)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	MainConfig = &cfg
	return nil
}
