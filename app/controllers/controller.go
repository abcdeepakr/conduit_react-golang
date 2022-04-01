package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	database "github.com/deepakr-28/conduit_golang_backend/app/database"
	model "github.com/deepakr-28/conduit_golang_backend/app/models"
)

func insertUser(user model.User) {

	inserted, err := database.Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 user in db with id: ", inserted.InsertedID)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	insertUser(user)

	json.NewEncoder(w).Encode(user)
}
