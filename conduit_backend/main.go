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
	var PORT = os.Getenv("HTTP_PLATFORM_PORT")
	database.ConnectDB()
	fmt.Println("MONGODB API")
	r := router.Router()
	fmt.Println("SERVER STARTED AT PORT " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
