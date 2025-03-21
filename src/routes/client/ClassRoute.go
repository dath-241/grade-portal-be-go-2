package routes_client

import (
	controller_client "Go2/controllers/client"
	middlewares_client "Go2/middlewares/client"

	"github.com/gin-gonic/gin"
)

// ClassRoute thiết lập các route cho việc quản lý lớp học
func ClassRoute(r *gin.RouterGroup) {
	// Route để lấy ra tất cả các class của account đó
	r.GET("/account", controller_client.HandleUserClasses)

	// Route để lấy ra chi tiết lớp học
	r.GET("/:id", controller_client.HandleClassDetail)

	// Route để đếm số lượng lớp học của một môn học
	r.GET("/count/:id", controller_client.HandleCountDocuments)

	// Route để giảng viên tạo lớp học kèm theo link csv bảng điểm
	r.POST("/create", middlewares_client.RequireTeacher, controller_client.HandleAddClass)

	// Route để giảng viên update link csv bảng điểm
	r.PATCH("/upload/:id", middlewares_client.RequireTeacher, controller_client.HandleUpdateClassCsvURL)

	// Route để giảng viên xóa lớp học
	r.DELETE("/delete/:id", middlewares_client.RequireTeacher, controller_client.HandleDeleteClass)
}
