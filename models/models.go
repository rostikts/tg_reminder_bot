package models

type User interface {
	GetToken() error
	Register() error
}

type Reminder struct {
	ID        uint   `json:"id"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	UserId    int    `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
