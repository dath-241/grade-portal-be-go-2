package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceClass struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	Semester      string        `bson:"semester"`
	Name          string        `bson:"name"` 
	CourseId      bson.ObjectID `bson:"course_id"`
	ListStudentMs []string      `bson:"listStudent_ms"`
	TeacherId     bson.ObjectID `bson:"teacher_id"`
	CreatedBy     bson.ObjectID `bson:"createdBy"` 
	UpdatedBy     bson.ObjectID `bson:"updatedBy"`
	CsvURL 				string   			`bson:"csv_url"`
	LastModified 	string    		`bson:"last_mod"`
}
type InterfaceClassStudent struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Semester  string        `bson:"semester"`
	Name      string        `bson:"name"` 
	CourseId  bson.ObjectID `bson:"course_id"`
	TeacherId bson.ObjectID `bson:"teacher_id"`
}

func ClassModel() *mongo.Collection {
	InitModel("gradeportal", "class")
	return collection
}
