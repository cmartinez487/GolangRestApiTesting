package main

import (
	"net/http"

	"github.com/cmartinez/GolangRestApiTesting/config"
	"github.com/cmartinez/GolangRestApiTesting/internal/handlers"
	"github.com/cmartinez/GolangRestApiTesting/internal/repository"
	"github.com/cmartinez/GolangRestApiTesting/internal/service"
	"github.com/cmartinez/GolangRestApiTesting/migrations"
	"github.com/gorilla/mux"
)

func main() {
	db := config.InitDB()

	// Ejecutar migraciones
	migrations.MigrateDB(db)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	router := mux.NewRouter().StrictSlash(true) // Agrega esto para evitar conflictos
	router.HandleFunc("/users", userHandler.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserHandler).Methods("GET") // Asegura que ID sea numérico
	router.HandleFunc("/users", userHandler.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", userHandler.PutUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUserHandler).Methods("DELETE")

	router.HandleFunc("/tasks", taskHandler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.GetTaskHandler).Methods("GET") // Asegura que ID sea numérico
	router.HandleFunc("/tasks", taskHandler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.PutTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
