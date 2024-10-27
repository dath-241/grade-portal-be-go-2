package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	Id          primitive.ObjectID `bson:"Id"`
	Course_name string             `bson:"Course_name" binding:"required"`
	Credit      int32              `bson:"Credit" binding:"required"`
	Class_id    []string           `bson:"Class_id" binding:"required"`
}
