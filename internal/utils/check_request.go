package utils

import "net/http"

func IsPost(r *http.Request) bool {
	return r.Method == "POST"
}

func IsGet(r *http.Request) bool {
	return r.Method == "GET"
}

func IsPut(r *http.Request) bool {
	return r.Method == "PUT"
}

func IsDelete(r *http.Request) bool {
	return r.Method == "DELETE"
}
