package rates

import (
	. ".."
	. "../../models/rates"
	"log"
)

var db, _ = OpenConnection()

func AddRate(rate *Rate) error {
	query := "INSERT INTO rates(positions, rate_in_dollars, project_id) VALUES(?,?,?)"

	result, err := db.Exec(query,
		rate.Position.Name,
		rate.RateInDollars,
		rate.Project.ID)

	if err != nil {
		log.Fatal("Can't add rate in DB")
		return err
	}
	id, _ := result.LastInsertId()
	rate.ID = uint64(id)
	return nil
}

func ChangeRate(rate *Rate) error {
	query := "UPDATE milestones SET rate_in_dollars=$1 VALUES(?)"

	_, err := db.Exec(query,
		rate.RateInDollars)

	if err != nil {
		log.Fatal("Can't update project rates in DB")
		return err
	}
	return nil
}

func GetProjectRates(projectId uint64) ([]Rate, error) {
	query := "SELECT id, positions, rate_in_dollars, project_id  FROM rates"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Can't get project rates")
		return nil, err
	}
	defer rows.Close()

	var rates []Rate
	for rows.Next() {
		var rate Rate
		err := rows.Scan(
			&rate.ID,
			&rate.Position.Name,
			&rate.RateInDollars,
			&rate.Project.ID)
		if err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, err
}
