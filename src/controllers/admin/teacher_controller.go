package controllers_admin

import (
	"be-go-2/config"
	"be-go-2/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gán giá trị bổ sung cho teacher
	teacher.ID = primitive.NewObjectID()
	teacher.Role = "teacher"
	teacher.CreatedAt = time.Now()
	teacher.ExpiredAt = time.Now()

	// Truy xuất collection
	teacherCollection := config.MongoClient.Database("grade-portal").Collection("teacher")

	// Thêm Teacher vào database
	_, err := teacherCollection.InsertOne(context.TODO(), teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, teacher)
}
