package features

import (
	. "../milestones"
	. "../projects"
	"time"
)

type Feature struct {
	ID                       uint64
	Title                    string
	Description              string
	BestCaseEstimateInDays   time.Duration
	MostLikelyEstimateInDays time.Duration
	WorstCaseEstimateInDays  time.Duration
	CreatedAt                time.Time
	LastUpdatedAt            time.Time
	Project                  Project
	Milestone                Milestone
}
