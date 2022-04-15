package userController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	database "github.com/deepakr-28/conduit_golang_backend/app/database"
	model "github.com/deepakr-28/conduit_golang_backend/app/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func authenticateUser(collection *mongo.Collection, user model.User) bool {

	var result model.User
	fmt.Println("username : ", user)
	err := collection.FindOne(context.TODO(), bson.M{"username": user.UserName}).Decode(&result)

	if err != nil {
		log.Fatal(err)
		return false
	} else {
		if result.Password == user.Password {
			return true
		} else {
			return false
		}

	}
}

func getCurrentUser(collection *mongo.Collection, token string) bool {
	// this function will return a user based on the jwt token
	// decode token here, get username, search username and return the user
	decodedToken := "deepak"

	if decodedToken == token {
		// var result model.User
		// fmt.Println("username : ", user)
		// err := collection.FindOne(context.TODO(), bson.M{"username": user.UserName}).Decode(&result)
		return true
	} else {
		return false
	}
}

func updateUser(collection *mongo.Collection, user model.User, token string) bool {

	// this function will return a user based on the jwt token
	// decode token here, get username, search username and return the user

	var result model.User
	decodedToken := "deepak"
	if decodedToken == token {
		filter := bson.M{"username": token} // get username from the token
		updatedData := bson.M{
			"$set": user,
		}

		upsert := true
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}

		err := collection.FindOneAndUpdate(context.TODO(), filter, updatedData, &opt).Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		return true
	} else {
		return false
	}
}

func getUser(collection *mongo.Collection, username string) model.User {

	var result model.User
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)

	if err != nil {
		result.UserName = "null"
		return result
	}
	return result

}

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
	if userNameExists {
		error.Error = true
		error.Message = "Username taken, please try another one"
		json.NewEncoder(w).Encode(error)
	} else {
		createUser(user)
		json.NewEncoder(w).Encode(user)
	}
}
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	Collection = database.Client.Database(databaseName).Collection(collectionName)

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.User
	var error model.Error

	_ = json.NewDecoder(r.Body).Decode(&user)
	response := authenticateUser(Collection, user)
	if response {

		responseData := model.AuthenticatedResponse{User: user, JsonToken: "something"}
		json.NewEncoder(w).Encode(responseData)
	} else {
		error.Error = true
		error.Message = "Password or Username Invalid, please check and re-enter"
		json.NewEncoder(w).Encode(error)
	}
}
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	Collection = database.Client.Database(databaseName).Collection(collectionName)

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.User
	var error model.Error

	_ = json.NewDecoder(r.Body).Decode(&user)
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	response := getCurrentUser(Collection, reqToken)

	if response {

		json.NewEncoder(w).Encode("currnt user")
	} else {
		error.Error = true
		error.Message = "user not logged in"
		json.NewEncoder(w).Encode(error)
	}
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	Collection = database.Client.Database(databaseName).Collection(collectionName)

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.User
	var error model.Error

	_ = json.NewDecoder(r.Body).Decode(&user)

	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		error.Error = true
		error.Message = "Bearer Token not found"
		json.NewEncoder(w).Encode(error)
		return
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	response := updateUser(Collection, user, reqToken)

	if response {
		json.NewEncoder(w).Encode(user)
	} else {
		error.Error = true
		error.Message = "user not logged in"
		json.NewEncoder(w).Encode(error)
	}
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	Collection = database.Client.Database(databaseName).Collection(collectionName)

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var error model.Error

	params := mux.Vars(r)
	response := getUser(Collection, params["username"])

	if response.UserName != "null" {
		json.NewEncoder(w).Encode(response)
	} else {
		error.Error = true
		error.Message = "User Not Found"
		json.NewEncoder(w).Encode(error)
	}
}
