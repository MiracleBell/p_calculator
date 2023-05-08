package projects

import (
	. ".."
	. "../../models/projects"
	"log"
)

var db, _ = OpenConnection()

func AddNewProject(project *Project) error {
	query := "INSERT INTO projects(title, description, client, creator_id, created_at, last_updated_at) VALUES(?,?,?,?,?,?)"

	result, err := db.Exec(query,
		project.Title,
		project.Description,
		project.Client,
		project.Creator.ID,
		project.CreatedAt,
		project.LastUpdatedAt)

	if err != nil {
		log.Fatal("Can't add project in DB")
		return err
	}
	id, _ := result.LastInsertId()
	project.ID = uint64(id)
	return nil
}

func GetProjects() ([]Project, error) {
	query := "SELECT id, title, description, client, creator_id, created_at, last_updated_at  FROM projects"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Can't get projects")
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		err := rows.Scan(
			&project.ID,
			&project.Title,
			&project.Description,
			&project.Client,
			&project.Creator.ID,
			&project.CreatedAt,
			&project.LastUpdatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, err
}
