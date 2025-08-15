package domain

import "time"

type User struct {
	ID         int
	Name       string
	TelegramID int64
	CreatedAt  time.Time
}
