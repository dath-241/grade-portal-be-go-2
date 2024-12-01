package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type InterfaceAccount struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Email     string        `bson:"email"`
	Name      string        `bson:"name"`
	Ms        string        `bson:"ms"`
	Faculty   string        `bson:"faculty"`
	Role      string        `bson:"role"`
	CreatedBy bson.ObjectID `bson:"createdBy"`
	ExpiredAt time.Time     `bson:"expiredAt"`
}

func AccountModel() *mongo.Collection {
	InitModel("gradeportal", "account")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "expiredAt", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatalf("Failed to create TTL index: %v", err)
	}
	return collection
}
