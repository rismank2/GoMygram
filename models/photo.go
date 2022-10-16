package models

import (
	"time"
)

type UserPhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Message struct {
	Message string `json:"message"`
}

type Photo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type ResponsePostPhoto struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseGetPhoto struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     UserPhoto `json:"User"`
}

type ResponsePutPhoto struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
