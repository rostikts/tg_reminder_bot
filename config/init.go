package config

import (
	"gopkg.in/telebot.v3"
	"log"
	"time"
)

var Bot *telebot.Bot

func InitBot() {
	var err error
	prefs := telebot.Settings{
		Token:  "5224732259:AAEFBjJPwbHFksEJ-e38Trf1gNX2MqfwL3s",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	Bot, err = telebot.NewBot(prefs)
	if err != nil {
		log.Panic("Bot was not started")
	}
}
