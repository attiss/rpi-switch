package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	PinAssignment PinAssignment `yaml:"PinAssignment"`
}

type PinAssignment struct {
	Relay1Pin uint8 `yaml:"Relay1Pin"`
	Relay2Pin uint8 `yaml:"Relay2Pin"`
	Relay3Pin uint8 `yaml:"Relay3Pin"`
	Relay4Pin uint8 `yaml:"Relay4Pin"`
	Relay5Pin uint8 `yaml:"Relay5Pin"`
	Relay6Pin uint8 `yaml:"Relay6Pin"`
	Relay7Pin uint8 `yaml:"Relay7Pin"`
	Relay8Pin uint8 `yaml:"Relay8Pin"`
}

func ReadConfig(filePath string) (Config, error) {
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
