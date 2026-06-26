package models

import "time"

type Queue struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_available"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
