package teams

import (
	. "../positions"
	. "../projects"
)

type TeamMember struct {
	ID                  uint64
	Position            Position
	DegreeOfInvolvement float64
	Project             Project
}

type Team struct {
	Teams []TeamMember
}
