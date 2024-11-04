package routes_admin

import (
	"be-go-2/config"

	"github.com/gin-gonic/gin"
)

// Main route include all routes use for the website?
func MainRoute(r *gin.Engine) {
	apiAdmin := config.PrefixAdmin()
	AuthRoutes(r.Group(apiAdmin))
	CourseRoutes(r)
}
