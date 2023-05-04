package milestones

import (
	. "../projects"
	"time"
)

type Milestone struct {
	ID            uint64
	Title         string
	Description   string
	StartDateTime time.Time
	EndDateTime   time.Time
	CreatedAt     time.Time
	LastUpdatedAt time.Time
	Project       Project
}
