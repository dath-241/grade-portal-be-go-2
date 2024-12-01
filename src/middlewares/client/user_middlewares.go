package middlewares_client

import (
	"Go2/helper"
	"Go2/models"
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func RequireUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{
			"status": "Fail",
			"message": "Yêu cầu cung cấp token"})
		c.Abort()
		return
	}

	// Kiểm tra định dạng Bearer token
	if !strings.HasPrefix(token, "Bearer ") {
		c.JSON(401, gin.H{
			"status": "Fail",
			"message": "Header Authorization không hợp lệ"})
		c.Abort()
		return
	}

	token = token[7:] // Loại bỏ tiền tố "Bearer "
	claims, err := helper.ParseJWT(token)
	if err != nil || claims == nil {
		c.JSON(401, gin.H{
			"status": "Fail",
			"message": "Người dùng chưa đăng nhập !"})
		c.Abort()
		return
	}

	var user models.InterfaceAccount
	collection := models.AccountModel()
	err = collection.FindOne(context.TODO(), bson.M{"_id": claims.ID}).Decode(&user)
	if err != nil {
		c.JSON(401, gin.H{
			"status": "Fail",
			"message": "Không tìm thấy người dùng"})
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}

func RequireTeacher(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"status": "Fail",
			"message": "Người dùng chưa đăng nhập !"})
		c.Abort()
		return
	}

	account := user.(models.InterfaceAccount)
	if account.Role != "teacher" {
		c.JSON(403, gin.H{
			"status": "Fail",
			"message": "Chỉ giáo viên mới có quyền truy cập"})
		c.Abort()
		return
	}

	c.Next()
}
