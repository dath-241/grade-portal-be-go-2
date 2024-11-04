package routes_admin

import (
	controllers_admin "be-go-2/controllers/admin"
	"github.com/gin-gonic/gin"
)

func TeacherRoutes(r *gin.RouterGroup) {
	r.POST("/teacher", controllers_admin.CreateTeacher)
}
