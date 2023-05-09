package teams

import (
	handler "../../handlers/teams"
	"net/http"
)

func TeamRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		PutProjectTeam(w, r)
	case "GET":
		GetProjectTeam(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func GetProjectTeam(w http.ResponseWriter, r *http.Request) {
	handler.GetMilestone(w, r)
}

func PutProjectTeam(w http.ResponseWriter, r *http.Request) {
	handler.PutMilestone(w, r)
}
