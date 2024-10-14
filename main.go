package main

import (
	"fmt"
	"log"
	"net/http"
	"task-management-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomePage).Methods("GET")
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/detail/{id}", controllers.Detail).Methods("GET")
	router.HandleFunc("/create", controllers.Create).Methods("POST")
	router.HandleFunc("/delete/{id}", controllers.Delete).Methods("DELETE")
	router.HandleFunc("/update/{id}", controllers.Update).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("This server running on port:8080")
}
