package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"task-management-api/entities"
	"task-management-api/models"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ini adalah home page route..")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	// Ambil data tasks dari model
	tasks := models.GetTasks() // Sesuaikan nama fungsi dengan model

	// Siapkan data dalam format map
	data := map[string]any{
		"tasks": tasks,
	}

	// Ubah data menjadi format JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		// Jika terjadi error saat konversi ke JSON, kirim status error
		http.Error(w, "Error while converting data to JSON", http.StatusInternalServerError)
		return
	}

	// Set header response sebagai JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Set status response

	// Tulis JSON ke response body
	w.Write(jsonData)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Lakukan parsing ID ke integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Error converting ID: %v", err)
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Panggil fungsi model untuk mendapatkan data task berdasarkan ID
	task, err := models.GetTaskByID(id)
	if err != nil {
		log.Printf("Error fetching task with ID %d: %v", id, err)
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	log.Printf("Sending task response: %+v", task) // Logging data yang akan dikirim
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Printf("Error encoding task data: %v", err)
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entities.Tasks

	// Parse body request
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Memanggil fungsi CreateTask di model
	err = models.CreateTask(task)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task created successfully"})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task entities.Tasks

	// Parse body request
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Memanggil fungsi UpdateTask di model
	err = models.UpdateTask(id, task)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Memanggil fungsi DeleteTask di model
	err = models.DeleteTask(id)
	if err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}
