package routes

import (
	"github.com/GuilleFB/go-party/handlers"
	"github.com/gorilla/mux"
)

func RoutesUsers(r *mux.Router)  {
	r.HandleFunc("/mux/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/mux/user/{id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/mux/user/edit/{id}", handlers.EditUserHandler).Methods("PATCH")
	r.HandleFunc("/mux/user/delete/{id}", handlers.DeleteUserHandler).Methods("DELETE")
}

// func RoutesUsers(r *mux.Router)  {
// 	r.Handle(http.MethodGet, "/mux/users", handlers.GetUsersHandler)
// 	r.handle(http.MethodGet, "/mux/user/{id}", handlers.GetUserHandler)
// 	r.handle(http.MethodGet, "/mux/user/edit/{id}", handlers.EditUserHandler)
// 	r.handle(http.MethodGet, "/mux/user/delete/{id}", handlers.DeleteUserHandler)
// }