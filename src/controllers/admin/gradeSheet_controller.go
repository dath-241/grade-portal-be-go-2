package controllers_admin

import (
	"be-go-2/config"
	"be-go-2/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetGradeSheet(c *gin.Context) {
	courseIDStr := c.Param("courseID")
	classIDStr := c.Param("classID")

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
	collection := config.MongoClient.Database("School").Collection("GradeSheet")

	// Query to find the document with matching courseID and classID
	var gradeSheet models.GradeSheet
	err = collection.FindOne(c.Request.Context(), bson.M{"courseID": courseID, "classID": classID}).Decode(&gradeSheet)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"message": "GradeSheet not found"})
		} else {
			c.JSON(http.StatusInternalServerError, errResponse(err))
		}
		return
	}

	// Return the retrieved GradeSheet document as JSON
	c.JSON(http.StatusOK, gradeSheet)
}
