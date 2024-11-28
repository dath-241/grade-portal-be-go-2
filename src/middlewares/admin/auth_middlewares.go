package middlewares_admin

import (
	// "Go2/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

// HandleRequireAuth kiểm tra và xác thực token của người dùng
func HandleRequireAuth(c *gin.Context) {
    // Lấy giá trị của header Authorization
    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(401, gin.H{
            "status": "Fail",
            "message": "Yêu cầu token"})
        c.Abort()
        return
    }

    // Kiểm tra định dạng Bearer token
    if len(token) > 7 && strings.HasPrefix(token, "Bearer") {
    } else {
        c.JSON(401, gin.H{
            "status": "Fail",
            "message": "Header Authorization không hợp lệ"})
        c.Abort()
        return
    }


}