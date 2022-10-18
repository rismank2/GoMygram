package controller

import (
	"MyGram/middleware"
	"MyGram/models"
	"MyGram/repository"
	"MyGram/service"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type SosialMediaHandlerInterf interface {
	SosialMedia(w http.ResponseWriter, r *http.Request)
}

func NewSosialMedia(db *sql.DB) SosialMediaHandlerInterf {
	return &SosialMediaHand{db: db}
}

type SosialMediaHand struct {
	db *sql.DB
}

// SosialMedia implements SosialMediaHandlerInterf
func (sm *SosialMediaHand) SosialMedia(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)

	ctx := r.Context()
	user := middleware.RunUser(ctx)

	fmt.Println(user)
	fmt.Println(user.Id)
	sosialmediaser := service.SocialMediaSerN()

	switch r.Method {
	case http.MethodGet:
		fmt.Println("Get Social Media")
		socialmedias := repository.SocmedGetRepo(sm.db)
		jsonData, _ := json.Marshal(&socialmedias)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonData)

	case http.MethodPost:
		fmt.Println("POST")
		var newSocialMedia models.SocialMedia
		json.NewDecoder(r.Body).Decode(&newSocialMedia)
		err := sosialmediaser.CekPostSocialMedia(newSocialMedia.Name, newSocialMedia.Social_Media_Url)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		} else {

			user_id := user.Id
			response := repository.SocmedPostRepo(sm.db, newSocialMedia, user_id)
			jsonData, _ := json.Marshal(&response)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(201)
			w.Write(jsonData)
		}

	case http.MethodPut:
		fmt.Println("PUT")
		if id != "" {
			var newSocialMedia models.SocialMedia
			json.NewDecoder(r.Body).Decode(&newSocialMedia)
			err := sosialmediaser.CekPostSocialMedia(newSocialMedia.Name, newSocialMedia.Social_Media_Url)
			if err != nil {
				w.Write([]byte(fmt.Sprint(err)))
			} else {
				response := repository.SocmedPutRepo(sm.db, newSocialMedia, id)
				jsonData, _ := json.Marshal(&response)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(200)
				w.Write(jsonData)
			}

		}

	case http.MethodDelete:
		fmt.Println("DELETE")
		if id != "" {
			message := repository.SocmedDelRepo(sm.db, id)
			jsonData, _ := json.Marshal(&message)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(jsonData)
		} else {
			err := errors.New("id is empty")
			w.Write([]byte(fmt.Sprint(err)))
		}

	}

}
