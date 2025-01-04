package bot

import (
	"go-telegram-bot/internal/api"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	var replyText string

	if strings.HasPrefix(message.Text, "/weather") {
		parts := strings.SplitN(message.Text, " ", 2)
		if len(parts) < 2 {
			replyText = "Введите команду в формате: /weather [город]"
		} else {
			city := parts[1]
			weather, err := api.GetWeather(city)
			if err != nil {
				replyText = "Не удалось получить погоду: " + err.Error()
			} else {
				replyText = weather
			}

		}
	} else {
		switch message.Text {
		case "/start":
			replyText = "Добро пожаловать! Я ваш помощник."
		case "Привет":
			replyText = "Привет! Чем могу помочь?"
		default:
			replyText = "Извините, я не понял ваш запрос."
		}
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, replyText)
	bot.Send(msg)

}
