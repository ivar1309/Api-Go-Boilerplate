package models

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ivar1309/Api-Go-Boilerplate/internal/db"
)

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
}

var Users = map[string]User{}

var ErrUserExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user not found")

func GetUserByUsername(username string) (User, error) {
	var u User
	query := `
		SELECT id, username, password_hash, role 
		FROM users 
		WHERE username=$1
	`

	err := db.GetDB().QueryRow(context.Background(), query, username).
		Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Role)

	if err == sql.ErrNoRows {
		return u, ErrUserNotFound
	}

	return u, err
}

func CreateUser(user User) error {
	query := `
		INSERT INTO users (username, password_hash, role) 
		VALUES ($1, $2, $3)
	`

	_, err := db.GetDB().Exec(context.Background(), query, user.Username, user.PasswordHash, user.Role)

	return err
}
