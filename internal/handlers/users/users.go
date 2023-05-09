package users

import (
	. "../../models/users"
	service "../../services/users"
	"encoding/json"
	"net/http"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	err := service.AddUser(&user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(user)
}
