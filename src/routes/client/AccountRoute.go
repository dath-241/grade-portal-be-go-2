package routes_client

import (
    controller_client "Go2/controllers/client"
    middlewares_client "Go2/middlewares/client"

    "github.com/gin-gonic/gin"
)

// AccountRoute thiết lập các route cho việc quản lý tài khoản người dùng
func AccountRoute(r *gin.RouterGroup) {
    // Route để đăng nhập
    r.POST("/login", controller_client.HandleLogin)

    // Route để đăng xuất, yêu cầu xác thực người dùng
    r.POST("/logout", middlewares_client.RequireUser, controller_client.HandleLogout)

    // Route để lấy thông tin tài khoản hiện tại, yêu cầu xác thực người dùng
    r.GET("/info", middlewares_client.RequireUser, controller_client.HandleAccount)

    // Route để lấy thông tin giáo viên theo ID
    r.GET("/:id", controller_client.HandleGetInfoByID)
}