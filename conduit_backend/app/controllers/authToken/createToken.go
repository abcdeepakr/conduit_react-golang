package tokenPackage

import (
	"fmt"
	"log"
	"time"

	model "github.com/deepakr-28/conduit_golang_backend/app/models"
	"github.com/golang-jwt/jwt"
)

func CreateToken(username string) model.Response {

	var hmacSampleSecret []byte
	var tokenCreationResponse model.Response

	claims := model.Payload{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 120).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		log.Fatal("err", err)
		// tokenCreationResponse.Error = true
		// tokenCreationResponse.Message = err.Error()
		return tokenCreationResponse
	}
	tokenCreationResponse.Error = false
	tokenCreationResponse.Message = tokenString
	return tokenCreationResponse
}

func VerifyToken(generatedToken string) string {
	// var hmacSampleSecret []byte
	tokenString := generatedToken
	fmt.Print(tokenString)
	type Payload struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	// https://pkg.go.dev/github.com/golang-jwt/jwt#NewWithClaims
	// Override time value for tests.  Restore default value after.

	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if err != nil {
		log.Fatal("ERROR", err)
		return ""
	}

	claims, ok := token.Claims.(*Payload)
	if ok && token.Valid {
		fmt.Println(claims)
	}

	return claims.Username
}
