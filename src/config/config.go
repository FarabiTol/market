package config

var MainConfig *Config

type Config struct {
	PostgresConfig
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
