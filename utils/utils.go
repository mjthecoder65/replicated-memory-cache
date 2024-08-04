package utils

import (
	"log"
	"os"

	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`
	Discovery   *api.Config `yaml:"discovery"`
	Environment struct {
		Value string `yaml:"value"`
	} `yaml:"environment"`

	Database struct {
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
	} `yaml:"database"`
}

func LoadConfig() *Config {
	file, err := os.ReadFile("config/config.yaml")

	if err != nil {
		panic("failed to read configuration file")
	}

	var config Config

	err = yaml.Unmarshal(file, &config)

	if err != nil {
		log.Fatal(err)
	}

	return &config
}
