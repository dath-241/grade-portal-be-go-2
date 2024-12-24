package middlewares_admin

import (
	"Go2/helper"
	"fmt"
	"net/http"
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
            "message": "Token không hợp lệ"})
        c.Abort()
        return
    }
    Claims, _ := helper.ParseJWT(token)
	if Claims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":    "Fail",
			"message": "Người dùng chưa đăng nhập",
		})
		c.Abort()
		return
	}
    fmt.Println(Claims.ID)
	c.Set("ID", Claims.ID)
	c.Next()

}