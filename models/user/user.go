package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/rostikts/tg_reminder_bot/config"
	"github.com/rostikts/tg_reminder_bot/constants"
	"log"
	"net/http"
	"strconv"
	"time"
)

type loginData struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func getLoginJSON(data *loginData) (*bytes.Buffer, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(payload)
	return body, nil
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AuthToken string
}

type responseLoginData struct {
	Status int       `json:"status"`
	Expire time.Time `json:"expire"`
	Token  string    `json:"token"`
}

func (u *User) GetToken() error {
	url := config.BackendURL + constants.LoginURL

	loginData := loginData{
		Id:       u.ID,
		Username: u.Name,
	}
	payload, err := getLoginJSON(&loginData)
	if err != nil {
		return err
	}
	log.Print(payload)
	response, err := http.Post(url, "application/json", payload)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New("smth gone wrong :(\nStatus code: " + strconv.Itoa(response.StatusCode))
	}
	var result responseLoginData
	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return err
	}
	u.AuthToken = result.Token
	return nil
}

func (u *User) Register() error {
	url := config.BackendURL + "/v1/users"
	payload, err := json.Marshal(u)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(payload)
	response, err := http.Post(url, "application/json", body)
	if response.StatusCode != 201 {
		return errors.New("smth gone wrong :(\nStatus code: " + strconv.Itoa(response.StatusCode))
	}
	return nil
}
