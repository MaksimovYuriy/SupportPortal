package models

import "time"

type Agent struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	IsAvailable bool      `json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
}
