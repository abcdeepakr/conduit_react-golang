package router

import (
	"github.com/gorilla/mux"

	articleController "github.com/deepakr-28/conduit_golang_backend/app/controllers/articles"
	userController "github.com/deepakr-28/conduit_golang_backend/app/controllers/users"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	// USER ROUTES
	router.HandleFunc("/api/users", userController.CreateUser).Methods("POST")                          // creates user
	router.HandleFunc("/api/users/login", userController.AuthenticateUser).Methods("POST")              // login user
	router.HandleFunc("/api/user", userController.GetCurrentUser).Methods("GET")                        // get loggedin user user
	router.HandleFunc("/api/user", userController.UpdateUser).Methods("PUT")                            // update user
	router.HandleFunc("/api/profiles/{username}", userController.GetUser).Methods("GET")                // get specific user
	router.HandleFunc("/api/profiles/{username}/follow", userController.FollowUser).Methods("POST")     // FOLLOW USER
	router.HandleFunc("/api/profiles/{username}/follow", userController.UnfollowUser).Methods("DELETE") // UNFOLLOW USER

	// Article Routes
	router.HandleFunc("/api/articles", articleController.CreateArticle).Methods("POST")          // Create Article
	router.HandleFunc("/api/articles/{slug}", articleController.GetArticle).Methods("GET")       // get specific article
	router.HandleFunc("/api/articles/{slug}", articleController.DeleteArticle).Methods("DELETE") // delete specific article

	return router
}
