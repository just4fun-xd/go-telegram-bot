package bot

import (
	"go-telegram-bot/config"
	"go-telegram-bot/internal/api"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	var replyText string

	switch {
	case strings.HasPrefix(message.Text, "/weather"):
		replyText = handleWeatherCommand(message.Text)
	case strings.HasPrefix(message.Text, "/dog"):
		replyText = handleDogCommand(message.Text)
	case message.Text == "/start":
		replyText = config.StartMessage
	case message.Text == "/help":
		replyText = config.HelpMessage
	default:
		replyText = "Неизвестная команда. Воспользуйтесь /help, чтобы увидеть список доступных команд."
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, replyText)
	bot.Send(msg)

}

func handleWeatherCommand(input string) string {
	parts := strings.SplitN(input, " ", 2)
	if len(parts) < 2 {
		return "Введите команду в формате: /weather [город]"
	}

	city := parts[1]
	weather, err := api.GetWeather(city)
	if err != nil {
		return "Не удалось получить погоду: " + err.Error()
	}
	return weather
}

func handleDogCommand(input string) string {
	parts := strings.SplitN(input, " ", 2)
	if len(parts) < 2 {
		imageUrl, err := api.GetRandomDogImage()
		if err != nil {
			return "Не удалось получить изображение собаки: " + err.Error()
		}
		return imageUrl
	}

	breed := strings.ToLower(parts[1])
	imageURL, err := api.GetDogImageByBreed(breed)
	if err != nil {
		return "Не удалось получить изображение собаки: " + err.Error()
	}
	return imageURL
}
