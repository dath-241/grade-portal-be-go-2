package models 

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InterfaceAdmin struct {
	ID        bson.ObjectID `bson:"_id,omitempty"` 
	Email     string        `bson:"email"`        
	Name      string        `bson:"name"`         
	Ms        string        `bson:"ms"`
	Faculty   string        `bson:"faculty"`
	CreatedBy bson.ObjectID `bson:"createdBy"` 
}

func AdminModel() *mongo.Collection {
	InitModel("gradeportal", "admin")
	return collection
}
