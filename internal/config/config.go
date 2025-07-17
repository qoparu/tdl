package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTP struct {
		Address string `yaml:"address"`
	} `yaml:"http"`
	MQTT struct {
		Broker   string `yaml:"broker"`
		ClientID string `yaml:"client_id"`
		Topic    string `yaml:"topic"`
	} `yaml:"mqtt"`
}

func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
