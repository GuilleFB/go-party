package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GuilleFB/go-party/db"
	"github.com/GuilleFB/go-party/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	} else {
		w.WriteHeader(http.StatusFound)
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks) // Displays in the user json all the tasks it has

	json.NewEncoder(w).Encode(&user)
}

// func PostUserHandler(w http.ResponseWriter, r *http.Request) {
func PostUserHandler(c *gin.Context) {
	var UserInput models.User

	// With Mux using net/http
	// if err := json.NewDecoder(r.Body).Decode(&UserInput); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// With Gin
	if err := c.ShouldBindJSON(&UserInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Where("username=?", UserInput.Username).Find(&UserInput)

	// With Mux using net/http
	// if UserInput.ID != 0 {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("username already used"))
	// 	return
	// }

	// With Gin
	if UserInput.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	passwordHash, errpassword := bcrypt.GenerateFromPassword([]byte(UserInput.Password), bcrypt.DefaultCost)

	// With Mux using net/http
	// if errpassword != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(errpassword.Error()))
	// 	return
	// }

	// With Gin
	if errpassword != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errpassword.Error()})
		return
	}

	// With Mux using net/http
	// json.NewDecoder(r.Body).Decode(&UserInput)

	// Overwrite only the necessary fields, keeping first_name and last_name
	UserInput.Password = string(passwordHash)

	// Saves the user in the database
	errcreateuser := db.DB.Create(&UserInput).Error

	// With Gin
	if errcreateuser != nil {
		c.Status(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": UserInput})

	}

	// With Mux using net/http
	// if errcreateuser != nil {
	// 	w.WriteHeader(http.StatusBadRequest) // 400
	// 	w.Write([]byte(errcreateuser.Error()))
	// }
	// json.NewEncoder(w).Encode(&UserInput)

}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)
	userID := params["id"]

	db.DB.First(&user, userID)

	if user.ID == 0 {
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

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Unscoped().Delete(&user) // Hard delete https://gorm.io/docs/delete.html#Delete-permanently
	// db.DB.Delete(&user) // Soft delete https://gorm.io/docs/delete.html#Soft-Delete

	w.WriteHeader(http.StatusOK)
}
