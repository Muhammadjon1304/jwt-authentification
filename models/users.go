package models

import "time"

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
}

type Login_User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
