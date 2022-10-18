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

type CommentHandlerInterf interface {
	Comment(w http.ResponseWriter, r *http.Request)
}

func NewComment(db *sql.DB) CommentHandlerInterf {
	return &CommentHand{db: db}
}

type CommentHand struct {
	db *sql.DB
}

// Comment implements CommentHandlerInterf
func (ch *CommentHand) Comment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	ctx := r.Context()
	user := middleware.RunUser(ctx)

	fmt.Println(user)
	fmt.Println(user.Id)
	commentservic := service.CommentServic()
	switch r.Method {
	case http.MethodGet:
		fmt.Println("Get Comments")
		comments := repository.CommentGetRepository(ch.db)
		jsonData, _ := json.Marshal(&comments)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(jsonData)

	case http.MethodPost:
		fmt.Println("POST")
		var comment models.Commment
		json.NewDecoder(r.Body).Decode(&comment)
		err := commentservic.CekPostComment(comment.Message)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		} else {
			user_id := user.Id
			repo := repository.CommentPostRepository(ch.db, comment, user_id)
			jsonData, _ := json.Marshal(&repo)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(201)
			w.Write(jsonData)

		}

	case http.MethodPut:
		fmt.Println("PUT")
		var newComment models.Commment
		json.NewDecoder(r.Body).Decode(&newComment)
		err := commentservic.CekPostComment(newComment.Message)
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
		} else {
			fmt.Print("test isi comment")
			if id != "" {
				repo := repository.CommentPutRepository(ch.db, newComment, id)
				jsonData, _ := json.Marshal(&repo)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(200)
				w.Write(jsonData)
			} else {
				err = errors.New("id not empty")
				w.Write([]byte(fmt.Sprint(err)))
			}
		}
	case http.MethodDelete:
		fmt.Println("DELETE")
		if id != "" {
			message := repository.CommentDeleteRepository(ch.db, id)
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
