package milestones

import (
	. "../../models/milestones"
	service "../../services/milestones"
	"encoding/json"
	"net/http"
)

func PostMilestone(w http.ResponseWriter, r *http.Request) {
	var milestone Milestone
	json.NewDecoder(r.Body).Decode(&milestone)
	err := service.AddMilestone(&milestone)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(milestone)
}

func PutMilestone(w http.ResponseWriter, r *http.Request) {
	var milestone Milestone
	json.NewDecoder(r.Body).Decode(&milestone)
	err := service.ChangeMilestone(&milestone)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(milestone)
}

func GetMilestone(w http.ResponseWriter, r *http.Request) {
	var id uint64
	json.NewDecoder(r.Body).Decode(&id)
	milestones, err := service.GetMilestones(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(milestones)
}

func DeleteMilestone(w http.ResponseWriter, r *http.Request) {
	var id uint64
	json.NewDecoder(r.Body).Decode(&id)
	err := service.DeleteMilestone(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
}
