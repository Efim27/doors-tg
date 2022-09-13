package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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
