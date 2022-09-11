package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (app App) handleCommand(message *tgbotapi.Message) (err error) {
	switch message.Command() {
	case "start":
		err = app.commandStart(message)
	default:
		err = app.commandUnknown(message)
	}

	return
}

func (app App) commandStart(message *tgbotapi.Message) (err error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hi!")
	_, err = app.Bot.Send(msg)
	if err != nil {
		return
	}

	return
}

func (app App) commandUnknown(message *tgbotapi.Message) (err error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Не известная команда")
	_, err = app.Bot.Send(msg)
	if err != nil {
		return
	}

	return
}
