package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"task-management-api/configs"
	"task-management-api/middleware"
	"task-management-api/routes"
)

func main() {
	// Load configuration and connect to the database
	configs.LoadConfig()
	configs.ConnectDB()

	// Initialize the router
	r := mux.NewRouter()

	// Register application routes
	routes.RouteIndex(r)

	// Apply CORS middleware
	handlerWithCORS := middleware.CORS(r)

	// Start the server
	log.Info("Server is running on port ", configs.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", configs.ENV.PORT), handlerWithCORS)
}
