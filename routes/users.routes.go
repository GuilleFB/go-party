package routes

import (
	"encoding/json"
	"net/http"

	"github.com/GuilleFB/go-party/db"
	"github.com/GuilleFB/go-party/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request)  {
	var users []models.User

	db.DB.Find(&users)
	
	json.NewEncoder(w).Encode(&users)
	
}

func GetUserHandler(w http.ResponseWriter, r *http.Request)  {
	var user models.User

	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request)  {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)

	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
	
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request)  {
	var user models.User

	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Unscoped().Delete(&user) // hard delete
	// db.DB.Delete(&user) // soft delete

	w.WriteHeader(http.StatusOK)
}