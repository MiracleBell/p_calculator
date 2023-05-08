package teams

import (
	. ".."
	. "../../models/teams"
	"log"
)

var db, _ = OpenConnection()

func AddTeamMember(teamMember *TeamMember) error {
	query := "INSERT INTO team_members(positions, degree_of_involvement, project_id) VALUES(?,?,?)"

	result, err := db.Exec(query,
		teamMember.Position.Name,
		teamMember.DegreeOfInvolvement,
		teamMember.Project.ID)

	if err != nil {
		log.Fatal("Can't add teamMember in DB")
		return err
	}
	id, _ := result.LastInsertId()
	teamMember.ID = uint64(id)
	return nil
}

func ChangeTeamMember(teamMember *TeamMember) error {
	query := "UPDATE team_members SET degree_of_involvement=$1 VALUES(?)"

	_, err := db.Exec(query,
		teamMember.DegreeOfInvolvement)

	if err != nil {
		log.Fatal("Can't update team member in DB")
		return err
	}
	return nil
}

func GetProjectTeam(projectId uint64) ([]TeamMember, error) {
	query := "SELECT id, positions, degree_of_involvement, project_id  FROM team_members"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Can't get project team")
		return nil, err
	}
	defer rows.Close()

	var team []TeamMember
	for rows.Next() {
		var teamMember TeamMember
		err := rows.Scan(
			&teamMember.ID,
			&teamMember.Position.Name,
			&teamMember.DegreeOfInvolvement,
			&teamMember.Project.ID)
		if err != nil {
			return nil, err
		}
		team = append(team, teamMember)
	}
	return team, err
}
