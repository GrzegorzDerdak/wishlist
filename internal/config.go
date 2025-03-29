package internal

import (
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	Port      string
	AppDomain string
	DSN       string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return &Config{
		Port:      os.Getenv("PORT"),
		AppDomain: os.Getenv("APP_DOMAIN"),
		DSN:       os.Getenv("DSN"),
	}
}
