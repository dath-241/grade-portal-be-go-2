package controller_client

import (
	"Go2/models"
	"context"
	"fmt"
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
			fmt.Println(item.MSSV)
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
		"message": "Tạo bảng điểm thành công",
	})
}

// HandlePatchResult xử lý yêu cầu cập nhật kết quả điểm
// HandlePatchResult xử lý việc cập nhật kết quả điểm theo ID
func HandlePatchResult(c *gin.Context) {
	param := c.Param("id")
	resultID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}

	var data models.InterfaceScore
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}

	// Tạo một bản đồ để lưu trữ các trường cần cập nhật
	updateFields := bson.M{}

	// Chỉ thêm các trường có trong yêu cầu
	if len(data.BT) > 0 {
		updateFields["bt"] = data.BT
	}
	if len(data.TN) > 0 {
		updateFields["tn"] = data.TN
	}
	if len(data.BTL) > 0 {
		updateFields["btl"] = data.BTL
	}
	if data.GK != 0 {
		updateFields["gk"] = data.GK
	}
	if data.CK != 0 {
		updateFields["ck"] = data.CK
	}

	// Nếu không có trường nào để cập nhật, trả về lỗi
	if len(updateFields) == 0 {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Không có dữ liệu nào để cập nhật",
		})
		return
	}

	collection := models.ResultScoreModel()
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": resultID}, bson.M{"$set": updateFields})
	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi cập nhật bảng điểm",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Cập nhật bảng điểm thành công",
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
	err := collectionCourse.FindOne(context.TODO(), bson.M{"ms": params[0]}).Decode(&course)
	if ; err != nil {
		if err == mongo.ErrNoDocuments{
			c.JSON(404, gin.H{
        "status":    "Fail",
        "message": "Không tìm thấy khóa học",
      })
      return
		}
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi truy vấn dữ liệu",
		})
		return
	}

	var result models.InterfaceResult
	collectionResult := models.ResultScoreModel()

	// Tìm kiếm kết quả điểm theo course_id và học kỳ
	err = collectionResult.FindOne(context.TODO(), bson.M{"course_id": course.ID, "semester": params[1]}).Decode(&result);
	if err != nil {
		if err == mongo.ErrNoDocuments{
			c.JSON(404, gin.H{
				"status":    "Fail",
				"message": "Không có bảng điểm",
			})
			return
		}
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi truy vấn dữ liệu",
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
			"message": "Lỗi khi truy vấn dữ liệu",
		})
		return
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &result); err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi giải mã dữ liệu",
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
