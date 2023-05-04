package projects

import (
	. "../users"
	"time"
)

type Project struct {
	ID            uint64
	Title         string
	Description   string
	Client        string
	Creator       User
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
