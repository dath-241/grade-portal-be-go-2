package helper

import (
	// "errors"
	// "fmt"
	// "os"
	// "time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Claims struct {
	ID                   bson.ObjectID `json:"id"` // Email được lưu trong token
	jwt.RegisteredClaims             
}

// func CreateJWT(id bson.ObjectID) string {}

// func ParseJWT(tokenString string) (*Claims, error) {}
