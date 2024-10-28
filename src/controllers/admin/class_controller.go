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

	//binding data xem co loi khong
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"position": "binding data error",
		})
		return
	}
	// kiem tra student id có hợp lệ không
	studentCollection := models.StudentModel()

	for _, id := range data.ListStudentId {
        stuID, err := bson.ObjectIDFromHex(id)
        if err!= nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "message": err.Error(),
                "position": "student id error",
            })
            return
        }
		var student bson.M
		err = studentCollection.FindOne(context.Background(), bson.M{
			"_id": stuID,
		}).Decode(&student)
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Student id not found",
                "position": "student id error",
			})
		}else if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
                "position": "student id error",
			})
		}
    }
    // kiem tra teacher id co ton tai khong
    
	// kiem tra courseID co ton tai khong
	courseID, err := bson.ObjectIDFromHex(data.CourseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"position": "course id error",
		})
		return
	}
	var course bson.M
	courseCollection := models.CourseModle()
	err = courseCollection.FindOne(context.Background(), bson.M{
		"_id": courseID,
	}).Decode(&course)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{
            "message": "Course id not found",
            "position": "course id error",
        })
        return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "message": err.Error(),
            "position": "course id error",
        })
        return
	}
	//	get class collection
	collection := models.ClassModel()
	
	// check duplicate class
	var class bson.M
	err = collection.FindOne(context.Background(), bson.M{
		"Semester": data.Semester,
		"CourseId": courseID,
		"Name": data.Name,
	}).Decode(&class)
	if err == mongo.ErrNoDocuments{
	}else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"position": "search class error",
		})
		return
	}else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Class adready exists",
		})
		return
	}

	// check if the middleware has set adminId
	id, exist := c.Get("ID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID is required, need set ID in middleware",
            "position": "id error",
		})
		return
	}
	// create new class
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
			"message": err.Error(),
			"position": "insert class error",
		})
	}
	// return success
	c.JSON(http.StatusOK, gin.H{
		"message": "Created class successfully",
		"class_id": result.InsertedID,
	})
}
