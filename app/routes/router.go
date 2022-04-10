package router

import (
	"github.com/gorilla/mux"

	userController "github.com/deepakr-28/conduit_golang_backend/app/controllers/users"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	// router.HandleFunc("/api/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/users", userController.CreateUser).Methods("POST")
	return router
}
