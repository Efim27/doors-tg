package handlers

import (
	"errors"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"smart-doors-tg/internal/requests"
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
	owner, err := app.ClientAPI.OwnerUserByTG(message.From.ID)
	if (err != nil) && !errors.Is(err, requests.ErrNoAccess) {
		return app.NotifyError(message.Chat.ID)
	}

	if errors.Is(err, requests.ErrNoAccess) {
		log.Println(app.NotifyAdminsNoAccess(owner))
		return app.SendTextMsg(message.Chat.ID, "У Вас нет доступа к помещению или Ваш аккаунт еще не привязан")
	}

	err = app.ClientAPI.CheckAccess(owner.Name, owner.Surname)
	if err != nil {
		log.Println(app.NotifyAdminsNoAccess(owner))
		return app.SendTextMsg(message.Chat.ID, "У Вас нет доступа к помещению")
	}

	adminNotifyText := fmt.Sprintf("%v %v открыл помещение", owner.Name, owner.Surname)
	log.Println(app.NotifyAdmins(adminNotifyText))
	return app.SendTextMsg(message.Chat.ID, "Вы успешно открыли помещение")
}

func (app App) handleMsgCancel(message *tgbotapi.Message) (err error) {
	owner, err := app.ClientAPI.OwnerUserByTG(message.From.ID)
	if (err != nil) && !errors.Is(err, requests.ErrNoAccess) {
		return app.NotifyError(message.Chat.ID)
	}

	if errors.Is(err, requests.ErrNoAccess) {
		log.Println(app.NotifyAdminsNoAccess(owner))
		return app.SendTextMsg(message.Chat.ID, "У Вас нет доступа к помещению или Ваш аккаунт еще не привязан")
	}

	adminNotifyText := fmt.Sprintf("%v %v отменил запроc на открытие помещения", owner.Name, owner.Surname)
	log.Println(app.NotifyAdmins(adminNotifyText))

	return app.SendTextMsg(message.Chat.ID, "Вы успешно открыли помещение")
}
