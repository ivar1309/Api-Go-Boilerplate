package models

type User struct {
	Username string
	Password string
	Role     string
}

var Users = map[string]User{}
