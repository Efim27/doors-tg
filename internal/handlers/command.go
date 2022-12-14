package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
	newUser := message.From
	err = app.ClientAPI.NewUserTG(newUser.ID, message.Chat.ID, newUser.UserName, newUser.FirstName, newUser.LastName, newUser.LanguageCode)
	if err != nil {
		return app.NotifyError(message.Chat.ID)
	}

	err = app.SendTextMsg(message.Chat.ID, "Вы успешно отправили заявку на привязку телеграм аккаунта")
	if err != nil {
		return
	}

	return
}

func (app App) commandUnknown(message *tgbotapi.Message) (err error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда")
	_, err = app.Bot.Send(msg)
	if err != nil {
		return
	}

	return
}
