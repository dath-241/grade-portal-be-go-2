package routes_admin

import (
	"be-go-2/controllers/admin"
	"be-go-2/middlewares/admin"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/login", controllers_admin.LoginController)
	r.POST("/logout", admin_middlewares.RequireAuth, controllers_admin.LogoutController)
	r.POST("/create", admin_middlewares.RequireAuth, admin_middlewares.ValidateDataAdmin, controllers_admin.CreateAdminController)
}
