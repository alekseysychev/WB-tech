package dto

import "time"

type Event struct {
	Id   int       `json:"id"`
	User int       `json:"user_id"`
	Name string    `json:"name"`
	Text string    `json:"text"`
	Date time.Time `json:"date"`
}
