package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

// this function will connect with the database
func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var databaseConnectionString = os.Getenv("MONGODB_CONNECTION_STRING")
	fmt.Println(databaseConnectionString)
	var connectionString = databaseConnectionString
	const dbName = "conduit_golang_backend"
	const collectionName = "users"

	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance is ready")
}
