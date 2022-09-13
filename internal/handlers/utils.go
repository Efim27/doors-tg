package handlers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"smart-doors-tg/internal/requests"
)

func (app App) NotifyAdmins(text string) (err error) {
	AdminListTG, err := app.ClientAPI.AdminListTG()
	if err != nil {
		return
	}

	for _, AdminChatID := range AdminListTG {
		msg := tgbotapi.NewMessage(AdminChatID, text)

		_, err = app.Bot.Send(msg)
		if err != nil {
			return
		}
	}

	return
}

func (app App) NotifyAdminsNoAccess(user requests.User) (err error) {
	text := fmt.Sprintf("Не удачная попытка входа от пользователя %v %v", user.Name, user.Surname)

	return app.NotifyAdmins(text)
}

func (app App) SendTextMsg(chatID int64, text string) (err error) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err = app.Bot.Send(msg)

	return
}

func (app App) NotifyError(chatID int64) (err error) {
	return app.SendTextMsg(chatID, "Ошибка")
}
