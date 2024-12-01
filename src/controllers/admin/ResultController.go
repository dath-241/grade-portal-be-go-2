package controller_admin

import (
	"Go2/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// HandleCreateResult dùng để tạo hoặc cập nhật bảng điểm.
func HandleCreateResult(c *gin.Context) {
	var data InterfaceResultScoreController
	// Binding dữ liệu
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu không hợp lệ",
		})
		return
	}
	// Kiểm tra ClassID có đúng định dạng không
	classID, err := bson.ObjectIDFromHex(data.ClassID)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Mã lớp học không đúng định dạng.",
		})
		return
	}
	// Kiểm tra lớp học có tồn tại không
	var classDetail models.InterfaceClass
	collectionClass := models.ClassModel()
	if err = collectionClass.FindOne(context.TODO(), bson.M{"_id": classID}).Decode(&classDetail); err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Không tìm thấy lớp học đó",
		})
		return
	}

	createdBy, _ := c.Get("ID")

	//Kiểm tra nếu bảng điểm đã tồn tại thì chỉ cập nhật
	collection := models.ResultScoreModel()
	var result models.InterfaceResult
	err = collection.FindOne(context.TODO(), bson.M{"class_id": classID}).Decode(&result)
	if err == nil {
		filter := bson.M{"class_id": classID};
		update := bson.M{"$set": bson.M{"score": data.SCORE, "updatedBy": createdBy, "expiredAt": time.Now().AddDate(0, 6, 0)}};
		_ , err := collection.UpdateOne(context.TODO(), filter, update);
		if err != nil {
			c.JSON(500, gin.H{
				"status":    "Fail",
				"message": "Cập nhật bảng điểm thât bại!.",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":    "Success",
			"message": "Cập nhật bảng điểm thành công",
		})
		return
	}
	//Nếu không tồn tại thì tạo mới
	collection.InsertOne(context.TODO(), bson.M{
		"semester":  classDetail.Semester,
		"course_id": classDetail.CourseId,
		"score":     data.SCORE,
		"class_id":  classID,
		"expiredAt": time.Now().AddDate(0, 6, 0),
		"createdBy": createdBy,
		"updatedBy": createdBy,
	})
	c.JSON(http.StatusOK, gin.H{
		"status":    "Success",
		"message": "Thêm bảng điểm thành công.",
	})
}

// HandleGetResult xử lý việc lấy bảng điểm theo ID lớp học.
func HandleGetResult(c *gin.Context) {
	//Lấy và kiểm tra ClassID có đúng định dạng hay không
	param := c.Param("id")
	classID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":    "Fail",
			"message": "Mã lớp học không đúng định dạng.",
		})
		return
	}
	//Kiểm tra bảng điểm có tồn tại hay không
	collection := models.ResultScoreModel()
	var data models.InterfaceResult
	if err = collection.FindOne(context.TODO(), bson.M{"class_id": classID}).Decode(&data); err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Không tìm thấy bảng điểm.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    "Success",
		"message": "Lấy bảng điểm thành công.",
		"data":   data,
	})
}
