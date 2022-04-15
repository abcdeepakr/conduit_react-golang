package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:name`
	UserName string             `json:username`
	Password string             `json:password`
}

type AuthenticatedResponse struct {
	User      User   "json:user"
	JsonToken string "json:jwt"
}

type Response struct {
	Error   bool   `json:error`
	Message string `json:message`
}

// the paylod for the jwt token
type Payload struct {
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type JWTMaker struct {
	secretKey string
}
