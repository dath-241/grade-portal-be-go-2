package controller_client

import (
	"Go2/models"
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// HandleResult xử lý yêu cầu lấy kết quả điểm của người dùng
func HandleResult(c *gin.Context) {
	data, _ := c.Get("user")
	param := c.Param("id")
	classID, _ := bson.ObjectIDFromHex(param)
	user := data.(models.InterfaceAccount)
	var result models.InterfaceResult
	collection := models.ResultScoreModel()

	// Tìm kiếm kết quả điểm theo class_id
	if err := collection.FindOne(context.TODO(), bson.M{"class_id": classID}).Decode(&result); err != nil {
		c.JSON(401, gin.H{
			"status":    "Fail",
			"message": "Bạn không có quyền truy cập bảng điểm này",
		})
		return
	}

	// Kiểm tra vai trò của người dùng
	if user.Role == "teacher" {
		c.JSON(200, gin.H{
			"status":   "Success",
			"data": result,
		})
		return
	} else if user.Role == "student" {
		for _, item := range result.SCORE {
			if item.MSSV == user.Ms {
				c.JSON(200, gin.H{
					"status":  "Success",
					"data": item,
				})
				return
			}
		}
	}

	c.JSON(401, gin.H{
		"status":    "Fail",
		"message": "Bạn không có quyền vào đây",
	})
}

// HandleCreateResult xử lý yêu cầu tạo mới kết quả điểm
func HandleCreateResult(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	var dataResult InterfaceResultScoreController

	// Lấy dữ liệu từ front end
	c.BindJSON(&dataResult)
	classID, err := bson.ObjectIDFromHex(dataResult.ClassID)
	if err != nil {
		c.JSON(404, gin.H{
			"status":    "Fail",
			"message": "Lớp chưa có giáo viên",
		})
		return
	}

	var classDetail models.InterfaceClass
	collectionClass := models.ClassModel()

	// Tìm kiếm chi tiết lớp học
	if err = collectionClass.FindOne(context.TODO(), bson.M{"_id": classID}).Decode(&classDetail); err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Không tìm thấy lớp học đó",
		})
		return
	}

	collection := models.ResultScoreModel()
	var result models.InterfaceResult

	// Kiểm tra xem đã có bản ghi result trước đó chưa
	err = collection.FindOne(context.TODO(), bson.M{"class_id": classID}).Decode(&result)
	if err == nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Bảng ghi của lớp học này đã được lưu trong database trước đó",
		})
		return
	}

	// Tạo mới bản ghi result
	collection.InsertOne(context.TODO(), bson.M{
		"semester":  classDetail.Semester,
		"course_id": classDetail.CourseId,
		"score":     dataResult.SCORE,
		"class_id":  classID,
		"expiredAt": time.Now().AddDate(0, 6, 0),
		"createdBy": user.ID,
		"updatedBy": user.ID,
	})

	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Cập nhật bảng điểm thành công",
	})
}

// HandlePatchResult xử lý yêu cầu cập nhật kết quả điểm
func HandlePatchResult(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	var dataResult InterfaceResultScoreController
	c.BindJSON(&dataResult)
	classID, _ := bson.ObjectIDFromHex(dataResult.ClassID)
	collection := models.ResultScoreModel()

	// Cập nhật kết quả điểm
	result, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"class_id": classID},
		bson.M{"$set": bson.M{
			"score":     dataResult.SCORE,
			"updatedBy": user.ID,
		}},
	)
	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi hệ thống",
		})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Thay đổi không hợp lệ",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Thay đổi thành công",
	})
}

// HandleCourseResult xử lý yêu cầu lấy kết quả điểm của khóa học
func HandleCourseResult(c *gin.Context) {
	data, _ := c.Get("user")
	account := data.(models.InterfaceAccount)
	param := c.Param("ms")
	params := strings.Split(param, "-")
	var course models.InterfaceCourse
	collectionCourse := models.CourseModel()

	// Tìm kiếm khóa học theo mã số
	if err := collectionCourse.FindOne(context.TODO(), bson.M{"ms": params[0]}).Decode(&course); err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "MS course sai",
		})
		return
	}

	var result models.InterfaceResult
	collectionResult := models.ResultScoreModel()

	// Tìm kiếm kết quả điểm theo course_id và học kỳ
	err := collectionResult.FindOne(context.TODO(), bson.M{"course_id": course.ID, "semester": params[1]}).Decode(&result);
	if err != nil {
		if err == mongo.ErrNoDocuments{
			c.JSON(200, gin.H{
				"status":    "Success",
				"message": "Không có bảng điểm",
			})
			return
		}
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "ID course sai",
		})
		return
	}

	// Kiểm tra điểm của người dùng
	for _, item := range result.SCORE {
		if item.MSSV == account.Ms {
			c.JSON(200, gin.H{
				"status":    "Success",
				"message": "Lấy điểm thành công",
				"data":   item.Data,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"status":    "Fail",
		"message": " Không tìm thấy điểm",
	})
}

// HandleAllResults xử lý yêu cầu lấy tất cả kết quả điểm của người dùng
func HandleAllResults(c *gin.Context) {
	data, _ := c.Get("user")
	account := data.(models.InterfaceAccount)
	collection := models.ResultScoreModel()
	var result []models.InterfaceResult

	// Tìm kiếm tất cả kết quả điểm của người dùng
	cursor, err := collection.Find(context.TODO(), bson.M{"score.mssv": account.Ms})
	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi server",
		})
		return
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &result); err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi server",
		})
		return
	}

	type score struct {
		Ms   string                `json:"ms"`
		Data models.InterfaceScore `json:"data"`
	}
	var scores []score

	// Lấy điểm của người dùng từ kết quả
	for _, item := range result {
		for _, sco := range item.SCORE {
			if sco.MSSV == account.Ms {
				collectionCourse := models.CourseModel()
				var course models.InterfaceCourse
				if err := collectionCourse.FindOne(context.TODO(), bson.M{"_id": item.CourseID}).Decode(&course); err != nil {
					c.JSON(400, gin.H{
						"status":    "Fail",
						"message": "MS course sai",
					})
					return
				}
				scores = append(scores, score{course.MS + "-" + item.Semester, sco.Data})
			}
		}
	}

	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Lấy điểm thành công",
		"data":  scores,
	})
}
