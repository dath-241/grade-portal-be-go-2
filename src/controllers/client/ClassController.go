package controller_client

import (
	"Go2/models"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// HandleTeacherClasses xử lý việc lấy danh sách lớp học của giáo viên.
func HandleTeacherClasses(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	if user.Role != "teacher" {
		c.JSON(401, gin.H{
			"status":    "Fail",
			"message": "Chỉ giáo viên mới được phép truy cập",
		})
		return
	}
	var classTeacherAll []models.InterfaceClass
	collection := models.ClassModel()
	cursor, err := collection.Find(context.TODO(), bson.M{
		"teacher_id": user.ID,
	})

	if err != nil {
		c.JSON(401, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi tìm lớp học",
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &classTeacherAll); err != nil {
		c.JSON(401, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi đọc dữ liệu lớp học",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":     "Success",
		"data": classTeacherAll,
	})
}

// HandleStudentClasses xử lý việc lấy danh sách lớp học của sinh viên.
func HandleStudentClasses(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	var classStudentAll []models.InterfaceClassStudent
	collection := models.ClassModel()
	fmt.Println(user)
	cursor, err := collection.Find(context.TODO(), bson.M{
		"listStudent_ms": user.Ms,
	})
	if err != nil {
		c.JSON(401, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi tìm lớp học",
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &classStudentAll); err != nil {
		c.JSON(401, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi đọc dữ liệu lớp học",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":     "Success",
		"data": classStudentAll,
	})
}

// HandleUserClasses xử lý việc lấy danh sách lớp học của người dùng.
func HandleUserClasses(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	if user.Role == "teacher" {
		HandleTeacherClasses(c)
		return
	} else if user.Role == "student" {
		HandleStudentClasses(c)
		return
	}
	c.JSON(400, gin.H{
		"status":    "Fail",
		"message": "Không tìm thấy người dùng",
	})
}

// HandleClassDetail xử lý việc lấy chi tiết lớp học.
func HandleClassDetail(c *gin.Context) {
	paramID := c.Param("id")
	id, _ := bson.ObjectIDFromHex(paramID)
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	var classDetail models.InterfaceClass
	collection := models.ClassModel()
	err := collection.FindOne(context.TODO(), bson.M{
		"_id": id,
	}).Decode(&classDetail)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Không tìm thấy lớp học",
		})
		return
	}
	if user.Role == "student" {
		var listStudent = classDetail.ListStudentMs
		for _, studentMs := range listStudent {
			if studentMs == user.Ms {
				c.JSON(200, gin.H{
					"status":        "Success",
					"message": "Lấy lớp học thành công",
					"data": classDetail,
				})
				return
			}
		}
		c.JSON(401, gin.H{
			"status":    "Fail",
			"message": "Chỉ sinh viên mới được phép truy cập",
		})
		return
	} else if user.Role == "teacher" {
		if user.ID != classDetail.TeacherId {
			c.JSON(401, gin.H{
				"status":    "Fail",
				"message": "Chỉ giáo viên mới được phép truy cập",
			})
			return
		}
		c.JSON(200, gin.H{
			"status":        "Success",
			"message": "Lấy lớp học thành công",
			"data": classDetail,
		})
		return
	}
	c.JSON(401, gin.H{
		"status":    "Fail",
		"message": "Chỉ sinh viên và giáo viên mới được phép truy cập",
	})
}

// HandleCountDocuments xử lý việc đếm số lượng lớp học của một môn học.
func HandleCountDocuments(c *gin.Context) {
	param := c.Param("id")
	courseID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Không tìm thấy môn học",
		})
		return
	}
	collection := models.ClassModel()
	count, err := collection.CountDocuments(context.TODO(), bson.M{"course_id": courseID})
	if err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Lỗi khi lấy môn học",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "Success",
		"data": count,
	})
}
