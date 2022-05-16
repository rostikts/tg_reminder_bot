package utils

import (
	"github.com/rostikts/tg_reminder_bot/models/user"
	"gopkg.in/telebot.v3"
	"log"
	"strconv"
	"time"
)

func GetCurrentUser(message telebot.Message) *user.User {
	currentUser, found := TokenCache.Get(strconv.FormatInt(message.Chat.ID, 10))

	if !found {

		currentUser = &user.User{
			ID:   int(message.Chat.ID),
			Name: message.Chat.Username,
		}
		if err := currentUser.(*user.User).GetToken(); err != nil {
			log.Printf("error during token extraction %s", err.Error())
		}
		TokenCache.Set(strconv.FormatInt(message.Chat.ID, 10), currentUser, time.Minute*50)
	}
	return currentUser.(*user.User)
}
