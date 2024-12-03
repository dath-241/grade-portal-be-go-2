package controller_admin

import (
	"Go2/models"
	"context"
	"strings"
	"time"

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

	// Remove dupplicate object with ID and Email in newAccounts
	newAccounts = removeDuplicates(newAccounts)

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
		// Check @hcmut.edu.vn, role and dupplicated
		if contains(emailSet, account.Email) || contains(idSet, account.Ms) || !CheckEmailAndRole(account.Email, account.Role) {
			invalidAccounts = append(invalidAccounts, account)
		} else {
			// Add field CreatedBy and ExpiredAt for valid account
			account.CreatedBy = createdBy
			account.ExpiredAt = time.Now().AddDate(5, 0, 0)
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

func removeDuplicates(accounts []InterfaceAccount) []InterfaceAccount {
	seen := make(map[string]map[string]bool) // Map store email and id has met
	var result []InterfaceAccount

	for _, account := range accounts {
		if seen[account.Email][account.Ms] {
			// Dupplicate, pass
			continue
		}

		// No dupplicate, add accounts to result
		if seen[account.Email] == nil {
			seen[account.Email] = make(map[string]bool)
		}
		seen[account.Email][account.Ms] = true
		result = append(result, account)
	}

	return result
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
	if strings.HasSuffix(email, "@hcmut.edu.vn") && (role == "student" || role == "teacher") {
		return true
	}
	return false
}
