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

type PhotoHand struct {
	db *sql.DB
}

// Photo implements PhotoHandlerInterf
func (ph *PhotoHand) Photo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	ctx := r.Context()
	user := middleware.RunUser(ctx)

	fmt.Println(user)
	fmt.Println(user.Id)
	switch r.Method {
	case http.MethodGet:
		fmt.Println("Get Photo")
		photos := repository.PhotoGetRepo(ph.db)
		jsonData, _ := json.Marshal(&photos)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonData)

	case http.MethodPost:
		fmt.Println("Post Photo")

		//penampungan rbody
		var newPhotos models.Photo
		json.NewDecoder(r.Body).Decode(&newPhotos)
		fmt.Println(newPhotos)
		//check validasi user
		photoserv := service.NewPhotoService()
		err := photoserv.CekPostPhoto(newPhotos.Title, newPhotos.Url)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		} else {
			//query insert
			user_id := user.Id
			response := repository.PhotoPostRepo(ph.db, newPhotos, user_id)
			jsonData, _ := json.Marshal(&response)
			w.Header().Add("Content-Type", "application/json")
			w.Write(jsonData)
			w.WriteHeader(201)
		}

	case http.MethodPut:
		fmt.Println("Put")
		if id != "" {

			var newPhotos models.Photo
			json.NewDecoder(r.Body).Decode(&newPhotos)
			photoserv := service.NewPhotoService()
			err := photoserv.CekPostPhoto(newPhotos.Title, newPhotos.Url)
			if err != nil {
				w.Write([]byte(fmt.Sprint(err)))
			} else {
				response := repository.PhotoPutRepo(ph.db, newPhotos, id)
				jsonData, _ := json.Marshal(&response)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(200)
				w.Write(jsonData)
			}

		} else {
			err := errors.New("Photo not found")
			w.Write([]byte(fmt.Sprint(err)))
		}
	case http.MethodDelete:
		fmt.Println("DELETE")
		if id != "" {
			message := repository.PhotoDeleteRepo(ph.db, id)
			jsonData, _ := json.Marshal(&message)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(jsonData)
		} else {
			err := errors.New("id not empty")
			w.Write([]byte(fmt.Sprint(err)))
		}
	}

}

type PhotoHandlerInterf interface {
	Photo(w http.ResponseWriter, r *http.Request)
}

func NewPhoto(db *sql.DB) PhotoHandlerInterf {
	return &PhotoHand{db: db}
}
