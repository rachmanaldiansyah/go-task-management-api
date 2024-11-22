package routes

import (
	"github.com/gorilla/mux"

	taskscontrollers "task-management-api/controllers"
)

func TaskRouter(r *mux.Router) {
	router := r.PathPrefix("/tasks").Subrouter()

	router.HandleFunc("", taskscontrollers.GetTasks).Methods("GET")
	router.HandleFunc("", taskscontrollers.CreateTask).Methods("POST")
	router.HandleFunc("/{id}/detail", taskscontrollers.GetTasksByID).Methods("GET")
	router.HandleFunc("/{id}/update", taskscontrollers.UpdateTask).Methods("POST")
	router.HandleFunc("/{id}/delete", taskscontrollers.DeleteTask).Methods("DELETE")
}
