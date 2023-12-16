package config

import "os"

type Config struct {
	Port                    string
	FirebaseAuthCredentials string
}

func New() *Config {
	return &Config{
		Port:                    os.Getenv("PORT"),
		FirebaseAuthCredentials: os.Getenv("FIREBASE_AUTH_CREDENTIALS"),
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
