package handlers

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (app App) handleMessage(message *tgbotapi.Message) (err error) {
	msgText := strings.ToLower(strings.TrimSpace(message.Text))

	if strings.Contains(msgText, "подтвердить") {
		err = app.handleMsgConfirm(message)
	} else if strings.Contains(msgText, "отмена") {
		err = app.handleMsgCancel(message)
	} else {
		err = app.commandUnknown(message)
	}

	return
}

func (app App) handleMsgConfirm(message *tgbotapi.Message) (err error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Вы успешно открыли помещение")
	_, err = app.Bot.Send(msg)

	err = app.NotifyAdmins("")
	if err != nil {
		return
	}

	return
}

func (app App) handleMsgCancel(message *tgbotapi.Message) (err error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Вы успешно отменили запрос на подтверждение")
	_, err = app.Bot.Send(msg)
	if err != nil {
		return
	}

	return
}
