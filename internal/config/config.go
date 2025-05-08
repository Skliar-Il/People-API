package config

import (
	"github.com/Skliar-Il/People-API/internal/transport/http/client"
	"github.com/Skliar-Il/People-API/pkg/database"
	"github.com/Skliar-Il/People-API/pkg/logger"
	"github.com/Skliar-Il/People-API/pkg/redis"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type serverCfg struct {
	HttpPort string `env:"SERVER_PORT_HTTP" default-env:"8080"`
}

type Config struct {
	Server   serverCfg       `env:"SERVER"`
	DataBase database.Config `env:"POSTGRES"`
	Redis    redis.Config    `env:"REDIS"`
	Client   client.Config   `env:"LINK"`
	Logger   logger.Config   `env:"LOGGER"`
}

func New() (*Config, error) {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
