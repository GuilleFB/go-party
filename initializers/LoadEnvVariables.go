package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables()  {
	var err = godotenv.Load()

	if err != nil{
		log.Fatal("Error loading .env file")
	} else {
		log.Println("Environment Variables Loaded")
	}
}

func VariablesDB() string {
	host := os.Getenv("POSTGRES_HOST")
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DB")
    port := os.Getenv("POSTGRES_PORT")

    // Construir la cadena DSN usando fmt.Sprintf
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)
	return dsn
}