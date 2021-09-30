package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Config of service
type Config struct {
	Address string
}

// ParseConfig of service
func ParseConfig(configPath string) (Config, error) {
	cfg := Config{}

	_, err := toml.DecodeFile(configPath, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("code 2: %w", err)
	}
	return cfg, nil
}
