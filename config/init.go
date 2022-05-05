package config

import (
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

var Bot *telebot.Bot

func InitBot() {
	var err error
	prefs := telebot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	Bot, err = telebot.NewBot(prefs)
	if err != nil {
		log.Panic("Bot was not started")
	}
	BackendURL = os.Getenv("BACKEND_URL")
}
