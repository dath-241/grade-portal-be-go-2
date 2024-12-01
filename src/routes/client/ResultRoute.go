package routes_client

import (
    controller_client "Go2/controllers/client"
    middlewares_client "Go2/middlewares/client"

    "github.com/gin-gonic/gin"
)

// ResultScoreRoute thiết lập các route cho kết quả điểm
func ResultRoute(r *gin.RouterGroup) {
    // Tạo kết quả điểm, yêu cầu quyền giáo viên
    r.POST("/create", middlewares_client.RequireTeacher, controller_client.HandleCreateResult)
    // Lấy tất cả kết quả điểm
    r.GET("/getmark", controller_client.HandleAllResults)
    // Lấy kết quả điểm của một khóa học cụ thể
    r.GET("/getmark/:ms", controller_client.HandleCourseResult)
    // Lấy kết quả điểm theo ID
    r.GET("/:id", controller_client.HandleResult)
    // Cập nhật kết quả điểm theo ID, yêu cầu quyền giáo viên
    r.PATCH("change", middlewares_client.RequireTeacher, controller_client.HandlePatchResult)
}