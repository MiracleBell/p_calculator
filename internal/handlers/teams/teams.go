package teams

import (
	. "../../models/teams"
	service "../../services/teams"
	"encoding/json"
	"net/http"
)

func PostTeamMember(w http.ResponseWriter, r *http.Request) {
	var member TeamMember
	json.NewDecoder(r.Body).Decode(&member)
	err := service.AddTeamMember(&member)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(member)
}

func PutMilestone(w http.ResponseWriter, r *http.Request) {
	var member TeamMember
	json.NewDecoder(r.Body).Decode(&member)
	err := service.UpdateTeamMember(&member)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(member)
}

func GetMilestone(w http.ResponseWriter, r *http.Request) {
	var id uint64
	json.NewDecoder(r.Body).Decode(&id)
	members, err := service.GetProjectTeam(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(members)
}
