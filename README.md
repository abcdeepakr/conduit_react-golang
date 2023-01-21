
# Table of contents
1. [Why this project?](#why-this-project?)
2. [Running backend locally](#run-locally)
3. [File structure](#file-structure)
4. [Models](#models)
5. [Controllers](#controllers)
6. [Routes](#routes)

## Why this project?<a name="why-this-project?"></a>

Conduit or [realworld](https://realworld.io) is a project on GitHub which allows people to learn new stacks and technologies.
It helps people to learn and deploy same project in different frameworks.
I wanted to learn Go for quite a while now, and hence started with this project. 

This project will give up an idea about the following concepts
- File structure in Go
- REST APIs in Go
- Connecting to MongoDB in Go
- Creating REST APIs in Go, with real time data
- working with JWT in Go
- Authentication and authorization in Go

## Running Backend Locally<a name="run-locally"></a>

Make sure you have the following tools installed/ready

- Go Programming Language
- .env
	- MONGODB_CONNECTION_STRING
	- DATABASE_NAME

Enter command `go run main.go` this should probably start the server locally
You are then good to make the API calls using the following [routes](#routes)

This is one of the first project that I took after going over the basics of Go. I am learning, please let me know the places where things are over complicated, and places where wrong patters are used.

### Testing locally on Postman

#### Creating a user
```
[POST] : localhost:4000/api/users
 
Body
Raw (JSON):
{
    "name" : "rajma chawal",
    "username" : "rajmaaaaa",
    "password" : "12345"
}

Header
Authorization : Bearer <YOUR BEARER TOKEN>

```

Similarly you can make other requests using this example 
## File structure<a name="file-structure"></a>
```text
|__Conduit_backend
   |__app
	   |__Controllers
	   |__Database
	   |__models
	   |__Routes
   |__.env
   |__main.go
```

Let's break down these files into simpler modules

The parent folder is conduit_backend

### main.go

Main.go is the entry point of every go project, it is created once we run the command `go mod init main.go` . 
our main.go file is responsible for
- Call the functions in the `database` folder that creates a connection with the database
- Create a Router which starts a local server.
-  This port number is something that could cause deployment issues, this is something I faces, but using 3000 on heroku worked fine for me.

### App

The App folder contains the major folders required to implement the backend such as
- Controllers
	- This folder contains out business logic
- Models
	- This folder contains the files which creates different interfaces in Go
- Routes
	- This folder contains a file which creates a router using mux router, which is a package in go used to create routers.
- Database
	- As the name suggets this folder is used to connect with the database, which is MongoDB in our case.


#### Controllers<a name="controllers"></a>

We use the following packages here, and most of the code crunching is done in this folder

- Mongo - to perform mongodb operations
- godotenv - accessing env variables
- bson - working with JSON data

Divided in to 3 parts
- Articles
- Auth
- users
Each folder works on respective operations

#### Models<a name="models"></a>
This directory is crucial to create interfaces, and validating them
Example of an Interface
```go
type User struct {
    ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Name      string             `json:name`
    UserName  string             `json:username`
    Password  string             `json:password`
    Followers []string           `json:followers`
    Following []string           `json:following`
    Articles  []string           `json:articles`
}
```

#### Routes<a name="routes"></a>

Controllers are imported in this file
```go

import (
    "github.com/gorilla/mux"
    articleController "github.com/deepakr-28/conduit_golang_backend/app/controllers/articles"
    userController "github.com/deepakr-28/conduit_golang_backend/app/controllers/users"
)
```


Following routes are currently present in the project

```GO
// User Routes

POST: /api/users // CREATE USER
POST: /api/users/login // LOGIN USER
GET:  /api/user //GET USER
PUT:  /api/user //UPDATE USER
GET:  /api/profiles/{username} //GET USER INFO
POST: /api/profiles/{username}/follow // LIKE USER
POST: /api/profiles/{username}/follow // FOLLOW USER

// Article Routes
POST: /api/articles // NEW ARTICLE
GET: /api/articles/{slug} //GET ARTICLE
DELETE: /api/articles/{slug} // DELETE ARTICLE
```