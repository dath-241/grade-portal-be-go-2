package routes_client

import (
    controller_client "Go2/controllers/client"

    "github.com/gin-gonic/gin"
)

// CourseRoute thiết lập các route cho việc quản lý khóa học
func CourseRoute(r *gin.RouterGroup) {
    // Route để lấy ra chi tiết khóa học
    r.GET("/:id", controller_client.HandleGetCourseByID)
}