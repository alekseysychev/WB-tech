package dto

import "time"

type Event struct {
	Id     int       `json:"id"`
	Text   string    `json:"text"`
	Date   time.Time `json:"date"`
	UserId int       `json:"user_id"`
}
