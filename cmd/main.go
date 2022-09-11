package main

import (
	"log"

	"smart-doors-tg/internal/config"
	"smart-doors-tg/internal/handlers"
)

func main() {
	appConfig, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalln(err)
	}

	app, err := handlers.NewApp(appConfig)
	if err != nil {
		log.Fatalln(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
