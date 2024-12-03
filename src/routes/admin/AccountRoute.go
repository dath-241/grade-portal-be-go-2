package routes_admin

import (
	"github.com/gin-gonic/gin"
)

// AccountRoute thiết lập các route cho tài khoản
func AccountRoute(r *gin.RouterGroup) {
	// Create new admin account
	r.POST("/create", controller_admin.HandleCreateAccount)

	// Find admin account by ID
	r.GET("/:id", controller_admin.HandleGetAccountByID)

	// Get all teacher accounts
	r.GET("/teacher", controller_admin.HandleGetTeacherAccounts)

	// Get all student accounts
	r.GET("/student", controller_admin.HandleGetStudentAccounts)

	// Delete account by ID
	// r.DELETE("/delete/:id", controller_admin.HandleDeleteAccount)

	// Update account by ID
	// r.PATCH("/change/:id", controller_admin.HandleUpdateAccount)
}
