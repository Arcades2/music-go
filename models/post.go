package models

import (
	"time"
)

type Post struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	Id          int       `json:"id"`
	User        int       `json:"user"`
	Enabled     bool      `json:"enabled"`
}
