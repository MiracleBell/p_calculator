package positions

import (
	. ".."
	. "../../models/positions"
	"log"
)

var db, _ = OpenConnection()

func GetPositions() ([]Position, error) {
	query := "SELECT name, defaultRate  FROM positions"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Can't get positions")
		return nil, err
	}
	defer rows.Close()

	var positions []Position
	for rows.Next() {
		var position Position
		err := rows.Scan(
			&position.Name,
			&position.DefaultRate)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, err
}

func GetByName(name string) (Position, error) {
	query := "SELECT name, defaultRate  FROM positions WHERE name LIKE ?"

	var position Position
	err := db.QueryRow(query, name).Scan(
		&position.Name,
		&position.DefaultRate)

	if err != nil {
		log.Fatal("Can't get position by name")
		return Position{}, err
	}

	return position, err
}
