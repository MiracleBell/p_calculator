package features

import (
	. "../../utils"
	"net/http"
)

func FeaturesRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		AddNewFeature(w, r)
	case "GET":
		GetProjectsFeatures(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func AddNewFeature(w http.ResponseWriter, r *http.Request) {
	//TODO: do someth
}

func GetProjectsFeatures(w http.ResponseWriter, r *http.Request) {

}

func PutFeature(w http.ResponseWriter, r *http.Request) {
	if IsPut(r) {

	}
	http.Error(w, "Invalid request method", 405)
}
