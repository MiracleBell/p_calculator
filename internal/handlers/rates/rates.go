package rates

import (
	. "../../models/rates"
	service "../../services/rates"
	"encoding/json"
	"net/http"
)

func PostRate(w http.ResponseWriter, r *http.Request) {
	var rate Rate
	json.NewDecoder(r.Body).Decode(&rate)
	err := service.AddRate(&rate)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(rate)
}

func PutRate(w http.ResponseWriter, r *http.Request) {
	var rate Rate
	json.NewDecoder(r.Body).Decode(&rate)
	err := service.ChangeRate(&rate)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(rate)
}

func GetRates(w http.ResponseWriter, r *http.Request) {
	var projectId uint64
	json.NewDecoder(r.Body).Decode(&projectId)
	rates, err := service.GetProjectRates(projectId)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(rates)
}
