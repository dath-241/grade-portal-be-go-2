package controller_client

import (
	"Go2/models"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// HandleLogin xử lý việc đăng nhập.
func HandleLogin(c *gin.Context) {

}

// HandleLogout xử lý việc đăng xuất.
func HandleLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Đăng xuất thành công",
	})
}

// HandleAccount lấy thông tin tài khoản hiện tại.
func HandleAccount(c *gin.Context) {
	user, _ := c.Get("user")
	if user == "" {
		c.JSON(401, gin.H{
			"status":  "Fail",
			"message": "Yêu cầu đăng nhập",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Thành công",
		"data":    user,
	})
}

// HandleGetInfoByID xử lý việc lấy thông tin giáo viên theo ID.
func HandleGetInfoByID(c *gin.Context) {
	param := c.Param("id")
	teacherID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	collection := models.AccountModel()
	var teacher struct {
		Name  string `bson:"name"`
		Email string `bson:"email"`
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": teacherID, "role": "teacher"}).Decode(&teacher)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"status":  "Fail",
				"message": "Không tìm thấy giảng viên",
			})
			return
		}
		c.JSON(500, gin.H{
			"status":  "Fail",
			"message": "Lỗi khi truy vấn dữ liệu",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Thành công",
		"data":    teacher,
	})
}
