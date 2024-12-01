package models

import (
	"Go2/config"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

var collection *mongo.Collection

func initModel(database string, col string) {
	if config.MongoClient == nil {
		log.Fatal("MongoDB client chưa được khởi tạo")
	}
	collection = config.MongoClient.Database(database).Collection(col)
}
