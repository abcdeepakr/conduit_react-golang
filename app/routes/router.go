package router

import (
	"github.com/gorilla/mux"

	userController "github.com/deepakr-28/conduit_golang_backend/app/controllers/users"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	// router.HandleFunc("/api/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/users", userController.CreateUser).Methods("POST")             // creates user
	router.HandleFunc("/api/users/login", userController.AuthenticateUser).Methods("POST") // login user
	router.HandleFunc("/api/user", userController.GetCurrentUser).Methods("GET")           // login user
	return router
}