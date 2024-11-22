package taskscontrollers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"task-management-api/configs"
	"task-management-api/helpers"
	"task-management-api/models"
)

// GetTasks retrieves a list of all tasks from the database and sends them in the response.
// If an error occurs during the database query, it responds with an internal server error.
func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	if err := configs.DB.Find(&tasks).Error; err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(w, http.StatusOK, "List of tasks", tasks)
}

// CreateTask creates a new task in the database based on the JSON payload of the incoming request.
// If an error occurs during the database query, it responds with an internal server error.
// If the task is created successfully, it responds with a 201 status code and an empty response body.
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var tasks models.Task

	if err := json.NewDecoder(r.Body).Decode(&tasks); err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := configs.DB.Create(&tasks).Error; err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(w, http.StatusCreated, "Task created successfully", nil)
}

// GetTasksByID retrieves a task from the database by the given id and sends it in the response.
// If the task is not found, it responds with a 404 status code and an appropriate error message.
// If an error occurs during the database query, it responds with an internal server error.
func GetTasksByID(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var tasks models.Task

	if err := configs.DB.First(&tasks, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, http.StatusNotFound, "Task not found", nil)
			return
		}

		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(w, http.StatusOK, "Task detail", tasks)
}

// UpdateTask retrieves a task from the database by the given id, updates its fields with the given
// request body, and sends a success response. If the task is not found, it responds with a 404 status
// code and an appropriate error message. If an error occurs during the database query, it responds
// with an internal server error.
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var tasks models.Task

	if err := configs.DB.First(&tasks, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, http.StatusNotFound, "Task not found", nil)
			return
		}

		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&tasks); err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := configs.DB.Where("id = ?", id).Updates(&tasks).Error; err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(w, http.StatusOK, "Task updated successfully", nil)
}

// DeleteTask retrieves a task from the database by the given id, deletes it, and sends a success
// response. If the task is not found, it responds with a 404 status code and an appropriate error
// message. If an error occurs during the database query, it responds with an internal server error.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var tasks models.Task

	res := configs.DB.Delete(&tasks, id)
	if res.Error != nil {
		helpers.Response(w, http.StatusInternalServerError, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helpers.Response(w, http.StatusNotFound, "Task not found", nil)
		return
	}

	helpers.Response(w, http.StatusOK, "Task deleted successfully", nil)
}
