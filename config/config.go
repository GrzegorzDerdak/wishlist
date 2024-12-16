package config

import "os"

type Config struct {
	Port        string
	AppDomain   string
	DatabaseUrl string
}

func Load() *Config {
	return &Config{
		Port:        os.Getenv("PORT"),
		AppDomain:   os.Getenv("APP_DOMAIN"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}
}
