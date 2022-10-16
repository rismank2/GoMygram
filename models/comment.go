package models

import "time"

type Commment struct {
	Id        int       `json:"id"`
	User_id   int       `json:"user_id"`
	Photo_id  int       `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ResponsePostComment struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	Photo_id  int       `json:"photo_id"`
	User_id   int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseGetComment struct {
	Id        int                     `json:"id"`
	Message   string                  `json:"message"`
	Photo_id  int                     `json:"photo_id"`
	User_id   int                     `json:"user_id"`
	UpdatedAt time.Time               `json:"updated_at"`
	CreatedAt time.Time               `json:"created_at"`
	User      ResponseCommentGetUser  `json:"User"`
	Photo     ResponseCommentGetPhoto `json:"Photo"`
}

type ResponseCommentGetUser struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ResponseCommentGetPhoto struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Caption string `json:"caption"`
	Url     string `json:"photo_url"`
	User_id int    `json:"user_id"`
}

type ResponsePutComment struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"photo_url"`
	User_id   int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
