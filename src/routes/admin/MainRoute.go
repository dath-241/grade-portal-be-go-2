package routes_admin

import (
	"Go2/config"
	middlewares_admin "Go2/middlewares/admin"

	"github.com/gin-gonic/gin"
)

// MainRoute thiết lập các route chính cho ứng dụng
func MainRoute(r *gin.Engine) {
	prefixAdmin := config.BASE_URL_ADMIN()

	// Các route không yêu cầu đăng nhập
	AuthRoute(r.Group(prefixAdmin)) // admin/api

	// Tạo nhóm route yêu cầu xác thực
	protectedGroup := r.Group(prefixAdmin)
	protectedGroup.Use(middlewares_admin.HandleRequireAuth)

	// Các route yêu cầu đăng nhập
	ResultRoute(protectedGroup.Group("/result"))   // Route cho bảng điểm
	AccountRoute(protectedGroup.Group("/account")) // Route cho tài khoản
	ClassRoute(protectedGroup.Group("/class"))     // Route cho lớp học
	CourseRoute(protectedGroup.Group("/course"))   // Route cho khóa học
	HallOfFameRoute(protectedGroup.Group("/HOF"))  // Route cho Hall of Fame
}
