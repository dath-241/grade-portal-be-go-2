package admin_middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ValindateEmail(email string) bool {
	return strings.HasSuffix(email, "@hcmut.edu.vn")
}

func ValidateDataAdmin(c *gin.Context) {


}