package controller_admin

import (
	"github.com/gin-gonic/gin"
)

// HandleCreateAccount xử lý việc tạo tài khoản mới.
func HandleCreateAccount(c *gin.Context) {

}

// CheckEmailAndRole kiểm tra đuôi email và role
func CheckEmailAndRole(email string, role string) bool {
	return true
}

// HandleGetAccountByID xử lý việc lấy thông tin tài khoản theo ID.
func HandleGetAccountByID(c *gin.Context) {

}

// HandleGetTeacherAccounts xử lý việc lấy thông tin tài khoản giáo viên.
func HandleGetTeacherAccounts(c *gin.Context) {

}

// HandleGetStudentAccounts xử lý việc lấy thông tin tài khoản sinh viên.
func HandleGetStudentAccounts(c *gin.Context) {

}

// HandleDeleteAccount xử lý việc xóa tài khoản.
func HandleDeleteAccount(c *gin.Context) {

}

// HandleUpdateAccount xử lý việc cập nhật thông tin tài khoản.
func HandleUpdateAccount(c *gin.Context) {

}
