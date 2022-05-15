package main

import (
	"github.com/rostikts/tg_reminder_bot/config"
	"github.com/rostikts/tg_reminder_bot/handlers"
	"gopkg.in/telebot.v3"
)

func main() {
	config.InitBot()
	config.Bot.Handle("/start", func(ctx telebot.Context) error {
		return ctx.Send(config.BackendURL)
	})
	h := handlers.NewHandler()
	config.Bot.Handle("/create", h.CreateHandler)
	config.Bot.Start()
}
