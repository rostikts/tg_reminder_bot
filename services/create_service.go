package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rostikts/tg_reminder_bot/config"
	"github.com/rostikts/tg_reminder_bot/constants"
	"github.com/rostikts/tg_reminder_bot/models"
	"github.com/rostikts/tg_reminder_bot/models/user"
	"github.com/rostikts/tg_reminder_bot/utils"
	"gopkg.in/telebot.v3"
	"net/http"
	"strconv"
	"strings"
)

type ReminderService struct {
	data *models.Reminder
	user user.User
}

func NewReminderService() ReminderService {
	return ReminderService{
		data: &models.Reminder{},
		user: user.User{},
	}
}

func (s ReminderService) CreateReminder(message *telebot.Message) error {
	if err := s.parseCreatePayload(message.Payload); err != nil {
		return err
	}
	s.user = *utils.GetCurrentUser(*message)
	s.data.UserId = s.user.ID
	s.data.Completed = false
	if err := s.sendCreateRequest(); err != nil {
		return err
	}
	return nil
}

func (s ReminderService) parseCreatePayload(payload string) error {
	payloadList := strings.SplitN(payload, " ", 3)
	s.data.Date = payloadList[0] + "T00:00:00Z"
	s.data.Time = payloadList[1] + ":0.0"
	s.data.Title = payloadList[2]
	return nil
}

func (s ReminderService) sendCreateRequest() error {
	url := config.BackendURL + constants.CreateReminderUrl
	client := http.DefaultClient

	payload, err := json.Marshal(s.data)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(payload)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	token := fmt.Sprintf("Bearer %s", s.user.AuthToken)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	var res interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	if resp.StatusCode != 201 {
		return errors.New("smth gone wrong :(\nStatus code: " + strconv.Itoa(resp.StatusCode))
	}
	return nil
}
