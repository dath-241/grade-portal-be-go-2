package controllers_admin

import (
	"be-go-2/config"
	"be-go-2/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	collection   *mongo.Collection
	dbName       = "Course"
	dbCollection = "Courses"
)

// CreateCoures func to create a course and post to DB
func CreateCourse(ctx *gin.Context) {

	var course models.Course
	if err := ctx.BindJSON(&course); err != nil {

		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	newUser := models.Course{
		CourseId:    course.CourseId,
		CourseName : course.CourseName,
		Credit:      course.Credit,
		Description: course.Description,
		
		//!TODO: CreatedBy stores ID of the Creator of the Course, temporary using random due to lack of data
		CreatedBy: primitive.NewObjectID(),
	}
	collection = config.MongoClient.Database(dbName).Collection(dbCollection)

	// result, err // use result to see the _id
	_, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newUser)
}

// GetCourse gets a course by its ID
func GetCourse(ctx *gin.Context) {

}
