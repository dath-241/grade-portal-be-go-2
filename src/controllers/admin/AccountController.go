package controller_admin

import (
	"Go2/models"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// HandleCreateAccount xử lý việc tạo tài khoản mới.
func HandleCreateAccount(c *gin.Context) {
	createdBy, _ := c.Get("ID") // Get creator's ID
	var newAccounts []InterfaceAccount

	// Check if the body of the request is valid
	if err := c.ShouldBindJSON(&newAccounts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Fail",
			"message": "Error when retrieve on database",
		})
		return
	}

	if err := cusor.All(context.TODO(), &existedAccounts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Fail",
				"message": "Error when creating account",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
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

// HandleGetAccountByID xử lý việc lấy thông tin tài khoản theo ID.
func HandleGetAccountByID(c *gin.Context) {
	idParam := c.Param("id")

	accountId, err := bson.ObjectIDFromHex(idParam)
	if err != nil { // Check if idParam is valid or not
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Fail",
			"message": "Invalid ID",
		})
		return
	}

	accountCol := models.AccountModel()
	var account models.InterfaceAccount

	// Find account in database by accountId
	err = accountCol.FindOne(context.TODO(), bson.M{"_id": accountId}).Decode(&account)
	if err != nil { // If there is an error when finding account
		if err == mongo.ErrNoDocuments { // Can not find account with accountId
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "Fail",
				"message": "Can not find account",
			})
			return
		}
		// Another error when using database
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Fail",
			"message": "Internal server error",
		})
		return
	}

	// No error, find account successfully, return that account data
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Find account successfully",
		"data":    account,
	})
}

// HandleGetTeacherAccounts xử lý việc lấy thông tin tất cả tài khoản giáo viên hoặc theo id
func HandleGetTeacherAccounts(c *gin.Context) {
	accountCol := models.AccountModel()

	query := c.Query("ms")

	if query == "" { // Get all teacher accounts
		var teachers []models.InterfaceAccount
		// Retriev database
		cursor, err := accountCol.Find(context.TODO(), bson.M{"role": "teacher"})

		if err != nil { // If there is an error when finding account
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{ // There are no teacher accounts database
					"status":  "Fail",
					"message": "Can not find account",
				})
				return
			}
			// Another error when using database
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Fail",
				"message": "Internal server error",
			})
			return
		}

		// Decoding cursor
		if err := cursor.All(context.TODO(), &teachers); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Fail",
				"message": "Error decoding account",
			})
			return
		}

		// There is no error
		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Find all teacher accounts successfully",
			"data":    teachers,
		})

	} else { // Get teacher account by `ms`
		var teacher models.InterfaceAccount
		err := accountCol.FindOne(context.TODO(), bson.M{"role": "teacher", "ms": query}).Decode(&teacher)

		if err != nil { // If there is an error when finding account
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{ // There are no teacher account with ms database
					"status":  "Fail",
					"message": "Can not find teacher account with ms",
				})
				return
			}
			// Another error when using database
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Fail",
				"message": "Internal server error",
			})
			return
		}

		// Maybe there is no error
		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Find teacher account by ID successfully",
			"data":    teacher,
		})
	}
}

// HandleGetStudentAccounts xử lý việc lấy thông tin tất cả tài khoản sinh viên hoặc theo id
func HandleGetStudentAccounts(c *gin.Context) {
	accountCol := models.AccountModel()

	query := c.Query("ms")

	if query == "" {
		var students []models.InterfaceAccount
		cursor, err := accountCol.Find(context.TODO(), bson.M{"role": "student"})

		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{
					"status":  "Fail",
					"message": "Can not find student account with ms",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Fail",
				"message": "Internal server error",
			})
			return
		}

		if err := cursor.All(context.TODO(), &students); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Fail",
				"message": "Error decoding account",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Find student account by ID successfully",
			"data":    students,
		})
	} else {
		var student models.InterfaceAccount
		err := accountCol.FindOne(context.TODO(), bson.M{"role": "student", "ms": query}).Decode(&student)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{
					"status":  "Fail",
					"message": "Can not find student account with ms",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Fail",
				"message": "Internal server error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Find student account by ID successfully",
			"data":    student,
		})
	}
}

// Delete account
func HandleDeleteAccount(c *gin.Context) {
	idParam := c.Param("id")
	accountID, err := bson.ObjectIDFromHex(idParam)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Fail",
			"message": "Can not find ID",
		})
		return
	}

	// Delete
	collection := models.AccountModel()
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": accountID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Fail",
			"message": "Can not delete this account",
		})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Fail",
			"message": "Can not find account with id",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Delete account successfully",
	})
}

func HandleUpdateAccount(c *gin.Context) {
	idParam := c.Param("id")
	createdBy, _ := c.Get("ID")
	accountId, err := bson.ObjectIDFromHex(idParam)

	if err != nil { // id is invalid
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Fail",
			"message": "Invalid id",
		})
		return
	}

	var updatedAccont InterfaceAccount
	if err := c.ShouldBindJSON(&updatedAccont); err != nil { // Body request is invalid
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Fail",
			"message": "Invalid data",
		})
		return
	}

	updatedAccont.CreatedBy = createdBy // CreatedBy -> UpdatedBy
	acccountCol := models.AccountModel()

	// Update that account
	filter := bson.M{"_id": accountId}
	updateData := bson.M{"$set": updatedAccont}
	result, err := acccountCol.UpdateOne(context.TODO(), filter, updateData)

	if err != nil { // Error when update in database
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Fail",
			"message": "Error when updating account",
		})
		return
	}

	if result.MatchedCount == 0 { // Can not find account with that id
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Fail",
			"message": "Can not find account with id",
		})
		return
	}

	// May be there is no error left
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Update account successfully",
	})
}
