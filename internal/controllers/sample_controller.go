package controllers

import (
	"encoding/json"
	"net/http"
)

func PublicEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "This route is public!"})
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("username").(string)
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome " + user + ", you're authenticated!"})
}

func AdminOnlyEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello Admin! You have access to this restricted route.",
	})
}
