package routes_admin

import (
	controllers_admin "be-go-2/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	adminRoutes := router.Group("/admin")
	{
		adminRoutes.POST("/teacher/create", controllers_admin.CreateTeacher)
	}
}
