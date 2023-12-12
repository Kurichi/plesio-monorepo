package config

import "os"

type Config struct {
	Port string
}

func New() *Config {
	return &Config{
		Port: os.Getenv("PORT"),
	}
}
