package main

import (
	"net/http"

	"github.com/cmartinez/GolangRestApiTesting/db"
	"github.com/cmartinez/GolangRestApiTesting/models"
	"github.com/cmartinez/GolangRestApiTesting/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnect()

	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Task{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	// User routes
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.PutUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// Task routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.PostTasksHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.PutTasksHandler).Methods("PUT")
	r.HandleFunc("/tasks/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
