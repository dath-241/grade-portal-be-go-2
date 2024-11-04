package controllers_admin

import (
	"be-go-2/models"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IAdminCreate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func IsAdminExist(collection *mongo.Collection, email string, password string) (bool, error) {

	filter := bson.M{
		"$or": []bson.M{
			{"email": email},
			{"password": password},
		}}

	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments { // Không có trùng lặp
		return false, nil
	} else if err != nil { // Lỗi khác
		return false, err
	}

	return true, nil // Có trùng lặp
}

func CreateAdmin(c *gin.Context) {

	var body IAdminCreate

	// Kiểm tra cấu trúc dữ liệu
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Dữ liệu không hợp lệ",
			"data":    nil,
		})
		return
	}

	// Kiểm tra các trường bắt buộc: Name, Email, Password
	if strings.TrimSpace(body.Name) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Họ và tên không được để trống",
			"data":    nil,
		})
		return
	}

	if strings.TrimSpace(body.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Mật khẩu không được để trống",
			"data":    nil,
		})
		return
	}

	if strings.TrimSpace(body.Email) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Email không được để trống",
			"data":    nil,
		})
		return
	}

	// Kiểm tra Email có hậu tố `@hcmut.edu.vn`
	if !strings.HasSuffix(body.Email, "@hcmut.edu.vn") {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Email phải là @hcmut.edu.vn",
			"data":    nil,
		})
		return
	}

	collection := models.AdminModel()

	// Kiểm tra trùng Email, Password
	isExist, err := IsAdminExist(collection, body.Email, body.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Đã có lỗi xảy ra",
			"data":    nil,
		})
		return
	}

	if isExist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Email hoặc mật khẩu đã được sử dụng",
			"data":    nil,
		})
		return
	}

	// Thêm Admin

	var data = models.InterfaceAdmin{
		Email:     body.Email,
		Password:  body.Password,
		Name:      body.Name,
		CreatedAt: time.Now(),
	}

	_, err = collection.InsertOne(context.TODO(), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Đã có lỗi xảy ra khi thêm Admin",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Thêm Admin thành công",
		"data":    data,
	})
}
