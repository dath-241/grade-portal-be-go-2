package routes_admin

import (
	"be-go-2/config"
	controllers_admin "be-go-2/controllers/admin"

	"github.com/gin-gonic/gin"
)

func MainRoute(r *gin.Engine) {
	r.POST("/class/create",controllers_admin.CreateClass)
	apiAdmin := config.PrefixAdmin()
	
	AuthRoutes(r.Group(apiAdmin))
}
