package teams

import . "../projects"

type Position struct {
	Name        string
	DefaultRate float64
}

type TeamMember struct {
	ID                  uint64
	Position            Position
	DegreeOfInvolvement float64
	Project             Project
}

type Team struct {
	Teams []TeamMember
}
