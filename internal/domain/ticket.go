package domain

import "time"

type Ticket struct {
	ID          int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
