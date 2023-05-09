package projects

import (
	handler "../../handlers/projects"
	"net/http"
)

func ProjectRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		AddNewProject(w, r)
	case "GET":
		GetProjectList(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func AddNewProject(w http.ResponseWriter, r *http.Request) {
	handler.PostMilestone(w, r)
}

func GetProjectList(w http.ResponseWriter, r *http.Request) {
	handler.GetMilestone(w, r)
}
