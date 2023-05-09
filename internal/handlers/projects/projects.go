package projects

import (
	. "../../models/projects"
	service "../../services/projects"
	"encoding/json"
	"net/http"
)

func PostMilestone(w http.ResponseWriter, r *http.Request) {
	var project Project
	json.NewDecoder(r.Body).Decode(&project)
	err := service.AddProject(&project)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(project)
}

func GetMilestone(w http.ResponseWriter, r *http.Request) {
	projects, err := service.GetProjects()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(projects)
}
