
# Telegram Bot for API Integration

[![Go](https://img.shields.io/badge/-Go-00ADD8?style=flat&logo=Go&logoColor=ffffff)](https://golang.org/)
[![Telegram](https://img.shields.io/badge/-Telegram-2CA5E0?style=flat&logo=telegram&logoColor=ffffff)](https://core.telegram.org/bots)

Telegram Bot for API Integration — это проект для тренировки работы с Telegram Bot API и интеграции сторонних API в Go-приложение.  
Бот предоставляет функции для получения текущей погоды и случайных изображений собак.

---

## **Особенности проекта**

- Поддержка команд для получения погоды через OpenWeatherMap API.
- Интеграция Dog CEO's Dog API для случайных изображений собак.
- Лёгкий для понимания и расширения код.

---

## **Технологии**

- **Язык:** Go
- **Telegram API:** [tgbotapi](https://github.com/go-telegram-bot-api/telegram-bot-api)
- **Погодный API:** [OpenWeatherMap](https://openweathermap.org/api)
- **API для собак:** [Dog CEO's Dog API](https://dog.ceo/dog-api/)

---

## **Как запустить проект**

### **1. Клонируйте репозиторий**

```bash
git clone https://github.com/your-repo/go-telegram-bot.git
cd go-telegram-bot
```

### **2. Настройка переменных окружения**

Создайте файл `.env` в корневой директории и добавьте ваши API-ключи:
```plaintext
TELEGRAM_BOT_TOKEN=ваш_токен_бота
WEATHER_API_KEY=ваш_ключ_openweathermap
```

### **3. Установите зависимости**

Убедитесь, что у вас установлен Go версии 1.20 или выше. Установите зависимости:
```bash
go mod tidy
```

### **4. Запустите бота**

```bash
go run cmd/main.go
```

---

## **Примеры использования**

### Получение погоды
```plaintext
/weather Москва
```
**Ответ**:
```
В городе Москва сейчас ясное небо, температура 15.3°C
```

### Случайное изображение собаки
```plaintext
/dog
```
**Ответ**:  
*(URL случайного изображения собаки)*

### Изображение собаки определенной породы
```plaintext
/dog husky
```
**Ответ**:  
*(URL изображения собаки породы хаски)*

---

## **Структура проекта**

```
go-telegram-bot/
├── README.md                  # Документация проекта
├── cmd/
│   └── main.go                # Точка входа
├── config/
│   ├── config.go              # Конфигурация приложения
│   └── constants.go           # Статические константы
├── go.mod                     # Зависимости проекта
├── go.sum                     # Контрольные суммы зависимостей
├── internal/
│   ├── api/                   # Взаимодействие с внешними API
│   │   ├── dog.go             # Работа с Dog CEO's Dog API
│   │   └── weather.go         # Работа с OpenWeatherMap API
│   ├── bot/                   # Логика Telegram-бота
│   │   ├── bot.go             # Инициализация бота
│   │   └── handler.go         # Обработка команд
│   └── utils/                 # Вспомогательные утилиты (пока пусто)
└── pck/                       # Дополнительные пакеты (если появятся)
```

---

## **Идеи для улучшения**

1. **Добавить больше API**:
   - Trivia API для создания викторин.
   - API для генерации цитат или шуток.
2. **Добавить кнопки Telegram**:
   - Использовать inline-кнопки для выбора пород собак или городов.
3. **Логирование**:
   - Добавить логирование для отладки запросов и ошибок.
4. **Тесты**:
   - Реализовать тесты для API и обработчиков команд.

---

## **Лицензия**

Этот проект распространяется под лицензией MIT. Вы можете использовать его для обучения и разработки собственных проектов.

---

## **Контакты**

Если у вас есть вопросы или предложения, свяжитесь со мной через Telegram или GitHub.
