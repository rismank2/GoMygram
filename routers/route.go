package router

import (
	controller "MyGram/controllers"
	"MyGram/middleware"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// handler user

const PORT = ":4828"

func RunRoute(db *sql.DB) {
	route := mux.NewRouter()
	usersHandler := controller.NewUsersHandler(db)
	registerHandler := controller.NewRegisterHandler(db)
	loginHandler := controller.UserLoginHandler(db)

	route.HandleFunc("/users/register", registerHandler.Register)
	route.HandleFunc("/users/login", loginHandler.Login)
	route.Handle("/users/{id}", middleware.Auth(http.HandlerFunc(usersHandler.UsersHandler))).Methods("PUT")
	route.Handle("/users", middleware.Auth(http.HandlerFunc(usersHandler.UsersHandler))).Methods("DELETE")

	//handler photo
	photoHandler := controller.NewPhoto(db)
	route.Handle("/photos", middleware.Auth(http.HandlerFunc(photoHandler.Photo))).Methods("GET")
	route.Handle("/photos", middleware.Auth(http.HandlerFunc(photoHandler.Photo))).Methods("POST")
	route.Handle("/photos/{id}", middleware.Auth(http.HandlerFunc(photoHandler.Photo))).Methods("PUT")
	route.Handle("/photos/{id}", middleware.Auth(http.HandlerFunc(photoHandler.Photo))).Methods("DELETE")

	//handler comment
	commentHandler := controller.NewComment(db)
	route.Handle("/comments", middleware.Auth(http.HandlerFunc(commentHandler.Comment))).Methods("GET")
	route.Handle("/comments", middleware.Auth(http.HandlerFunc(commentHandler.Comment))).Methods("POST")
	route.Handle("/comments/{id}", middleware.Auth(http.HandlerFunc(commentHandler.Comment))).Methods("PUT")
	route.Handle("/comments/{id}", middleware.Auth(http.HandlerFunc(commentHandler.Comment))).Methods("DELETE")

	//handler comment
	sosialmediaHandler := controller.NewSosialMedia(db)
	route.Handle("/sosialmedias", middleware.Auth(http.HandlerFunc(sosialmediaHandler.SosialMedia))).Methods("GET")
	route.Handle("/sosialmedias", middleware.Auth(http.HandlerFunc(sosialmediaHandler.SosialMedia))).Methods("POST")
	route.Handle("/sosialmedias/{id}", middleware.Auth(http.HandlerFunc(sosialmediaHandler.SosialMedia))).Methods("PUT")
	route.Handle("/sosialmedias/{id}", middleware.Auth(http.HandlerFunc(sosialmediaHandler.SosialMedia))).Methods("DELETE")
	log.Println("Server aktif di http://127.0.0.1" + PORT)
	log.Println("Tekan CTRL+C untuk keluar")
	srv := &http.Server{
		Handler:      route,
		Addr:         "0.0.0.0" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
