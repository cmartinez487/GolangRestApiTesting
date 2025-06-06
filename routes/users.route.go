package routes

import (
	"encoding/json"
	"net/http"

	"github.com/cmartinez/GolangRestApiTesting/db"
	"github.com/cmartinez/GolangRestApiTesting/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	if db.DB.Error != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error retrieving users: " + db.DB.Error.Error()))
	}
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a single user"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	//createdUser := db.DB.Create(&user)
	db.DB.Create(&user)

	if db.DB.Error != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error creating user: " + db.DB.Error.Error()))
	}
	w.WriteHeader(http.StatusCreated)

	//json.NewEncoder(w).Encode(&createdUser)
	w.Write([]byte("Create a new user: " + user.FirstName + " " + user.LastName + " with email: " + user.Email))
}

func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update an existing user"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete a user"))
}
