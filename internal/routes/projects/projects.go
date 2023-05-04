package projects

import (
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
	//TODO: do someth
}

func GetProjectList(w http.ResponseWriter, r *http.Request) {

}
