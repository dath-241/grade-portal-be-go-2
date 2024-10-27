package routes_admin

import (
	"be-go-2/config"

	"github.com/gin-gonic/gin"
)

func MainRoute(r *gin.Engine) {
	apiAdmin := config.PrefixAdmin()

	adminGroup := r.Group(apiAdmin)

	GradeSheet(adminGroup)
	AuthRoutes(adminGroup)
}
