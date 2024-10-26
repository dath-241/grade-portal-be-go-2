package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Name struct {
	LastName string `bson:"last_name" json:"LastName"`
	MFName   string `bson:"mf_name" json:"MFName"`
}

type Teacher struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	Name          Name                 `bson:"name" json:"name"`
	Email         string               `bson:"email" json:"email"`
	LastLogin     time.Time            `bson:"last_login" json:"LastLogin"`
	Role          string               `bson:"role" json:"role"`
	UpdatedAt     time.Time            `bson:"updated_at" json:"updated_at"`
	CreatedAt     time.Time            `bson:"created_at" json:"created_at"`
	TeacherID     primitive.ObjectID   `bson:"teacher_id" json:"teacher_id"`
	ManageClassID []primitive.ObjectID `bson:"manage_class_id" json:"manage_class_id"`
}
