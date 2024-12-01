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
	createdBy, _ := c.Get("ID")
	//Binding dữ liệu
	var newAccounts []InterfaceAccount
	if err := c.ShouldBindJSON(&newAccounts); err != nil {
		c.JSON(400, gin.H{
			"status": "Fail",
			"message": "Dữ liệu yêu cầu không hợp lệ"})
			return
	}
	//Lấy dữ liệu account để kiểm tra dupplicate
	accountCollection := models.AccountModel()
	var existingAccounts []models.InterfaceAccount
	cursor, err := accountCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{
			"status": "Fail",
			"message": "Lỗi khi lấy tài khoản từ cơ sở dữ liệu"})
		return
	}
	//giải mã dữ liệu từ con trỏ sang kết quả
	if err := cursor.All(context.TODO(), &existingAccounts); err != nil {
		c.JSON(500, gin.H{
			"status": "Fail",
			"message": "Lỗi khi giải mã tài khoản"})
		return
	}
	//Tạo map để check email và mã số có bị trùng hay không
	existingMap := make(map[string]bool)
	var validAccounts []InterfaceAccount
	var invalidAccounts []InterfaceAccount
	for _, account := range existingAccounts {
		existingMap[account.Email] = true
		existingMap[account.Ms] = true
	}
	//Kiểm tra email và role của tài khoản mới đúng định dạng và loại bỏ các tài khoản bị trùng.
	for _, newAccount := range newAccounts {
		if !existingMap[newAccount.Email] && !existingMap[newAccount.Ms] && CheckEmailAndRole(newAccount.Email, newAccount.Role){
			newAccount.CreatedBy = createdBy
			newAccount.ExpiredAt = time.Now().AddDate(5, 0, 0)
			validAccounts = append(validAccounts, newAccount)
		} else {
			invalidAccounts = append(invalidAccounts, newAccount)
		}
	}
	//Thêm tài khoản khả dụng vào cơ sở dữ liệu
	if len(validAccounts) > 0 {
		_, err := accountCollection.InsertMany(context.TODO(), validAccounts); 
		if err != nil {
			c.JSON(500, gin.H{
				"status": "Fail",
				"message": "Lỗi khi tạo tài khoản."})
			return
		}
	}
	c.JSON(200, gin.H{
		"status":            "Success",
		"invalidAccounts": invalidAccounts,
		"validAccounts":   validAccounts,
	})
}

// CheckEmailAndRole kiểm tra đuôi email và role
func CheckEmailAndRole(email string, role string) (bool){
	if(strings.HasSuffix(email, "@hcmut.edu.vn") && (role == "student" || role == "teacher")){
		return true;
	}
	return false
}
// HandleGetAccountByID xử lý việc lấy thông tin tài khoản theo ID.
func HandleGetAccountByID(c *gin.Context) {
	idParam := c.Param("id")
	//Kiểm tra accountID đúng định dạng
	accountID, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "accountID không đúng định dạng.",
		})
		return
	}

	accountCollection := models.AccountModel()
	var account models.InterfaceAccount
	err = accountCollection.FindOne(context.TODO(), bson.M{"_id": accountID}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"status":    "Fail",
				"message": "Không tìm thấy tài khoản."})
			return
		}
		c.JSON(500, gin.H{
				"status":    "Fail",
				"message": "Lỗi khi lấy tài khoản từ cơ sở dữ liệu."})
		return
	}

	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Tìm tài khoản thành công.",
		"data": account,
	})
}

// HandleGetTeacherAccounts xử lý việc lấy thông tin tài khoản giáo viên.
func HandleGetTeacherAccounts(c *gin.Context) {
	accountCollection := models.AccountModel()
	query := c.Query("ms")
	// Lấy tất cả giáo viên
	if query == "" {
		var teachers []models.InterfaceAccount
		cursor, err := accountCollection.Find(context.TODO(), bson.M{"role": "teacher"})
		if err != nil {
			//Không có giảng viên trong cơ sở dữ liệu
			if err == mongo.ErrNoDocuments {
				c.JSON(404, gin.H{
					"status":    "Fail",
					"message": "Không có tài khoản nào trong cơ sở dữ liệu.",
				})
				return
			}
			// Lấy dữ liệu bị lỗi
			c.JSON(500, gin.H{
				"status":    "Fail",
				"message": "Lỗi khi lấy dữ liệu từ cơ sở dữ liệu.",})
			return
		}
		// giải mã từ con trỏ sang kết quả
		if err := cursor.All(context.TODO(), &teachers); err != nil {
			c.JSON(500, gin.H{
					"status":    "Fail",
					"message": "Lỗi khi giải mã tài khoản.",})
			return
		}
		// trả về kết quả tìm được
		c.JSON(200, gin.H{
			"status":    "Success",
			"message":      "Tìm tài khoản thành công",
			"data": teachers,
		})
		return
	} else {
	// Lấy giáo viên theo mã số
		var teacher models.InterfaceAccount
		err := accountCollection.FindOne(context.TODO(), bson.M{"ms": query}).Decode(&teacher)
		if err != nil {
			//Không tồn tại giáo viên với mã số đã nhập
			if err == mongo.ErrNoDocuments {
				c.JSON(404, gin.H{
					"status":    "Fail",
					"message": "Không tìm thấy tài khoản"})
				return
			}
			// Lấy dữ liệu bị lỗi
			c.JSON(500, gin.H{
				"status":    "Fail",
				"message": "Lỗi khi lấy tài khoản từ cơ sở dữ liệu"})
			return
		}
		// trả về kết quả tìm được
		c.JSON(http.StatusOK, gin.H{
			"status":    "Success",
			"message":      "Tìm tài khoản thành công",
			"data": teacher,
		})
		return;
	}
}
// HandleGetStudentAccounts xử lý việc lấy thông tin tài khoản sinh viên.
func HandleGetStudentAccounts(c *gin.Context) {
	accountCollection := models.AccountModel()
	query := c.Query("ms")
	// Lấy tất cả sinh viên
	if query == "" {
		var students []models.InterfaceAccount
		cursor, err := accountCollection.Find(context.TODO(), bson.M{"role": "student"})
		if err != nil {
			//Không có sinh vien trong cơ sở dữ liệu
			if err == mongo.ErrNoDocuments {
				c.JSON(404, gin.H{
					"status":    "Fail",
					"message": "Không có tài khoản nào trong cơ sở dữ liệu"})
				return
			}
			//Lấy dữ liệu bị lỗi
			c.JSON(500, gin.H{
				"status":    "Fail",
				"message": "Lỗi khi lấy tài khoản"})
			return
		}
		//Giải mã từ con trỏ sang kết quả
		if err := cursor.All(context.TODO(), &students); err != nil {
			c.JSON(500, gin.H{
				"status":    "Fail",
				"mesage": "Lỗi khi giải mã tài khoản"})
			return
		}
		//Trả kết quả về
		c.JSON(200, gin.H{
			"status":    "Success",
			"message":      "Tìm thấy tài khoản thành công",
			"data": students,
		})
	} else {
	// Lấy sinh viên theo mã số
		var student models.InterfaceAccount
		err := accountCollection.FindOne(context.TODO(), bson.M{"ms": query}).Decode(&student)
		if err != nil {

			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{
					"status":    "Fail",
					"message": "Không tìm thấy tài khoản"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":    "Fail",
				"message": "Lỗi khi lấy tài khoản"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":    "Success",
            "message":      "Tìm thấy tài khoản thành công",
			"data": student,
		})
	}
}

// HandleDeleteAccount xử lý việc xóa tài khoản.
func HandleDeleteAccount(c *gin.Context) {
	idParam := c.Param("id")
	//Kiểm tra định dạng ID
	accountID, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "ID không đúng định dạng",
		})
		return
	}
	//Xóa tài khoản
	accountCollection := models.AccountModel()
	_ , err = accountCollection.DeleteOne(context.TODO(), bson.M{"_id": accountID})
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi xóa tài khoản",
		})
		return
	}
	//Trả kết quả
	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Xóa tài khoản thành công",
	})
}

// HandleUpdateAccount xử lý việc cập nhật thông tin tài khoản.
func HandleUpdateAccount(c *gin.Context) {
	idParam := c.Param("id")
	createdBy, _ := c.Get("ID")
	accountID, err := bson.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "ID không đúng định dạng",
		})
		return
	}
	var updatedAccount InterfaceAccount
	if err := c.ShouldBindJSON(&updatedAccount); err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu yêu cầu không hợp lệ"})
		return
	}
	updatedAccount.CreatedBy = createdBy
	accountCollection := models.AccountModel();
	_, err = accountCollection.UpdateOne(context.TODO(), bson.M{"_id": accountID}, bson.M{"$set": updatedAccount}); 
	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi cập nhật tài khoản vào cơ sở dữ liệu",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Cập nhật tài khoản thành công",
	})
}
