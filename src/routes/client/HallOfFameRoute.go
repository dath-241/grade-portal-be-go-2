package routes_client

import (
    controller_client "Go2/controllers/client"

    "github.com/gin-gonic/gin"
)

// HallOfFameRoute thiết lập route cho Hall of Fame
func HallOfFameRoute(r *gin.RouterGroup) {
    // Route để lấy tất cả Hall of Fame của học kỳ trước
    r.GET("/all", controller_client.GetAllPrevSemester)
}