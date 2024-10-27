package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GradeSheet struct {
	StudentID primitive.ObjectID `bson:"student_id" json:"student_id"`
	ClassID   primitive.ObjectID `bson:"class_id" json:"class_id"`
	Status    string             `bson:"status" json:"status"` // open or close
	Grade     GradeDetails       `bson:"grade" json:"grade"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type GradeDetails struct {
	Final      float64   `bson:"final" json:"final"`
	Midterm    float64   `bson:"midterm" json:"midterm"`
	Assignment []float64 `bson:"assignment" json:"assignment"`
	Exercise   []float64 `bson:"exercise" json:"exercise"`
}
