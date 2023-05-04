package teams

import (
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
	//TODO: do someth
}

func PutProjectTeam(w http.ResponseWriter, r *http.Request) {

}
