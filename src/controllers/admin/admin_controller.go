package controllers_admin

import (
	"be-go-2/models"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IAdminCreate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Faculty  string `json:"faculty"`
	Role     string `json:"role"`
}

func IsAdminExist(collection *mongo.Collection, email string, passowrd string, name string) (bool, error) {

	filter := bson.M{
		"email":    email,
		"password": passowrd,
		"name":     name,
	}

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

	var body IAdminCreate // body là biến mang cấu trúc của IAdminCreate

	// Kiểm tra request body
	if err := c.BindJSON(&body); err != nil { // Nếu request body có cấu trúc không đúng với IAdminCreate
		c.JSON(400, gin.H{ // Response lỗi dữ liệu không hợp lệ
			"success": false,
			"message": "Dữ liệu không hợp lệ",
			"data":    nil,
		})
		return
	}

	// Request body đã đúng cấu trúc
	// Truy suất vào collection
	collection := models.AdminModel()

	// Kiểm tra trùng email, password, name
	isExis, err := IsAdminExist(collection, body.Email, body.Password, body.Name)

	if err != nil {
		c.JSON(500, gin.H{ // Response lỗi khi kiểm tra trùng
			"success": false,
			"message": "Đã có lỗi xảy ra",
			"data":    nil,
		})
		return
	}
	if isExis {
		c.JSON(500, gin.H{ // Response lỗi dữ liệu đã tồn tại
			"success": false,
			"message": "Admin đã tồn tại",
			"data":    nil,
		})
		return
	}

	// Thêm nếu không có các lỗi trên
	_, err = collection.InsertOne(context.TODO(), bson.M{
		"email":    body.Email,
		"password": body.Password,
		"name":     body.Name,
		"faculty":  body.Faculty,
		"role":     body.Role,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Đã có lỗi xảy ra khi tạo Admin",
			"data":    nil,
		})
		return
	}

	// Tạo Admin thành công
	c.JSON(200, gin.H{
		"success": true,
		"message": "Thêm Admin thành công",
		"data": IAdminCreate{
			Email:    body.Email,
			Password: body.Password,
			Name:     body.Name,
			Faculty:  body.Faculty,
			Role:     body.Role,
		},
	})
}
