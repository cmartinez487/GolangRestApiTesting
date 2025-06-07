package routes

import (
	"encoding/json"
	"net/http"

	"github.com/cmartinez/GolangRestApiTesting/db"
	"github.com/cmartinez/GolangRestApiTesting/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	if db.DB.Error != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error retrieving users: " + db.DB.Error.Error()))
	}
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404 Not Found
		w.Write([]byte("Task not found"))
		return
	} else if db.DB.Error != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error retrieving user: " + db.DB.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)
	err := createdTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error creating user: " + db.DB.Error.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(&task)
}

func PutTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404 Not Found
		w.Write([]byte("Task not found"))
		return
	} else if db.DB.Error != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error retrieving task: " + db.DB.Error.Error()))
		return
	}

	var updatedData models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON format"))
		return
	}

	// Realizar la actualización con GORM
	if err := db.DB.Model(&task).Updates(updatedData).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500 Internal Server Error
		w.Write([]byte("Error updating task: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task updated successfully"))
}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404 Not Found
		w.Write([]byte("Task not found"))
		return
	} else if db.DB.Error != nil {
		w.WriteHeader(http.StatusBadRequest) //400 Bad Request
		w.Write([]byte("Error retrieving task: " + db.DB.Error.Error()))
		return
	}

	db.DB.Delete(&task, params["id"])

	w.WriteHeader(http.StatusOK)
}
