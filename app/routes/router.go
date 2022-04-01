package router

import (
	"github.com/gorilla/mux"

	controller "github.com/deepakr-28/conduit_golang_backend/app/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.CreateUser).Methods("POST")
	return router

}
