package models

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceStudent struct {
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Name     string        `json:"name"`
	UID      string        `json:"UID"`
	Faculty  string        `json:"faculty"`
	ClassID  []interface{} `json:"classID"`
}

func StudentModel() *mongo.Collection {
	initModel("gradeDB", "student")
	return collection
}
