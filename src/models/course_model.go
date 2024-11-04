package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Course model describe the Course model uses for the database
type Course struct {
	CourseId    string             `bson:"courseId" binding:"required"`
	CourseName  string             `bson:"courseName" binding:"required"`
	Credit      int32              `bson:"credit" binding:"required"`
	Description string             `bson:"description"`
	CreatedBy   primitive.ObjectID `bson:"createdBy"`
}
