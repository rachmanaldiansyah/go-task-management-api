package main

import (
	"log"
	"net/http"
	"task-management-api/config"
	"task-management-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomePage).Methods("GET")
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.GetTaskByID).Methods("GET")
	router.HandleFunc("/create", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/delete/{id}", controllers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", controllers.UpdateTask).Methods("PUT")

	// Jalankan server di port 8080
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
