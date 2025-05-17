package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	JwtSecret string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("env file not found")
	}

	cfg := &Config{
		Port:      os.Getenv("PORT"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}

	return cfg, nil
}