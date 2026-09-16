package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config содержит настройки сервиса temperature-api.
type Config struct {
	Address string `envconfig:"ADDRESS" default:":8081"`
}

// Load загружает конфигурацию из переменных окружения.
func Load() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, fmt.Errorf("not able to read environment variables: %w", err)
	}
	return &c, nil
}
