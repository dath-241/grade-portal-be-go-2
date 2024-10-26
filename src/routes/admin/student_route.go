package routes_admin

import (
	controllers_admin "be-go-2/controllers/admin"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.RouterGroup) {
	r.POST("/create", controllers_admin.CreateStudent)
}
