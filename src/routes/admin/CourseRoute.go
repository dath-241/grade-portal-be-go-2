package routes_admin

import (
	controller_admin "Go2/controllers/admin"

	"github.com/gin-gonic/gin"
)

// CourseRoute thiết lập các route cho việc quản lý khóa học
func CourseRoute(r *gin.RouterGroup) {
	// Route để tạo khóa học mới
	r.POST("/create", controller_admin.HandleCreateCourse)

	// Route để lấy chi tiết khóa học theo ID
	r.GET("/:id", controller_admin.HandleGetCourseByID)

	// Route để lấy ra tất cả khóa học
	r.GET("/all", controller_admin.HandleGetAllCourses)

	// Route để xóa khóa học theo ID
	r.DELETE("/delete/:id", controller_admin.HandleDeleteCourse)

	// Route để cập nhật thông tin khóa học theo ID
	r.PATCH("/change/:id", controller_admin.HandleUpdateCourse)

}
