package handlers

import (
	"github.com/rostikts/tg_reminder_bot/services"
	"gopkg.in/telebot.v3"
	"log"
)

type Handler struct {
	service services.ReminderService
}

func NewHandler() Handler {
	return Handler{service: services.NewReminderService()}
}

func (h Handler) CreateHandler(ctx telebot.Context) error {
	message := ctx.Message()
	err := h.service.CreateReminder(message)
	log.Printf("recieved: %s", err)
	if err != nil {
		ctx.Send(err.Error())
		return err
	}
	ctx.Send("The reminder is created successfully")
	return nil
}
