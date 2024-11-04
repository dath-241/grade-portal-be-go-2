package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GradeSheet struct {
	ClassID   primitive.ObjectID `bson:"classID"`
	CourseID  primitive.ObjectID `bson:"courseID"`
	ExpiredAt time.Time          `bson:"expiredAt"`
	CreatedBy primitive.ObjectID `bson:"createdBy"`
	UpdatedBy primitive.ObjectID `bson:"updatedBy,omitempty"`
	GradeData []StudentGrades    `bson:"gradeSheet,omitempty"`
}

type StudentGrades struct {
	StudentID string `bson:"studentID"`
	Grades    Grades `bson:"grades"`
}

type Grades struct {
	Homework   []float64 `bson:"Homework"`
	Lab        []float64 `bson:"Lab"`
	Assignment []float64 `bson:"Assignment"`
	Midterm    float64   `bson:"Midterm"`
	Final      float64   `bson:"Final"`
}
