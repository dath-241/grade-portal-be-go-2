package controller_admin

import (
	"Go2/helper"
	"Go2/models"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// HandleCreateCourse xử lý việc tạo khóa học mới.
func HandleCreateCourse(c *gin.Context) {
	var courseData InteraceCourse

	// Kiểm tra parse data vào có lỗi không
	if err := c.BindJSON(&courseData); err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu không hợp lệ",
		})
		return
	}

	collection := models.CourseModel()

	// Kiểm tra xem khóa học có bị trùng không
	isDuplicate, err := CheckDuplicateCourse(collection, courseData.Ms, courseData.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi kiểm tra dữ liệu",
		})
		return
	}

	// Nếu khóa học đã tồn tại, trả về lỗi
	if isDuplicate {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Khóa học đã tồn tại",
		})
		return
	}

	if courseData.BT+courseData.TN+courseData.BTL+courseData.GK+courseData.CK != 100 {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Sai hệ số, tổng hệ số tối đa là 100",
		})
		return
	}

	// Thêm nếu không bị trùng lặp
	createdBy, _ := c.Get("ID")
	_, err = collection.InsertOne(context.TODO(), bson.M{
		"ms":        courseData.Ms,
		"credit":    courseData.Credit,
		"name":      courseData.Name,
		"desc":      courseData.Desc,
		"createdBy": createdBy,
		"HS":        [5]int{courseData.BT, courseData.TN, courseData.BTL, courseData.GK, courseData.CK},
	})

	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi tạo khóa học",
		})
		return
	}

	// Trả về kết quả thành công
	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Tạo khóa học thành công",
	})
}

// CheckDuplicateCourse kiểm tra xem khóa học có bị trùng không.
func CheckDuplicateCourse(collection *mongo.Collection, ms string, name string) (bool, error) {
	if ms == "" {
		return true, errors.New("lỗi ms không có")
	}

	// Kiểm tra xem có bản ghi nào không
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{
		"ms": ms,
	}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, nil // Không tìm thấy bản ghi
	} else if err != nil {
		return false, err // Có lỗi khác
	}

	return true, nil // Tìm thấy bản ghi trùng
}

// HandleGetCourseByID xử lý việc lấy môn học theo mã môn học.
func HandleGetCourseByID(c *gin.Context) {
	param := c.Param("id")
	courseID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "ID không hợp lệ",
		})
		return
	}

	var course models.InterfaceCourse
	collection := models.CourseModel()

	if err := collection.FindOne(context.TODO(), bson.M{"_id": courseID}).Decode(&course); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"status":    "Fail",
				"message": "Không tìm thấy môn học",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi lấy môn học",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Lấy môn học thành công",
		"data":  course,
	})
}

// HandleGetAllCourses xử lý việc lấy tất cả các khóa học.
func HandleGetAllCourses(c *gin.Context) {
	var allCourses []models.InterfaceCourse
	collection := models.CourseModel()
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi lấy dữ liệu",
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &allCourses); err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi đọc dữ liệu từ cursor",
		})
		return
	}
	semester := helper.SetSemester()
	c.JSON(200, gin.H{
		"status":       "Success",
		"message":    "Lấy ra tất cả khóa học thành công",
		"data": 		allCourses,
		"semester":   semester,
	})
}

// HandleDeleteCourse xử lý việc xóa khóa học.
func HandleDeleteCourse(c *gin.Context) {
	param := c.Param("id")
	courseID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "ID không hợp lệ",
		})
		return
	}
	collection := models.CourseModel()
	if _, err = collection.DeleteOne(context.TODO(), bson.M{"_id": courseID}); err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi xóa khóa học",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    "Success",
		"message": "Xóa khóa học thành công",
	})
}

// HandleUpdateCourse xử lý việc cập nhật thông tin khóa học.
func HandleUpdateCourse(c *gin.Context) {
	param := c.Param("id")
	courseID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "error",
			"message": "ID không hợp lệ",
		})
		return
	}
	var courseData struct {
		Ms        string `json:"ms"`
		Credit    int    `json:"credit"`
		Name      string `json:"name"`
		Desc      string `json:"desc"`
		UpdatedBy any    `json:"updatedBy" bson:"updatedBy"`
	}
	if err = c.BindJSON(&courseData); err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu không hợp lệ",
		})
		return
	}
	collection := models.CourseModel()
	adminID, _ := c.Get("ID")
	courseData.UpdatedBy = adminID
	fmt.Print(courseData)
	if _, err = collection.UpdateOne(context.TODO(), bson.M{"_id": courseID}, bson.M{"$set": courseData}); err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi cập nhật khóa học",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    "Success",
		"message": "Cập nhật khóa học thành công",
	})
}
