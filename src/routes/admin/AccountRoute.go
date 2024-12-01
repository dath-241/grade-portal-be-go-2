package routes_admin

import (
	controller_admin "Go2/controllers/admin"

	"github.com/gin-gonic/gin"
)

// AccountRoute thiết lập các route cho tài khoản
func AccountRoute(r *gin.RouterGroup) {
	// Route để tạo danh sách các tài khoản mới
	r.POST("/create", controller_admin.HandleCreateAccount)

	// Route để tìm tài khoản theo ID
	r.GET("/:id", controller_admin.HandleGetAccountByID)

	// Route để lấy tất cả tài khoản giáo viên
	r.GET("/teacher", controller_admin.HandleGetTeacherAccounts)

	// Route để lấy tất cả tài khoản sinh viên
	r.GET("/student", controller_admin.HandleGetStudentAccounts)

	// Route để xóa tài khoản theo ID
	r.DELETE("/delete/:id", controller_admin.HandleDeleteAccount)

	// Route để cập nhật thông tin tài khoản theo ID
	r.PATCH("/change/:id", controller_admin.HandleUpdateAccount)
}
