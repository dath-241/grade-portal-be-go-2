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
	CreatedAt time.Time     `bson:"createdAt"`
	CreatedBy bson.ObjectID `bson:"createdBy"`
}

func AdminModel() *mongo.Collection {
	initModel("gradeDB", "admin")
	return collection
}
