package routes

import "github.com/gorilla/mux"

func RouteIndex(r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()

	TaskRouter(api)
}
