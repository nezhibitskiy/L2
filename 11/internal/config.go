package internal

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Ip   string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return &Config{
		Port: os.Getenv("APP_PORT"),
		Ip:   os.Getenv("APP_HOST"),
	}
}
