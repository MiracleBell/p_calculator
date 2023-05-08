package positions

import (
	db "../../database/positions"
	. "../../models/positions"
)

func GetPositions() ([]Position, error) {
	return db.GetPositions()
}

func GetPositionByName(name string) (Position, error) {
	return db.GetByName(name)
}
