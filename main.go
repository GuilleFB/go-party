package main

import (
	"net/http"

	"github.com/GuilleFB/go-party/db"
	"github.com/GuilleFB/go-party/models"
	"github.com/GuilleFB/go-party/routes"
	"github.com/gorilla/mux"
)

func main()  {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r:=	mux.NewRouter()
	
	r.HandleFunc("/", routes.HomeHandler).Methods("GET")
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/create", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/user/delete/{id}", routes.DeleteUserHandler).Methods("DELETE")
	
	http.ListenAndServe(":8000",r)
}