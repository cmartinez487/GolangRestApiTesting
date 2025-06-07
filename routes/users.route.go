package routes

import (
	"encoding/json"
	"net/http"

	"github.com/cmartinez/GolangRestApiTesting/db"
	"github.com/cmartinez/GolangRestApiTesting/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	if db.DB.Error != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error retrieving users: " + db.DB.Error.Error()))
		return
	}
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404 Not Found
		w.Write([]byte("User not found"))
		return
	} else if db.DB.Error != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error retrieving user: " + db.DB.Error.Error()))
		return
	}

	db.DB.Model(&user).Association("Task").Find(&user.Task) // Load associated tasks

	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error creating user: " + db.DB.Error.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(&user)
}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404 Not Found
		w.Write([]byte("User not found"))
		return
	} else if db.DB.Error != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error retrieving user: " + db.DB.Error.Error()))
		return
	}

	var updatedData models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON format"))
		return
	}

	// Realizar la actualización con GORM
	if err := db.DB.Model(&user).Updates(updatedData).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
		w.Write([]byte("Error updating user: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404 Not Found
		w.Write([]byte("User not found"))
		return
	} else if db.DB.Error != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error retrieving user: " + db.DB.Error.Error()))
		return
	}

	//db.DB.Unscoped().Delete(&user) // Use Unscoped to delete the user permanently
	db.DB.Delete(&user, params["id"])

	w.WriteHeader(http.StatusOK)
}
