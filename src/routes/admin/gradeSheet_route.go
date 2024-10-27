package routes_admin

import (
	controllers_admin "be-go-2/controllers/admin"

	"github.com/gin-gonic/gin"
)

func GradeSheet(r *gin.RouterGroup) {
	r.POST("/gradesheet/create", controllers_admin.CreateGradeSheet)
}
