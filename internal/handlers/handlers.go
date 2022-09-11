package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"smart-doors-tg/internal/config"
	"smart-doors-tg/internal/requests"
)

type App struct {
	Bot       *tgbotapi.BotAPI
	UpdatesTG tgbotapi.UpdatesChannel
	Config    config.Config
	ClientAPI requests.ClientAPI
	Logger    *log.Logger
}

func NewApp(config config.Config) (app *App, err error) {
	app = new(App)
	app.Config = config

	app.Bot, err = tgbotapi.NewBotAPI(config.TokenTG)
	if err != nil {
		return
	}

	app.ClientAPI = *requests.NewClientAPI(app.Config.ControlAPIAddr)

	app.Bot.Debug = app.Config.Debug
	log.Printf("Authorized on account %s\n", app.Bot.Self.UserName)
	return
}

func (app *App) loadUpdatesCh() (updates tgbotapi.UpdatesChannel, err error) {
	if !app.Config.IsWebhook {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 3600
		updates = app.Bot.GetUpdatesChan(u)

		return
	}

	file, err := os.Open("./cert/cert.pem")
	if err != nil {
		return
	}

	certFR := tgbotapi.FileReader{
		Name:   "cert.pem",
		Reader: file,
	}

	whAddr := fmt.Sprintf("https://%v:%v/%v", app.Config.AppAddr, app.Config.AppPort, app.Config.TokenTG)
	wh, _ := tgbotapi.NewWebhookWithCert(whAddr, certFR)

	_, err = app.Bot.Request(wh)
	if err != nil {
		return
	}

	info, err := app.Bot.GetWebhookInfo()
	if err != nil {
		return
	}

	if info.LastErrorDate != 0 {
		err = errors.New(fmt.Sprintf("Telegram callback failed: %s", info.LastErrorMessage))
		return
	}

	updates = app.Bot.ListenForWebhook("/" + app.Config.TokenTG)
	go http.ListenAndServeTLS(fmt.Sprintf("0.0.0.0:%v", app.Config.AppPort), "./cert/cert.pem", "./cert/key.pem", nil)

	return
}

func (app *App) Run() (err error) {
	app.UpdatesTG, err = app.loadUpdatesCh()
	if err != nil {
		return
	}

	for update := range app.UpdatesTG {
		// ignore any non-Message updates
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			err = app.handleCommand(update.Message)
			if err != nil {
				log.Println(err)
			}

			continue
		}

		err = app.handleMessage(update.Message)
		if err != nil {
			log.Println(err)
		}
	}

	return
}
