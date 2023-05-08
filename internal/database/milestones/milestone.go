package milestones

import (
	. ".."
	. "../../models/milestones"
	"log"
)

var db, _ = OpenConnection()

func AddNewMilestone(milestone *Milestone) error {
	query := "INSERT INTO milestones(title, description, start_date_time, end_date_time, created_at, last_updated_at, project_id) VALUES(?,?,?,?,?,?,?)"

	result, err := db.Exec(query,
		milestone.Title,
		milestone.Description,
		milestone.StartDateTime,
		milestone.EndDateTime,
		milestone.CreatedAt,
		milestone.LastUpdatedAt,
		milestone.Project.ID)

	if err != nil {
		log.Fatal("Can't add milestone in DB")
		return err
	}
	id, _ := result.LastInsertId()
	milestone.ID = uint64(id)
	return nil
}

func ChangeMilestone(milestone *Milestone) error {
	query := "UPDATE milestones SET title=$1, description=$2, start_date_time=$3, end_date_time=$4, last_updated_at=$5 VALUES(?,?,?,?,?)"

	_, err := db.Exec(query,
		milestone.Title,
		milestone.Description,
		milestone.StartDateTime,
		milestone.EndDateTime,
		milestone.LastUpdatedAt)

	if err != nil {
		log.Fatal("Can't update milestone in DB")
		return err
	}
	return nil
}

func GetProjectMilestones(projectId uint64) ([]Milestone, error) {
	query := "SELECT id, title, description, start_date_time, end_date_time, created_at, last_updated_at  FROM milestones"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Can't get milestones")
		return nil, err
	}
	defer rows.Close()

	var milestones []Milestone
	for rows.Next() {
		var milestone Milestone
		err := rows.Scan(
			&milestone.ID,
			&milestone.Title,
			&milestone.Description,
			&milestone.StartDateTime,
			&milestone.EndDateTime,
			&milestone.CreatedAt,
			&milestone.LastUpdatedAt)
		if err != nil {
			return nil, err
		}
		milestones = append(milestones, milestone)
	}
	return milestones, err
}

func DeleteMilestoneById(id uint64) error {
	deleteQuery := "DELETE FROM milestone WHERE id=?"
	_, err := db.Query(deleteQuery, id)
	if err != nil {
		log.Fatal("Can't remove milestone")
	}
	return nil
}
