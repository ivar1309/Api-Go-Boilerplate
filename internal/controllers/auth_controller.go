package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ivar1309/Api-Go-Boilerplate/internal/models"
	"github.com/ivar1309/Api-Go-Boilerplate/internal/utils"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	if _, exists := models.Users[req.Username]; exists {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	hashedPwd, _ := utils.HashPassword(req.Password)
	user := models.User{Username: req.Username, Password: hashedPwd, Role: "user"}
	models.Users[req.Username] = user

	json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	user, exists := models.Users[req.Username]
	if !exists || !utils.CheckPassword(req.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(user.Username, user.Role, time.Hour*24)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
