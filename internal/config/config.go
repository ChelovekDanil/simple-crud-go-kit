package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Port string `yaml:"server"`
	} `yaml:"server"`
	Database struct {
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Name string `yaml:"name"`
		Db   string `yaml:"db"`
	} `yaml:"database"`
}

const (
	configPath = "/home/asscamper/godir/crud/configs/config.yaml"
)

func MustLoad() *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
