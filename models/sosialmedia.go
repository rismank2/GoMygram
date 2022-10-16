package models

import "time"

type SocialMedia struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Social_Media_Url string `json:"social_media_url"`
	User_id          int    `json:"user_id"`
}

type ResponsePostSocialMedia struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Social_Media_Url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
}

type ResponsePutSocialMedia struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Social_Media_Url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type SocialMedias struct {
	SocialMedia ResponseGetSocialMedia `json:"social_media"`
}

type ResponseGetSocialMedia struct {
	Id               int                        `json:"id"`
	Name             string                     `json:"name"`
	Social_Media_Url string                     `json:"social_media_url"`
	User_id          int                        `json:"UserId"`
	CreatedAt        time.Time                  `json:"createdAt"`
	UpdatedAt        time.Time                  `json:"updatedAt"`
	User             ResponseUserGetSocialMedia `json:"User"`
}

type ResponseUserGetSocialMedia struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Url      string `json:"profile_image_url"`
}
