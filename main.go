package main

import (
	env "./internal/env"
	"./pkg/Telegram"
	ws "./pkg/WaitShutDown"
	"./pkg/db"
	"context"
	"github.com/joho/godotenv"
	"log"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := env.New()
	ctx, cancel := context.WithCancel(context.Background())
	go db.GetDataFromDB(ctx)
	go Telegram.TelegramBot(ctx, conf)

	ws.WaitShutdown(cancel)
}
