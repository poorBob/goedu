package models

import "time"

type Message struct {
	ID       int       `json:"id"`
	Uuid     string    `json:"uuid"`
	DateTime time.Time `json:"dateTime"`
	Content  string    `json:"content"`
}
