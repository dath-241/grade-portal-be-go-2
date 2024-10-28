package models 
import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Course struct {
	MS        string `bson:"ms"`
	CourseID  string `bson:"course_id,omitempty"`
	Credit    int    `bson:"credit"`
	Name      string `bson:"name"`
	Desc      string `bson:"desc"`
	CreatedBy string `bson:"createdby"`
}

func CourseModle() *mongo.Collection {
	initModel("gradeDB", "course")
	return collection
}