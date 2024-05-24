package db

import (
	"log"

	"github.com/GuilleFB/go-party/initializers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection()  {
	var error error
	DB, error = gorm.Open(postgres.Open(initializers.VariablesDB()), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else{
		log.Println("Data Base Connected")
	}
}