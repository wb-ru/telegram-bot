package Telegram

import (
	env "../../internal/env"
	"../models"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

//Подключение и запуск работы телеграм бота
func TelegramBot(ctx context.Context, config env.ConfigAPI) {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI(config.TokenApi) //боевой токен
	if err != nil {
		log.Panic(err)
	}
	log.Println("Start Bot")

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	upd, err1 := bot.GetUpdatesChan(ucfg)
	if err1 != nil {
		log.Println("bot.GetUpdatesChan err: ", err1)
	}

	for {
		select {
		case reply := <-models.ChanOrderPay:
			// Создаем сообщение
			msg := tgbotapi.NewMessage(config.ChatId, reply) //боевой чат
			// и отправляем его
			_, err := bot.Send(msg)
			if err != nil {
				log.Println("Error to send message", err)
			}
		case update := <-upd:
			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			if update.Message.Text == "/status" {
				// Создаем сообщение
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Программа работает!")
				// и отправляем его
				_, err := bot.Send(msg)
				if err != nil {
					log.Println("Error to send message", err)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
