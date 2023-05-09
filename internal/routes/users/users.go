package users

import (
	handler "../../handlers/users"
	"net/http"
)

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handler.PostUser(w, r)
	}
	http.Error(w, "Invalid request method", 405)
}
