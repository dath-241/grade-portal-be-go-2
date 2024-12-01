package routes_admin

import (
	"github.com/gin-gonic/gin"

	controller_admin "Go2/controllers/admin"
)

// HallOfFameRoute thiết lập các route cho việc quản lý Hall of Fame
func HallOfFameRoute(r *gin.RouterGroup) {
	// Route để cập nhật Hall of Fame
	r.POST("/update", controller_admin.HandleCreateHallOfFame)
	// Route để lấy tất cả Hall of Fame của học kỳ trước
	r.GET("/all", controller_admin.HandleGetPrevSemesterAllHallOfFame)
}
