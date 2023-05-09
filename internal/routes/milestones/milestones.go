package milestones

import (
	handler "../../handlers/milestones"
	"net/http"
)

func MilestoneRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		AddNewMilestone(w, r)
	case "GET":
		GetProjectMilestones(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func MilestoneRouterWithId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		PutMilestone(w, r)
	case "DELETE":
		DeleteMilestone(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func AddNewMilestone(w http.ResponseWriter, r *http.Request) {
	handler.PostMilestone(w, r)
}

func GetProjectMilestones(w http.ResponseWriter, r *http.Request) {
	handler.GetMilestone(w, r)
}

func PutMilestone(w http.ResponseWriter, r *http.Request) {
	handler.PutMilestone(w, r)
}

func DeleteMilestone(w http.ResponseWriter, r *http.Request) {
	handler.DeleteMilestone(w, r)
}
