package handlers

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

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)
	userID := params["id"]

	db.DB.First(&user, userID)

	if user.ID == 0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	
	json.NewDecoder(r.Body).Decode(&user)

	// db.DB.Save(&user) // Save is a combination function. If save value does not contain primary key, it will execute Create, otherwise it will execute Update (with all fields).
	db.DB.Model(&user).Updates(user) // https://gorm.io/docs/update.html#Updates-multiple-columns

	json.NewEncoder(w).Encode(user)

	w.WriteHeader(http.StatusOK)
    w.Write([]byte("User successfully updated"))
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

	db.DB.Unscoped().Delete(&user) // Hard delete https://gorm.io/docs/delete.html#Delete-permanently
	// db.DB.Delete(&user) // Soft delete https://gorm.io/docs/delete.html#Soft-Delete

	w.WriteHeader(http.StatusOK)
}