package teams

import (
	db "../../database/teams"
	. "../../models/teams"
)

func AddTeamMember(member *TeamMember) error {
	return db.AddTeamMember(member)
}

func UpdateTeamMember(member *TeamMember) error {
	return db.ChangeTeamMember(member)
}

func GetProjectTeam(projectId uint64) ([]TeamMember, error) {
	return db.GetProjectTeam(projectId)
}
