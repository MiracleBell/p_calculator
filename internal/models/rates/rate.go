package rates

import (
	. "../projects"
	. "../teams"
)

type Rate struct {
	ID            uint64
	Position      Position
	RateInDollars float64
	Project       Project
}
