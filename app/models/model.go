package model

import "go.mongodb.org/mongo-driver/bson/primitive"

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

type Error struct {
	Error   bool   `json:error`
	Message string `json:message`
}
