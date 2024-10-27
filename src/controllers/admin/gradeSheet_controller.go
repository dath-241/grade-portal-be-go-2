package controllers_admin

import (
	"be-go-2/config"
	"be-go-2/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func CreateGradeSheet(c *gin.Context) {
	var gradeSheet models.GradeSheet

	if err := c.ShouldBindJSON(&gradeSheet); err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	gradeSheet.StudentID = primitive.NewObjectID() // Create a new ObjectID
	gradeSheet.ClassID = primitive.NewObjectID()   // Create a new ObjectID
	gradeSheet.CreatedAt = time.Now()
	gradeSheet.UpdatedAt = time.Now()
	collection := config.MongoClient.Database("School").Collection("GradeSheet")

	if _, err := collection.InsertOne(c.Request.Context(), gradeSheet); err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusCreated, gradeSheet)
}
