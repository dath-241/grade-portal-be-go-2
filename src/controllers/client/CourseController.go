package controller_client

import (
	"Go2/models"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// HandleGetCourseByID xử lý việc lấy môn học theo mã môn học.
func HandleGetCourseByID(c *gin.Context) {
	param := c.Param("id")
	courseID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}

	var course models.InterfaceCourse
	collection := models.CourseModel()

	err = collection.FindOne(context.TODO(), bson.M{"_id": courseID}).Decode(&course)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"status":  "Fail",
				"message": "Không tìm thấy khóa học",
			})
			return
		}
		c.JSON(500, gin.H{
			"status":  "Fail",
			"message": "Lỗi khi lấy khóa học",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Lấy khóa học thành công",
		"data":    course,
	})
}
