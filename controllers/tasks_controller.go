package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-management-api/entities"

	"github.com/gorilla/mux"
)

var tasks []entities.Tasks

func AllTasks() {
	task := entities.Tasks{
		ID:        "1",
		TaskName:  "New Projects",
		TaskDesc:  "You must lead the project and finish it",
		CreatedBy: "2024-10-14",
		UpdatedBy: "2024-10-14",
	}

	tasks = append(tasks, task)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ini adalah home page route..")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if taskId["id"] == tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}
	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}

func Create(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
