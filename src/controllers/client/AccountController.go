package controller_client

import (
	"Go2/helper"
	"Go2/models"
	"context"
	"os"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// HandleLogin xử lý việc đăng nhập.
func HandleLogin(c *gin.Context) {
	var body InterfaceAccount
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Invalid data",
		})
		return
	}
	payload, err := idtoken.Validate(context.Background(), body.IDToken, os.Getenv("YOUR_CLIENT_ID"))
	if err != nil {
		c.JSON(401, gin.H{"error": "Token không hợp lệ"})
		return
	}

	email, validEmail := payload.Claims["email"].(string)
	if !validEmail {
		c.JSON(400, gin.H{"error": "Không lấy được thông tin người dùng"})
		return
	}

	collection := models.AccountModel()
	var user models.InterfaceAccount
	err = collection.FindOne(context.TODO(), bson.M{"email": email}).Decode((&user))
	if err != nil {
		c.JSON(400, gin.H{"error": "Không lấy được thông tin người dùng trong dữ liệu."})
		return
	}
	token := helper.CreateJWT(user.ID)
	c.SetCookie("token", token, 3600*24, "/", "", false, true)
	c.JSON(200, gin.H{
		"code":  "Success",
		"token": token,
		"role":  user.Role,
	})
}

// HandleLogout xử lý việc đăng xuất.
func HandleLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{
		"code":    "Success",
		"message": "Đăng xuất thành công",
	})
}

// HandleGetInfoByID xử lý việc lấy thông tin giáo viên theo ID.
func HandleGetInfoByID(c *gin.Context) {
	param := c.Param("id")
	teacherID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    "error",
			"message": "Teacher ID sai",
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
		c.JSON(400, gin.H{
			"code":    "error",
			"message": "Teacher ID sai",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "success",
		"teacher": teacher,
	})
}

// handleAccount lấy thông tin tài khoản hiện tại.
func HandleAccount(c *gin.Context) {
	user, _ := c.Get("user")
	if user == "" {
		c.JSON(401, gin.H{
			"code":    "error",
			"message": "Không có người dùng",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "success",
		"user": user,
	})
}
