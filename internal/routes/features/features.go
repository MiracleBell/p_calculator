package features

import (
	handler "../../handlers/features"
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
	handler.PostFeature(w, r)
}

func GetProjectsFeatures(w http.ResponseWriter, r *http.Request) {
	handler.GetFeature(w, r)
}

func PutFeature(w http.ResponseWriter, r *http.Request) {
	if IsPut(r) {
		handler.PutFeature(w, r)
	}
	http.Error(w, "Invalid request method", 405)
}
