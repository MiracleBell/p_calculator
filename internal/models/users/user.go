package users

import "time"

type User struct {
	ID            uint64
	Login         string
	Password      string
	Email         string
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
