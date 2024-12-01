package routes_client

import (
    controller_client "Go2/controllers/client"

    "github.com/gin-gonic/gin"
)

// HomeRouter thiết lập route cho trang chủ
func HomeRouter(r *gin.RouterGroup) {
    // Route cho trang chủ
    r.GET("/", controller_client.HomePage)
}