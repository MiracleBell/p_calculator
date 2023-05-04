package users

import "net/http"

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//TODO: do someth...
	}
	http.Error(w, "Invalid request method", 405)
}
