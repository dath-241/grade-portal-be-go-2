package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Class struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	ClassName          string   `bson:"class_name"` 
	CourseId      bson.ObjectID `bson:"course_id"`
	Semester      string        `bson:"semester"`
	ListStudentId []string      `bson:"listStudent_id"`
	TeacherId     bson.ObjectID `bson:"teacher_id"`
	CreatedBy     bson.ObjectID `bson:"createdBy"` 
	UpdatedBy     bson.ObjectID `bson:"updatedBy"`
}

func ClassModel() *mongo.Collection {
	initModel("gradeDB", "class")
	return collection
}
