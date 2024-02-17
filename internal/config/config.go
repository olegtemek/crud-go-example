package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Address     string        `env:"APP_URL"`
	Timeout     time.Duration `env:"HTTP_TIMEOUT"`
	LoggerLevel string        `env:"LOGGER_LEVEL"` // DEBUG || INFO || WARN || ERROR
	DatabaseUrl string        `env:"DATABASE_URL"`
}

func NewConfig() (*Config, error) {

	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)

	if err != nil {
		return &cfg, err
	}

	return &cfg, nil
}
