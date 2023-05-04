package rates

import (
	"net/http"
)

func RatesRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		PutProjectRates(w, r)
	case "GET":
		GetProjectRates(w, r)
	default:
		http.Error(w, "Invalid request method", 405)
	}
}

func GetProjectRates(w http.ResponseWriter, r *http.Request) {
	//TODO: do someth
}

func PutProjectRates(w http.ResponseWriter, r *http.Request) {

}
