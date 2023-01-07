package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/deepakr-28/conduit_golang_backend/app/database"
	router "github.com/deepakr-28/conduit_golang_backend/app/routes"
)

func main() {
	var databaseConnectionString = os.Getenv("PORT")
	database.ConnectDB()
	fmt.Println("MONGODB API")
	r := router.Router()
	fmt.Println("SERVER STARTING...")
	fmt.Println("SERVER STARTED AT PORT " + databaseConnectionString)
	log.Fatal(http.ListenAndServe(":"+databaseConnectionString, r))
	fmt.Println("SERVER STARTED AT PORT " + databaseConnectionString)
}
