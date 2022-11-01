package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	ListenAddress string        `yaml:"ListenAddress"`
	ListenPort    int           `yaml:"ListenPort"`
	PinAssignment PinAssignment `yaml:"PinAssignment"`
}

func (c Config) GetListenAddress() string {
	return fmt.Sprintf("%s:%d", c.ListenAddress, c.ListenPort)
}

func (c Config) validate() error {
	if len(c.PinAssignment) == 0 {
		return errors.New("pin assignment is missing")
	}

	return nil
}

type PinAssignment map[string]uint8

func readConfig(filePath string) (Config, error) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(fileBytes, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func GetConfig(filePath string) (Config, error) {
	cfg, err := readConfig(filePath)
	if err != nil {
		return Config{}, err
	}

	if err := cfg.validate(); err != nil {
		return Config{}, err
	}

	if cfg.ListenAddress == "" {
		cfg.ListenAddress = "0.0.0.0"
	}
	if cfg.ListenPort == 0 {
		cfg.ListenPort = 8080
	}

	return cfg, nil
}
