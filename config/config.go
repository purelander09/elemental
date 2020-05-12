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
		Url      string `yaml:"url" env:"DB_HOST" env-description:"Database host" env-default:"localhost"`
		Username string `yaml:"username" env:"DB_USERNAME" env-description:"Database user name" env-default:"username"`
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
