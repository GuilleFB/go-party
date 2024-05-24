package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GuilleFB/go-party/db"
	"github.com/GuilleFB/go-party/initializers"
	"github.com/GuilleFB/go-party/models"
	"github.com/GuilleFB/go-party/routes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func init() {
	initializers.LoadEnvVariables()
	db.DBConnection()
}
func main() {
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	// Utilizando Gin
	rgin := gin.Default()
	routes.RoutesTask(rgin)

	// Utilizando Gorilla Mux
	rmux := mux.NewRouter()
	routes.RoutesUsers(rmux)
	routes.RoutesIndex(rmux)

	mainRouter := http.NewServeMux()
	// Añadir Gin al enrutador principal
	mainRouter.Handle("/gin/", rgin)
	// Añadir Gorilla Mux al enrutador principal
	mainRouter.Handle("/mux/", rmux)

	port := os.Getenv("LISTEN_PORT")
	message := fmt.Sprintf("Listening on port %s", port)
	log.Println(message)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}
