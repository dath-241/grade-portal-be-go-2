package routes_client

import (
    middlewares_client "Go2/middlewares/client"

    "github.com/gin-gonic/gin"
)

// MainRoute thiết lập các route chính cho ứng dụng
func MainRoute(r *gin.Engine) {
    // Route cho trang chủ
    HomeRouter(r.Group("/"))
    // Route cho tài khoản
    AccountRoute(r.Group("/api"))
    
    // Nhóm các route yêu cầu quyền người dùng
    protectedGroup := r.Group("/api")
    protectedGroup.Use(middlewares_client.RequireUser)
    
    // Route cho Hall of Fame
    HallOfFameRoute(protectedGroup.Group("/HOF"))
    // Route cho lớp học
    ClassRoute(protectedGroup.Group("/class"))
    // Route cho khóa học
    CourseRoute(protectedGroup.Group("/course"))
    // Route cho kết quả điểm
    ResultRoute(protectedGroup.Group("/result"))
}