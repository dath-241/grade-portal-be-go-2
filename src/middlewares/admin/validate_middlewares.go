package middlewares_admin

import (
	controller_admin "Go2/controllers/admin"
	"strings"

	"github.com/gin-gonic/gin"
)

// ValidateEmail kiểm tra email có đuôi @hcmut.edu.vn
func ValidateEmail(email string) bool {
	return strings.HasSuffix(email, "@hcmut.edu.vn")
}

// ValidateMS kiểm tra mã số sinh viên không rỗng
func ValidateMS(ms string) bool {
	return ms != ""
}

// ValidateDataAdmin kiểm tra dữ liệu admin
func ValidateDataAdmin(c *gin.Context) {
	var data controller_admin.InterfaceAdminController
	c.BindJSON(&data)
	if !ValidateEmail(data.Email) || !ValidateMS(data.Ms) {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu không hợp lệ",
			"data":    data,
		})
		c.Abort()
		return
	}
	c.Set("adminData", data)
	c.Next()
}
