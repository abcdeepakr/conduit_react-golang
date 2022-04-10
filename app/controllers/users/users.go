package userController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	database "github.com/deepakr-28/conduit_golang_backend/app/database"
	model "github.com/deepakr-28/conduit_golang_backend/app/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var Collection *mongo.Collection

const collectionName = "users"

var databaseName string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseName = os.Getenv("DATABASE_NAME")

}

func createUser(user model.User) {
	Collection = database.Client.Database(databaseName).Collection(collectionName)

	inserted, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created a new user with id: ", inserted.InsertedID)

}

func checkUsername(collection *mongo.Collection, user string) bool {
	// data, err := collection.Distinct(context.Background(), user, bson.A{})
	var result model.User
	// opts := options.FindOne().SetSort(bson.D{{"username", 1}})

	// filter := bson.D{{"username", user}}
	fmt.Println("username : ", user)
	err := collection.FindOne(context.TODO(), bson.M{"username": user}).Decode(&result)
	fmt.Println(err)
	if err != nil {
		return false
	} else {
		return true
	}
}

// func authenticateUser() {}

// func getUser() {}

// func updateUser() {}

// PARENT FUNCTIONS WHICH WILL BE USED IN ROUTER FILES ARE DEFINED BELOW.
// FUNCTIONS ABOVE ARE HELPER METHODS WHICH ARE CALLED FROM THE FUNCTIONS BELOW.

func CreateUser(w http.ResponseWriter, r *http.Request) {
	Collection = database.Client.Database(databaseName).Collection(collectionName)

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.User
	var error model.Error
	_ = json.NewDecoder(r.Body).Decode(&user)

	userNameExists := checkUsername(Collection, user.UserName)
	fmt.Println(userNameExists)
	if userNameExists {
		error.Error = true
		error.Message = "Username taken, please try another one"
		json.NewEncoder(w).Encode(error)
	} else {
		createUser(user)
		json.NewEncoder(w).Encode(user)

	}

}
