package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/deepakr-28/conduit_golang_backend/app/database"
	router "github.com/deepakr-28/conduit_golang_backend/app/routes"
)

func main() {

	database.ConnectDB()
	fmt.Println("MONGODB API")
	r := router.Router()
	fmt.Println("SERVER STARTING...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("SERVER STARTED AT PORT 4000")
}
