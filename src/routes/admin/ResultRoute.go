package routes_admin

import (
	controller_admin "Go2/controllers/admin"

	"github.com/gin-gonic/gin"
)

// ResultScoreRoute thiết lập các route cho việc quản lý bảng điểm
func ResultRoute(r *gin.RouterGroup) {
	// Route để tạo bảng điểm mới
	r.POST("/create", controller_admin.HandleCreateResult)

	// Route để lấy bảng điểm theo ID lớp học
	r.GET("/:id", controller_admin.HandleGetResult)
}
