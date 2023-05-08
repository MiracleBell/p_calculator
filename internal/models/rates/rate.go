package rates

import (
	. "../positions"
	. "../projects"
)

type Rate struct {
	ID            uint64
	Position      Position
	RateInDollars float64
	Project       Project
}
