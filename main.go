package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"task-management-api/configs"
	"task-management-api/routes"
)

func main() {
	configs.LoadConfig()
	configs.ConnectDB()

	r := mux.NewRouter()
	routes.RouteIndex(r)

	log.Info("Server is running on port ", configs.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", configs.ENV.PORT), r)
}
