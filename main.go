package main

import (
	"net/http"

	"github.com/GuilleFB/go-party/db"
	"github.com/GuilleFB/go-party/initializers"
	"github.com/GuilleFB/go-party/models"
	"github.com/GuilleFB/go-party/routes"
	"github.com/gorilla/mux"
)

func init()  {
	initializers.LoadEnvVariables()
	db.DBConnection()	
}
func main()  {
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r:=	mux.NewRouter()
	
	routes.RoutesUsers(r)
	
	http.ListenAndServe(":8000",r)
}