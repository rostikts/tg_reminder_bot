package main

import (
	"github.com/rostikts/tg_reminder_bot/config"
	"gopkg.in/telebot.v3"
)

func main() {
	config.InitBot()
	config.Bot.Handle("/start", func(ctx telebot.Context) error {
		return ctx.Send("Hello!")
	})
	config.Bot.Start()
}
