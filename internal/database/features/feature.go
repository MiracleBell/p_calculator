package features

import (
	. "../"
	. "../../models/features"
	"log"
)

var db, _ = OpenConnection()

func AddNewFeature(feature *Feature) error {
	query := "INSERT INTO features(title, description, best_case_estimate, most_likely_estimate, worst_case_estimate, created_at, last_updated_at, project_id, milestone_id) VALUES(?,?,?,?,?,?,?,?,?)"

	result, err := db.Exec(query,
		feature.Title,
		feature.Description,
		feature.BestCaseEstimateInDays,
		feature.MostLikelyEstimateInDays,
		feature.WorstCaseEstimateInDays,
		feature.CreatedAt,
		feature.LastUpdatedAt,
		feature.Project.ID,
		feature.Milestone.ID)

	if err != nil {
		log.Fatal("Can't add feature in DB")
		return err
	}
	id, _ := result.LastInsertId()
	feature.ID = uint64(id)
	return nil
}

func ChangeFeature(feature *Feature) error {
	query := "UPDATE features SET title=$1, description=$2, best_case_estimate=$3, most_likely_estimate=$4, worst_case_estimate=$5, last_updated_at=$6, milestone_id=$7 VALUES(?,?,?,?,?,?,?)"

	_, err := db.Exec(query,
		feature.Title,
		feature.Description,
		feature.BestCaseEstimateInDays,
		feature.MostLikelyEstimateInDays,
		feature.WorstCaseEstimateInDays,
		feature.LastUpdatedAt,
		feature.Milestone.ID)

	if err != nil {
		log.Fatal("Can't update feature in DB")
		return err
	}
	return nil
}

func GetProjectFeatures(projectId uint64) ([]Feature, error) {
	query := "SELECT id, title, description, best_case_estimate, most_likely_estimate, worst_case_estimate, created_at, last_updated_at  FROM features"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Can't get features")
		return nil, err
	}
	defer rows.Close()

	var features []Feature
	for rows.Next() {
		var feature Feature
		err := rows.Scan(
			&feature.ID,
			&feature.Title,
			&feature.Description,
			&feature.BestCaseEstimateInDays,
			&feature.MostLikelyEstimateInDays,
			&feature.WorstCaseEstimateInDays,
			&feature.CreatedAt,
			&feature.LastUpdatedAt)
		if err != nil {
			return nil, err
		}
		features = append(features, feature)
	}
	return features, err
}
