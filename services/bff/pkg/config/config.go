package config

import "os"

type Config struct {
	Port            string
	TreeServiceAddr string
}

func New() *Config {
	return &Config{
		Port:            os.Getenv("PORT"),
		TreeServiceAddr: os.Getenv("TREE_SERVICE_ADDR"),
	}
}
