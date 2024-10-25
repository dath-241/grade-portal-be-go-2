package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceAdmin struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
	Name      string        `bson:"name"`
	Faculty   string        `bson:"faculty"`
	Role      string        `bson:"role"`
	CreatedAt time.Time     `bson:"createdAt"`
	ExpiredAt time.Time     `bson:"expiredAt"`
}

func AdminModel() *mongo.Collection {
	initModel("gradeDB", "admin")
	return collection
}
