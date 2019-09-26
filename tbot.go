package main

import (
  "log"

  "github.com/davecgh/go-spew/spew"
  "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
  TOKEN string = "909673712:AAEQ7YZdxJRoPhz90rPucxOPGTPTFCL6wf0"
  DEBUG bool   = false
  TIMEOUT int  = 60
)

func main() {
  // подключаемся к боту с помощью токена
  bot, err := tgbotapi.NewBotAPI(TOKEN)
  if err != nil {
    log.Panic(err)
  }

  bot.Debug = DEBUG
  log.Printf("Авторизация в аккаунте @%s", bot.Self.UserName)

  // инициализируем канал, куда будут прилетать обновления от API
  u := tgbotapi.NewUpdate(0)
  u.Timeout = TIMEOUT

  updates, err := bot.GetUpdatesChan(u)
  if err != nil {
    log.Fatal(err)
  }

  // читаем обновления из канала
  for update := range updates {
    if (update.Message != nil) {
      OnMessage(bot, update.Message)
    }
  }
}

// Отправляем сообщение в телеграм канал
func OnMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

  var msg tgbotapi.Chattable
  var sendMessage string
  //var sendType string = "text"

  userName := message.From.UserName
  chatID   := message.Chat.ID

  log.Printf("[%s] %d", userName, chatID)
  spew.Dump(message) // выводим то что пришло (Для отладки!!!)

  if message.IsCommand() {
    command := message.Command()

    switch command {
      case "start" :
        sendMessage = "Добро пожаловать, " + message.Chat.FirstName + " " + message.Chat.LastName + "!"
        msg = tgbotapi.NewMessage(chatID, sendMessage)
      case "help" :
        sendMessage = "Список доступных комманд:\n1: /BACK_TO_WORK\n2: /start"
        msg = tgbotapi.NewMessage(chatID, sendMessage)
      case "BACK_TO_WORK" :
        sendMessage = "photo.jpg"
        msg = tgbotapi.NewPhotoUpload(chatID, sendMessage)
      case "sticker" :
        sendMessage = "CAADAgAD8QIAApzW5wooEroAAWAF9ykWBA"
        msg = tgbotapi.NewStickerShare(chatID, sendMessage)
      default :
        sendMessage = command
    }

    _, err := bot.Send(msg) // отправляем сообщение
    if err != nil {
      log.Println(err) // распечатываем ошибки, если есть
    }
  }
}