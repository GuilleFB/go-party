package routes

import (
	"github.com/GuilleFB/go-party/handlers"
	"github.com/gorilla/mux"
)

func RoutesUsers(r *mux.Router)  {
	r.HandleFunc("/mux/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/create", handlers.PostUserHandler).Methods("POST")
	r.HandleFunc("/user/edit/{id}", handlers.EditUserHandler).Methods("PATCH")
	r.HandleFunc("/user/delete/{id}", handlers.DeleteUserHandler).Methods("DELETE")
}