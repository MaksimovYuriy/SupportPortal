package models

import "time"

type Ticket struct {
	ID          int64     `json:"id"`
	QueueID     *int64    `json:"queue_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
