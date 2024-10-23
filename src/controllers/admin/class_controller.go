package controllers_admin

import (
	"be-go-2/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateClass(c *gin.Context) {
	var data  struct{
		Semester      string   `json:"semester"`
		Name          string   `json:"name"` // nhom lop
		CourseId      string   `json:"course_id"`
		ListStudentId []string `json:"listStudent_id"`
		TeacherId     string   `json:"teacher_id"`
	}

	//bindding data xem co loi khong
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			// "message": "data error",
		})
		return
	}

	//check class da ton tai chua
	courseID, err := bson.ObjectIDFromHex(data.CourseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			// "message": "course id error",
		})
		return
	}

	//get collection
	collection := models.ClassModel()
	
	// check duplicate
	err = collection.FindOne(context.Background(), bson.M{
		"Semester": data.Semester,
		"CourseId": courseID,
		"Name": data.Name,
	}).Decode(&data)
	if err == mongo.ErrNoDocuments{
	}else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			// "message": err.Error(),
			"message": "search error",
		})
		return
	}else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Class adready exists",
		})
		return
	}

	// create new class
	id, _ := c.Get("ID")
	
	result, err := collection.InsertOne(context.Background(), bson.M{
		"Name": data.Name,
		"CourseId": courseID,
		"Semester": data.Semester,
		"ListStudentId": data.ListStudentId,
		"TeacherId": data.TeacherId,
		"CreatedBy": id,
		"UpdatedBy": id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			// "message": err.Error(),
			"message": "insert error",
		})
	}
	// return success
	c.JSON(http.StatusOK, gin.H{
		"message": "Created class successfully",
		"class_id": result.InsertedID,
	})
}
