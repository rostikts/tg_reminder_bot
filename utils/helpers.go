package utils

import (
	"github.com/rostikts/tg_reminder_bot/models/user"
	"gopkg.in/telebot.v3"
	"log"
)

func GetCurrentUser(message telebot.Message) *user.User {
	currentUser := user.User{
		ID:   int(message.Chat.ID),
		Name: message.Chat.Username,
	}
	if err := currentUser.GetToken(); err != nil {
		log.Printf("error during token extraction %s", err.Error())
	}
	return &currentUser
}
