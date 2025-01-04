package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToket string
}

func LoadConfit() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		BotToket: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}, nil
}
