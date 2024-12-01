package routes_admin

import (
	controller_admin "Go2/controllers/admin"

	"github.com/gin-gonic/gin"
)

// ClassRoute thiết lập các route cho việc quản lý lớp học
func ClassRoute(r *gin.RouterGroup) {
	// Route để tạo 1 lớp học mới
	r.POST("/create", controller_admin.HandleCreateClass)

	// Route để lấy ra chi tiết lớp học
	r.GET("/:id", controller_admin.HandleGetClassByID)

	// Route để lấy ra tất cả lớp học của id account đó
	r.GET("/account/:id", controller_admin.HandleGetAllClassesByAccountID)

	// Route để lấy ra tất cả lớp học của id course đó
	r.GET("/course/:id", controller_admin.HandleGetClassesByCourseID)

	// Route để thêm học sinh vào lớp học đó
	r.PATCH("/add", controller_admin.HandleAddStudentsToClass)

	// Route để xóa lớp học theo id lớp học
	r.DELETE("/delete/:id", controller_admin.HandleDeleteClass)

	// Route để chỉnh sửa thông tin lớp học
	r.PATCH("/change/:id", controller_admin.HandleUpdateClass)
}
