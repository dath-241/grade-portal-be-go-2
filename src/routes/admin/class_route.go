package routes_admin

import(
	controllers_admin "be-go-2/controllers/admin"
	"github.com/gin-gonic/gin"
)
func ClassRoutes(r *gin.RouterGroup){
	r.POST("/class/create",controllers_admin.CreateClass)
}