package controller_admin

import (
	"Go2/models"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// HandleCreateAccount xử lý việc tạo tài khoản mới.
func HandleCreateAccount(c *gin.Context) {
	createdBy, _ := c.Get("ID") // Get creator's ID

	var newAccounts []InterfaceAccount

	// Check if the body of the request is valid
	if err := c.ShouldBindJSON(&newAccounts); err != nil {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Invalid data",
		})
		return
	}

	// Check dupplicate
	accountCol := models.AccountModel()
	var existedAccounts []models.InterfaceAccount

	cusor, err := accountCol.Find(context.TODO(), bson.M{})

	if err != nil { // Error when retrieve database
		c.JSON(500, gin.H{
			"status":  "Fail",
			"message": "Error when retrieve on database",
		})
		return
	}

	if err := cusor.All(context.TODO(), &existedAccounts); err != nil {
		c.JSON(500, gin.H{
			"status":  "Fail",
			"message": "Error decoding account",
		})
		return
	}

	// Arrays to store dupplicating email and id
	var emailSet, idSet []string
	for _, account := range existedAccounts {
		emailSet = append(emailSet, account.Email)
		idSet = append(idSet, account.Ms)
	}
	// Classify valid and invalid accounts
	var validAccounts, invalidAccounts []InterfaceAccount
	for _, account := range newAccounts {
		if contains(emailSet, account.Email) || contains(idSet, account.Ms) {
			invalidAccounts = append(invalidAccounts, account)
		} else {
			validAccounts = append(validAccounts, account)
		}
	}

	// Add valid accounts to database
	if len(validAccounts) > 0 {
		_, err := accountCol.InsertMany(context.TODO(), validAccounts)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "Fail",
				"message": "Error when creating account",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"status":          "success",
		"invalidAccounts": invalidAccounts,
		"validAccount":    validAccounts,
	})
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// CheckEmailAndRole kiểm tra đuôi email và role
func CheckEmailAndRole(email string, role string) bool {
	return true
}

// HandleGetAccountByID xử lý việc lấy thông tin tài khoản theo ID.
func HandleGetAccountByID(c *gin.Context) {

}

// HandleGetTeacherAccounts xử lý việc lấy thông tin tài khoản giáo viên.
func HandleGetTeacherAccounts(c *gin.Context) {

}

// HandleGetStudentAccounts xử lý việc lấy thông tin tài khoản sinh viên.
func HandleGetStudentAccounts(c *gin.Context) {

}

// HandleDeleteAccount xử lý việc xóa tài khoản.
func HandleDeleteAccount(c *gin.Context) {

}

// HandleUpdateAccount xử lý việc cập nhật thông tin tài khoản.
func HandleUpdateAccount(c *gin.Context) {

}
