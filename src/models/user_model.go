package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Student struct {
	UID       string        	`bson:"uid"`
	Name      string        	`bson:"name"`
	Faculty   string        	`bson:"faculty"`
	Role      string        	`bson:"role"`
	CreatedBy bson.ObjectID 	`bson:"created_by"`
	CreateAt bson.ObjectID 		`bson:"create_at"`
	ExpiredAt time.Time     	`bson:"expired_at"`
	Subjects []bson.ObjectID 	`bson:"subjects"`
}

func StudentModel() *mongo.Collection {
    initModel("gradeDB", "user")
    return collection
}

type Teacher struct {
	UID       string        	`bson:"uid"`
	Name      string        	`bson:"name"`
	Faculty   string        	`bson:"faculty"`
	Role      string        	`bson:"role"`
	CreatedBy bson.ObjectID 	`bson:"created_by"`
	CreateAt bson.ObjectID     	`bson:"create_at"`
	ExpiredAt time.Time     	`bson:"expired_at"`
	Subjects []bson.ObjectID    `bson:"subjects"`
}

func TeacherModel() *mongo.Collection {
    initModel("gradeDB", "user")
    return collection
}