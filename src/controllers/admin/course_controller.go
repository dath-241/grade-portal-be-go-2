package controllers_admin

import (
	"be-go-2/config"
	"be-go-2/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	collection   *mongo.Collection
	dbName       = "Course"
	dbCollection = "Courses"
)

// CreateCoures func to create a course and post to DB
func CreateCourse(ctx *gin.Context) {
	// c, cancel := context.WithTimeoutCause(context.Background(), 10*time.Second)

	var course models.Course

	if err := ctx.BindJSON(&course); err != nil {

		ctx.JSON(http.StatusBadRequest, err) 
		return
	}

	newUser := models.Course{
		Id:          primitive.NewObjectID(),
		Course_name: course.Course_name,
		Credit:      course.Credit,
		Class_id:    course.Class_id,
	}
	collection = config.MongoClient.Database(dbName).Collection(dbCollection)

	_, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newUser)
}
