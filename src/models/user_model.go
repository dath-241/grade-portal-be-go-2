package models

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceStudent struct {
}

func StudentModel() *mongo.Collection {
	initModel("gradeDB", "student")
	return collection
}
