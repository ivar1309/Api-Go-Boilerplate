package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ivar1309/Api-Go-Boilerplate/internal/models"
	"github.com/ivar1309/Api-Go-Boilerplate/internal/utils"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Message string       `json:"message"`
	Tokens  utils.Tokens `json:"tokens"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	_, err := models.GetUserByUsername(req.Username)
	if err == nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	hashedPwd, _ := utils.HashPassword(req.Password)
	user := models.User{
		Username:     req.Username,
		PasswordHash: hashedPwd,
		Role:         "user",
	}

	err = models.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	tokens, _ := utils.GenerateTokens(user.Username, user.Role)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{Message: "User created", Tokens: tokens})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	json.NewDecoder(r.Body).Decode(&req)

	user, err := models.GetUserByUsername(req.Username)
	if err != nil {
		http.Error(w, "No such user", http.StatusNotFound)
		return
	}
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokens, _ := utils.GenerateTokens(user.Username, user.Role)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{Message: "Login successful", Tokens: tokens})
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	var body struct {
		RefreshToken string `json:"refreshToken"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	claims, err := utils.ValidateToken(body.RefreshToken, true)
	if err != nil {
		http.Error(w, "Invalid refresh", http.StatusUnauthorized)
		return
	}

	tokens, _ := utils.GenerateTokens(claims.Username, claims.Role)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{Message: "Refresh successful", Tokens: tokens})
}
