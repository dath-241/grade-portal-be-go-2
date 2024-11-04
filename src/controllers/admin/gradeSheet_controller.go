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
	courseIDStr := c.Param("courseID")
	classIDStr := c.Param("classID")
	// Convert string IDs to ObjectID
	courseID, err := primitive.ObjectIDFromHex(courseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	classID, err := primitive.ObjectIDFromHex(classIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var gradeSheet models.GradeSheet

	if err := c.ShouldBindJSON(&gradeSheet); err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	gradeSheet.CourseID = courseID
	gradeSheet.ClassID = classID
	gradeSheet.ExpiredAt = time.Now().AddDate(1, 0, 0)
	collection := config.MongoClient.Database("School").Collection("GradeSheet")

	if _, err := collection.InsertOne(c.Request.Context(), gradeSheet); err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusCreated, gradeSheet)
}
