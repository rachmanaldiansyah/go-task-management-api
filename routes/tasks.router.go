package routes

import (
	"github.com/gorilla/mux"

	taskscontrollers "task-management-api/controllers"
)

func TaskRouter(r *mux.Router) {
	router := r.PathPrefix("/tasks").Subrouter()

	router.HandleFunc("", taskscontrollers.GetTasks).Methods("GET")
	router.HandleFunc("/create", taskscontrollers.CreateTask).Methods("POST")
	router.HandleFunc("/detail/{id}", taskscontrollers.GetTasksByID).Methods("GET")
	router.HandleFunc("/update/{id}", taskscontrollers.UpdateTask).Methods("POST")
	router.HandleFunc("/delete/{id}", taskscontrollers.DeleteTask).Methods("DELETE")
}
