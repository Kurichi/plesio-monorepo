package config

import "os"

type Config struct {
	Port                    string
	TreeServiceAddr         string
	FirebaseAuthCredentials string
}

func New() *Config {
	return &Config{
		Port:                    os.Getenv("PORT"),
		TreeServiceAddr:         os.Getenv("TREE_SERVICE_ADDR"),
		FirebaseAuthCredentials: os.Getenv("FIREBASE_AUTH_CREDENTIALS"),
	}
}
