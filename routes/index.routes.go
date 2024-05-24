package routes

import (
	"github.com/GuilleFB/go-party/handlers"
	"github.com/gorilla/mux"
)

func RoutesIndex(r *mux.Router)  {
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
}