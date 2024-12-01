package routes_client

import (
    controller_client "Go2/controllers/client"

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
}