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

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
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
