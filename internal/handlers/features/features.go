package features

import (
	. "../../models/features"
	. "../../services/features"
	"encoding/json"
	"net/http"
)

func PostFeature(w http.ResponseWriter, r *http.Request) {
	var feature Feature
	json.NewDecoder(r.Body).Decode(&feature)
	err := AddFeature(&feature)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(feature)
}

func PutFeature(w http.ResponseWriter, r *http.Request) {
	var feature Feature
	json.NewDecoder(r.Body).Decode(&feature)
	err := UpdateFeature(&feature)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(feature)
}

func GetFeature(w http.ResponseWriter, r *http.Request) {
	var id uint64
	json.NewDecoder(r.Body).Decode(&id)
	features, err := GetFeatures(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(features)
}
