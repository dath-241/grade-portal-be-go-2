package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Teacher struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Name      string             `bson:"name" json:"name"`
	UID       string             `bson:"UID" json:"UID"`
	Faculty   string             `bson:"faculty" json:"faculty"`
	Role      string             `bson:"role" json:"role"`
	CreatedBy primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	ExpiredAt time.Time          `bson:"expiredAt" json:"expiredAt"`
}
