package routes_admin

import (
	controllers_admin "be-go-2/controllers/admin"

	"github.com/gin-gonic/gin"
)

// CourseRoutes include all the routes for courses
func CourseRoutes(app *gin.Engine) {
	api := app.Group("/course")
	{
		api.POST("/create", controllers_admin.CreateCourse)
	}
}
