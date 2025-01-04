package main

import (
	"go-telegram-bot/config"
	"go-telegram-bot/internal/bot"
	"log"
)

func main() {
	cfg, err := config.LoadConfit()
	if err != nil {
		log.Fatalf("Ошибка закгрузки конфигурации: %v", err)
	}

	err = bot.StartBot(cfg.BotToket)
	if err != nil {
		log.Fatalf("Ошибка запуска бота: %v", err)
	}
}
