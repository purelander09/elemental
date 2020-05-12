package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Port string `yaml:"port" env:"SRV_PORT,PORT" env-description:"Server port" env-default:"3333"`
	}
	Database struct {
		Host string `yaml:"host" env:"DB_HOST" env-description:"Database host" env-default:"localhost"`
	}
}

func GetConfiguration(configPath string) Config {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return cfg
}
