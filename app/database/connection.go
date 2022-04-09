package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// var Collection *mongo.Collection
var Client *mongo.Client

// this function will connect with the database
func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var databaseConnectionString = os.Getenv("MONGODB_CONNECTION_STRING")
	fmt.Println(databaseConnectionString)

	var connectionString = databaseConnectionString

	// const dbName = "conduit_golang_backend"
	// const collectionName = "users"

	// clientOption := options.Client().ApplyURI(connectionString)
	// client, err := mongo.Connect(context.TODO(), clientOption)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Collection = client.Database(dbName).Collection(collectionName)
	// fmt.Println("Collection instance is ready", &Collection)

	Client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	// fmt.Printf(Clien)
	databases, err := Client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DATABASES ARE : ", databases[0])

}
