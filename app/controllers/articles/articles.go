package articleController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	tokenPackage "github.com/deepakr-28/conduit_golang_backend/app/controllers/authToken"
	"github.com/deepakr-28/conduit_golang_backend/app/database"
	model "github.com/deepakr-28/conduit_golang_backend/app/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var ArticleCollection *mongo.Collection
var UserCollection *mongo.Collection

var databaseName string

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseName = os.Getenv("DATABASE_NAME")
}

func authenticateUser(token string) bool {
	username := tokenPackage.VerifyToken(token) // TODO RETURN A VALUE IN VERIFY TOKEN
	if username != "" {
		return true
	} else {
		return false
	}
}

func createArticle(article model.Article, token string) {

	var user model.User
	// USER COLLECTION IS USED TO UPDATE THE ARTICLES IN USER
	ArticleCollection = database.Client.Database(databaseName).Collection("articles")
	UserCollection = database.Client.Database(databaseName).Collection("users")

	// inserted
	inserted, err := ArticleCollection.InsertOne(context.Background(), article)
	if err != nil {
		log.Fatal(err)
	}
	// GET USERNAME FROM AUTH TOKEN, AND UPDATES THEIR ARTICLE ARRAY WITH SLUGS
	username := tokenPackage.VerifyToken(token)

	// FIND A USER, MAKE CHANGES, AND UDPATE
	// TODO TRY IF WE CAN UPDATE IN 1 QUERY
	findUsererror := UserCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if findUsererror != nil {
		log.Fatal(findUsererror)
	}

	user.Articles = append(user.Articles, article.Slug)
	filter := bson.M{"username": username} // get username from the token
	updatedData := bson.M{
		"$set": user,
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	// THIS QUERY UPDATES THE USER
	updateUserError := UserCollection.FindOneAndUpdate(context.TODO(), filter, updatedData, &opt).Decode(&user)

	if updateUserError != nil {
		log.Fatal(updateUserError)
	}

	fmt.Println("Created a new article with id: ", inserted.InsertedID)
	fmt.Println("UPATED USER DATA", user)
}

func getArticle(slug string) model.Article {
	ArticleCollection = database.Client.Database(databaseName).Collection("articles")
	var article model.Article
	err := ArticleCollection.FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&article)

	if err != nil {

		log.Fatal("ERROR : ", err)
	}
	fmt.Println(article)
	return article
}

func updateArticle(article model.Article) model.Article {

	var updatedArticle model.Article
	ArticleCollection = database.Client.Database(databaseName).Collection("articles")

	filter := bson.M{"slug": article.Slug} // get username from the token
	updatedData := bson.M{
		"$set": article,
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := ArticleCollection.FindOneAndUpdate(context.TODO(), filter, updatedData, &opt).Decode(&updatedArticle)

	if err != nil {
		log.Fatal(err)
	}

	return updatedArticle
}

func deleteArticle(slug string) bool {

	var article model.Article
	ArticleCollection = database.Client.Database(databaseName).Collection("articles")

	filter := bson.M{"slug": slug}
	err := ArticleCollection.FindOneAndDelete(context.TODO(), filter).Decode(&article)
	if err != nil {
		// log.Fatal("ERROR : ", err)
		fmt.Println(err)
		return false
	}
	fmt.Println(slug, " Deleted successfully")
	return true
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var article model.Article
	var error model.Response
	_ = json.NewDecoder(r.Body).Decode(&article)
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	userNameExists := authenticateUser(reqToken)

	if userNameExists {
		createArticle(article, reqToken)
		json.NewEncoder(w).Encode(article)
	} else {
		error.Error = true
		error.Message = "You Are not Authenticated"
		json.NewEncoder(w).Encode(error)
	}
}

func GetArticle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var error model.Response

	params := mux.Vars(r)
	response := getArticle(params["slug"])
	if response.Slug != "" {
		json.NewEncoder(w).Encode(response)
	} else {
		error.Error = true
		error.Message = "Article Not Found"
		json.NewEncoder(w).Encode(error)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var article model.Article
	var error model.Response

	_ = json.NewDecoder(r.Body).Decode(&article)

	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		error.Error = true
		error.Message = "Bearer Token not found"
		json.NewEncoder(w).Encode(error)
		return
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	response := updateArticle(article)

	username := authenticateUser(reqToken)

	if username {
		json.NewEncoder(w).Encode(response)
	} else {
		error.Error = true
		error.Message = "user not logged in"
		json.NewEncoder(w).Encode(error)
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var error model.Response

	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		error.Error = true
		error.Message = "Bearer Token not found"
		json.NewEncoder(w).Encode(error)
		return
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	username := authenticateUser(reqToken)
	params := mux.Vars(r)
	if username {

		response := deleteArticle(params["slug"])
		if response {
			error.Error = false
			error.Message = "Deleted Successfully"
			json.NewEncoder(w).Encode(error)

		} else {
			error.Error = true
			error.Message = "We Faced Some problem While deleting your article"
			json.NewEncoder(w).Encode(error)
		}

	} else {
		error.Error = true
		error.Message = "user not logged in"
		json.NewEncoder(w).Encode(error)
	}
}
