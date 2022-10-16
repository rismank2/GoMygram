package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseRegister struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type ResponseUpdate struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}
