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

type PubSubConfig struct {
	ProjectID string
	TopicName string
}

func NewPubSubConfig() *PubSubConfig {
	return &PubSubConfig{
		ProjectID: os.Getenv("PROJECT_ID"),
		TopicName: os.Getenv("TOPIC_NAME"),
	}
}
