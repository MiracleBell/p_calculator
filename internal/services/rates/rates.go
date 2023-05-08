package rates

import (
	db "../../database/rates"
	. "../../models/rates"
)

func AddRate(rate *Rate) error {
	return db.AddRate(rate)
}

func ChangeRate(rate *Rate) error {
	return db.ChangeRate(rate)
}

func GetProjectRates(projectId uint64) ([]Rate, error) {
	return db.GetProjectRates(projectId)
}
