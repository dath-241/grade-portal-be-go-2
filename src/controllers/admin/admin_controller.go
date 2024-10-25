package controllers_admin

import (
	"be-go-2/models"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IAdminCreate struct {
	Name struct {
		LastName string `json:"LastName"`
		MFName   string `json:"MFName"`
	} `json:"name"`
	Email string `json:"email"`
}

func IsAdminExist(collection *mongo.Collection, email string) (bool, error) {

	filter := bson.M{
		"email": email,
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

	var body IAdminCreate
	var data User

	// Kiểm tra cấu trúc dữ liệu
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Dữ liệu không hợp lệ",
			"data":    nil,
		})
		return
	}

	// Kiểm tra các trường bắt buộc: FirstName, MFName, Email
	if strings.TrimSpace(body.Name.LastName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Họ không được để trống",
			"data":    nil,
		})
		return
	}

	if strings.TrimSpace(body.Name.MFName) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Tên không được để trống",
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

	// Kiểm tra trùng Email
	isExist, err := IsAdminExist(collection, body.Email)

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
			"message": "Email đã được sử dụng",
			"data":    nil,
		})
		return
	}

	// Thêm Admin
	data.ID = primitive.NewObjectID().Hex()
	data.Name.LastName = body.Name.LastName
	data.Name.MFName = body.Name.MFName
	data.Email = body.Email
	data.Role = "Admin"
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.AdminInfo = &AdminInfo{
		AdminID: data.ID,
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
