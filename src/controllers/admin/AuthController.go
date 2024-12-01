package controller_admin

import (
	"Go2/helper"
	"Go2/models"
	"context"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// HandleLogin xử lý việc đăng nhập.
func HandleLogin(c *gin.Context) {
	var loginData AuthController
	// Lấy dữ liệu từ front end
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Dữ liệu yêu cầu không hợp lệ"})
		return
	}
	payload, err := idtoken.Validate(context.Background(), loginData.IDToken, os.Getenv("YOUR_CLIENT_ID"))
	if err != nil {
		fmt.Println("Không có token:", err)
		c.JSON(401, gin.H{
			"status":  "Fail",
			"message": "Token không hợp lệ"})
		return
	}
	// Lấy ra email
	email, emailOk := payload.Claims["email"].(string)
	if !emailOk {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Không lấy được thông tin người dùng"})
		return
	}
	// Tìm kiếm người dùng đã có trong database không
	collection := models.AdminModel()
	var user models.InterfaceAdmin
	err = collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "Fail",
			"mesage": "Không lấy được thông tin người dùng trong cơ sở dữ liệu"})
		return
	}
	token := helper.CreateJWT(user.ID)
	c.SetCookie("token", token, 3600*24, "/", "", true, true)
	c.JSON(200, gin.H{
		"status":  "Success",
		"token": token,
	})
}

// HandleLogout xử lý việc đăng xuất.
func HandleLogout(c *gin.Context) {
	c.SetCookie("token", "", 3600*24, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{
		"status":    "Success",
		"message": "Đăng xuất thành công",
	})
}

// HandleCreateAdmin xử lý việc tạo tài khoản admin mới.
func HandleCreateAdmin(c *gin.Context) {
	//Lấy admindata từ mdw
	adminData, _ := c.Get("adminData")
	data := adminData.(InterfaceAdminController)
	collection := models.AdminModel()
	//Lấy ID người tạo từ mdw
	createdBy, _ := c.Get("ID")
	//Check xem admin đó đã tồn tại hay chưa
	var existingAdmin models.InterfaceAdmin
	err := collection.FindOne(context.TODO(), bson.M{
		"$or": bson.A{
			bson.M{"email": data.Email},
			bson.M{"ms": data.Ms},
		},
	}).Decode(&existingAdmin)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":    "Fail",
			"message": "Bảng ghi của admin này đã được lưu trong database trước đó",
		})
		return
	}
	//Thêm admin và cơ sở dữ liệu
	collection.InsertOne(context.TODO(), bson.M{
		"email":     data.Email,
		"name":      data.Name,
		"faculty":   data.Faculty,
		"ms":        data.Ms,
		"createdBy": createdBy,
	})
	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Tạo tài khoản admin thành công",
	})
}

// HandleProfile xử lý việc lấy thông tin tài khoản admin.
func HandleProfile(c *gin.Context) {
	ID, _ := c.Get("ID")
	//Lấy thông tin admin từ cơ sở dữ liệu
	collection := models.AdminModel()
	var user models.InterfaceAdmin
	err := collection.FindOne(context.TODO(), bson.M{"_id": ID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Fail",
			"message": "Không lấy được thông tin người dùng trong dữ liệu.",
		})
		return
	}
	//Trả kết quả
	c.JSON(http.StatusOK, gin.H{
		"status":    "Success",
		"message": "Thành công",
		"data":    user,
	})
}
