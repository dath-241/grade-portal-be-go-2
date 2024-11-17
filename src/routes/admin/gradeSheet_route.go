package routes_admin

import (
	controllers_admin "be-go-2/controllers/admin"

	"github.com/gin-gonic/gin"
)

// GradeSheet sets up the routes for handling grade sheets
func GradeSheet(r *gin.RouterGroup) {
	r.POST("/teacher/course/:courseID/class/:classID/upload", controllers_admin.CreateGradeSheet)
	r.GET("/teacher/course/:courseID/class/:classID/grades", controllers_admin.GetGradeSheet)
}
