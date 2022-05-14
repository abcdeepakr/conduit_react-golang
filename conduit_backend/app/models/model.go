package model

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:name`
	UserName  string             `json:username`
	Password  string             `json:password`
	Followers []string           `json:followers`
	Following []string           `json:following`
	Articles  []string           `json:articles`
}

type Article struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Slug        string             `json:slug`
	Title       string             `json:title`
	Description string             `json:description`
	Body        string             `json:body`
	Tags        []string           `json:tags`
	CreatedAt   time.Time          `json:created_at`
	Author      string             `json:author`
}

type AuthenticatedResponse struct {
	User      User   "json:user"
	JsonToken string "json:jwt"
}

type Response struct {
	Error   bool   `json:error`
	Message string `json:message`
}

type Payload struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// the paylod for the jwt token

type JWTMaker struct {
	secretKey string
}
