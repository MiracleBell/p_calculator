package milestones

import (
	. "../../utils"
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
	if IsPost(r) {
		//TODO: do someth
	}
	http.Error(w, "Invalid request method", 405)
}

func GetProjectMilestones(w http.ResponseWriter, r *http.Request) {
	if IsGet(r) {

	}
	http.Error(w, "Invalid request method", 405)
}

func PutMilestone(w http.ResponseWriter, r *http.Request) {
	if IsPut(r) {

	}
	http.Error(w, "Invalid request method", 405)
}

func DeleteMilestone(w http.ResponseWriter, r *http.Request) {
	if IsDelete(r) {

	}
	http.Error(w, "Invalid request method", 405)
}
