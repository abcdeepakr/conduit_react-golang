package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	database "github.com/deepakr-28/conduit_golang_backend/app/database"
	model "github.com/deepakr-28/conduit_golang_backend/app/models"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

// const dbName = "conduit_golang_backend"
const collectionName = "users"

var databaseName string

func insertUser(user model.User) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseName = os.Getenv("DATABASE_NAME")
	res, err := database.Client.ListDatabaseNames(context.Background(), bson.M{})
	// defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	Collection = database.Client.Database(databaseName).Collection(collectionName)

	inserted, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 user in db with id: ", inserted.InsertedID)

}
