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

type IStudentCreate struct {
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Name     string        `json:"name"`
	UID      string        `json:"UID"`
	Faculty  string        `json:"faculty"`
	ClassID  []interface{} `json:"classID"`
}

func IsStudentExist(collection *mongo.Collection, UID string, Email string) (bool, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"UID": UID},
			{"email": Email},
		},
	}

	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments { // Không có trùng lặp
		return false, nil
	} else if err != nil { // Lỗi khác
		return false, err
	}

	return true, nil
}

func CreateStudent(c *gin.Context) {
	var body IStudentCreate
	// Kiểm tra body của request có hợp lệ không
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Dữ liệu không hợp lệ",
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

	collection := models.StudentModel()

	isExist, err := IsStudentExist(collection, body.UID, body.Email)

	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Đã có lỗi xảy ra",
			"data":    nil,
		})
		return
	}

	if isExist {
		c.JSON(500, gin.H{
			"success": false,
			"message": "MSSV hoặc Email đã tồn tại",
			"data":    nil,
		})
		return
	}

	_, err = collection.InsertOne(context.TODO(), bson.M{
		"email":     body.Email,
		"password":  body.Password,
		"name":      body.Name,
		"UID":       body.UID,
		"faculty":   body.Faculty,
		"role":      "student",
		"createdBy": time.Now(),
		"createdAt": time.Now(),
		"expiredAt": time.Now(),
	})

	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Đã có lỗi xảy ra",
			"data":    nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Tạo thành công",
		"data": bson.M{
			"email":     body.Email,
			"password":  body.Password,
			"name":      body.Name,
			"UID":       body.UID,
			"faculty":   body.Faculty,
			"role":      "student",
			"createdBy": primitive.NewObjectID(),
			"createdAt": time.Now(),
			"expiredAt": time.Now(),
		},
	})
}
