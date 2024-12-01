package routes_admin

import (
	controller_admin "Go2/controllers/admin"
	middlewares_admin "Go2/middlewares/admin"

	"github.com/gin-gonic/gin"
)

// AuthRoute thiết lập các route cho việc xác thực và quản lý tài khoản admin
func AuthRoute(r *gin.RouterGroup) {
	// Route để đăng nhập
	r.POST("/login", controller_admin.HandleLogin)

	// Route để đăng xuất, yêu cầu xác thực
	r.POST("/logout", middlewares_admin.HandleRequireAuth, controller_admin.HandleLogout)

	// Route để tạo tài khoản admin mới, yêu cầu xác thực và xác thực dữ liệu admin
	r.POST("/create", middlewares_admin.HandleRequireAuth, middlewares_admin.ValidateDataAdmin, controller_admin.HandleCreateAdmin)

	// Route để lấy thông tin profile của admin hiện tại, yêu cầu xác thực
	r.GET("/profile", middlewares_admin.HandleRequireAuth, controller_admin.HandleProfile)

}
