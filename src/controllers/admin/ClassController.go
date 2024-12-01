package controller_admin

import (
	"Go2/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// HandleCreateClass xử lý việc tạo lớp học mới.
func HandleCreateClass(c *gin.Context) {
	var data InterfaceClass
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu không hợp lệ",
		})
		return
	}
	courseID, err := bson.ObjectIDFromHex(data.CourseId)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Course ID không hợp lệ",
		})
		return
	}
	teacherID, err := bson.ObjectIDFromHex(data.TeacherId)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Teacher ID không hợp lệ",
		})
		return
	}
	
	collection := models.ClassModel()

	// Kiểm tra xem lớp học có bị trùng không
	isDuplicate, err := CheckDuplicateClass(collection, data.Semester, courseID, data.Name, teacherID)
	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi kiểm tra dữ liệu",
		})
		return
	}

	// Nếu lớp học đã tồn tại, trả về lỗi
	if isDuplicate {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Lớp học đã tồn tại",
		})
		return
	}

	// Thêm nếu không bị trùng lặp
	createdBy, _ := c.Get("ID")

	_, err = collection.InsertOne(context.TODO(), bson.M{
		"semester":       data.Semester,
		"name":           data.Name,
		"course_id":      courseID,
		"listStudent_ms": data.ListStudentMs,
		"teacher_id":     teacherID,
		"createdBy":      createdBy,
		"updatedBy":      createdBy,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi tạo lớp học",
		})
		return
	}

	// Trả về kết quả thành công
	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Tạo lớp học thành công",
	})
}

// CheckDuplicateClass kiểm tra xem lớp học có bị trùng không.
func CheckDuplicateClass(collection *mongo.Collection, semester string, courseID bson.ObjectID, name string, teacherID bson.ObjectID) (bool, error) {

	// Sử dụng FindOne để kiểm tra xem có bản ghi nào không
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{
		"semester":  semester,
		"course_id": courseID,
		"name":      name,
		"teacher_id": teacherID,
	}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, nil // Không tìm thấy bản ghi
	} else if err != nil {
		return false, err // Có lỗi khác
	}

	return true, nil // Tìm thấy bản ghi trùng
}

// CheckStudentOrTeacher hỗ trợ kiểm tra student hay teacher
func CheckStudentOrTeacher(c *gin.Context, id string, mssv *string) bool { // Student -> true, Teacher -> false
	collection := models.AccountModel()
	// Chuyển đổi id từ string sang ObjectID
	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":    "Fail",
			"message": "Lỗi định dạng dữ liệu"})
		return false // Xử lý lỗi và trả về false
	}

	cursor, err := collection.Find(context.TODO(), bson.M{
		"_id":  objectID,
		"role": "student",
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":    "Fail",
			"message": "Lỗi hệ thống"})
		return false // Xử lý lỗi và trả về false
	}
	defer cursor.Close(context.TODO()) // Đảm bảo đóng cursor sau khi sử dụng

	// Kiểm tra xem có tài liệu nào không
	if cursor.Next(context.TODO()) {
		// Nếu có tài liệu, trả về true
		var user models.InterfaceAccount
		cursor.Decode(&user)
		*mssv = user.Ms
		return true
	} else if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":    "Fail",
			"message": "Lỗi hệ thống."})
		return false
	}

	// Nếu không có tài liệu nào, trả về false
	return false
}

// HandleGetAllClassesByAccountID xử lý việc lấy tất cả lớp học theo account_id
func HandleGetAllClassesByAccountID(c *gin.Context) {
	accountID := c.Param("id")

	collection := models.ClassModel()
	var mssv string


	// Tìm tất cả lớp học mà giáo viên hoặc sinh viên với account_id tham gia
	isStudent := CheckStudentOrTeacher(c, accountID, &mssv)
	var filter bson.M
	if isStudent {
		filter = bson.M{"listStudent_ms": bson.M{"$in": []string{mssv}}} // Nếu là student
	} else {
		id, _ := bson.ObjectIDFromHex(accountID)
		filter = bson.M{"teacher_id": id} // Nếu là teacher
	}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Lỗi truy xuất dữ liệu"})
	}
	var classes []models.InterfaceClass
	// Đọc dữ liệu từ cursor
	if err := cursor.All(context.TODO(), &classes); err != nil {
		c.JSON(500, gin.H{
				"status":    "Fail",
				"message": "Lỗi khi giải mã tài khoản.",})
		return
	}

	// Trả về danh sách lớp học
	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Lấy lớp học thành công",
		"data": classes,
	})
}

// HandleGetClassByID xử lý việc lấy chi tiết lớp học theo ID
func HandleGetClassByID(c *gin.Context) {
	param := c.Param("id")
	classID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "ID không hợp lệ",
		})
		return
	}

	var class models.InterfaceClass
	collection := models.ClassModel()

	if err := collection.FindOne(context.TODO(), bson.M{"_id": classID}).Decode(&class); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"status":  "Fail",
				"message": "Không tìm thấy lớp học",
			})
		} else {
			c.JSON(500, gin.H{
				"status":    "Fail",
				"message": "Lỗi khi lấy dữ liệu từ cơ sở dữ liệu",
			})
			return
		}
		return
	}

	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Lấy lớp học thành công",
		"data":   class,
	})
}

// HandleGetClassesByCourseID xử lý việc lấy tất cả lớp học theo mã môn học
func HandleGetClassesByCourseID(c *gin.Context) {
	param := c.Param("id")
	courseID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "ID không hợp lệ",
		})
		return
	}
	collection := models.ClassModel()
	cursor, err := collection.Find(context.TODO(), bson.M{"course_id": courseID})
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Không tìm thấy lớp học",
		})
		return
	}
	var classes []models.InterfaceClass
	// Đọc dữ liệu từ cursor
	if err := cursor.All(context.TODO(), &classes); err != nil {
		c.JSON(500, gin.H{
				"status":    "Fail",
				"message": "Lỗi khi giải mã tài khoản.",})
		return
	}
	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Lấy lớp học thành công",
		"data": classes,
	})
}

// HandleAddStudentsToClass xử lý việc thêm học sinh vào lớp học
func HandleAddStudentsToClass(c *gin.Context) {
	var request InterfaceAddStudentClassController

	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	collection := models.ClassModel()
	classID, err := bson.ObjectIDFromHex(request.ClassId)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "ID lớp học không hợp lệ",
		})
		return
	}
	filter := bson.M{"_id": classID}
	update := bson.M{
		"$addToSet": bson.M{
			"listStudent_ms": bson.M{
				"$each": request.ListStudentMs,
			},
		},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi thêm học sinh vào lớp học",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Thêm học sinh vào lớp học thành công",
	})
}

// HandleUpdateClass xử lý việc cập nhật thông tin lớp học
func HandleUpdateClass(c *gin.Context) {
	param := c.Param("id")
	classID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "ID không hợp lệ",
		})
		return
	}
	var data InterfaceChangeClassController
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu không hợp lệ",
		})
		return
	}
	if data.Name == "" {
		c.JSON(400, gin.H{
      "status":    "Fail",
      "message": "Tên lớp không được để trống",
    })
    return
	}
	if data.Semester == "" {
		c.JSON(400, gin.H{
      "status":    "Fail",
      "message": "Semester không được để trống",
    })
    return
	}
	var courseID bson.ObjectID
	courseIDStr, _ := data.CourseId.(string)
	if courseIDStr != "" {
		courseID, err = bson.ObjectIDFromHex(courseIDStr)
		if err != nil {
			c.JSON(400, gin.H{
				"status":    "Fail",
				"message": "Course ID không hợp lệ",
			})
			return
		}
		data.CourseId = courseID
	}
	teacherIDStr, _ := data.TeacherId.(string)
	if teacherIDStr != "" {
		teacherID, err := bson.ObjectIDFromHex(teacherIDStr)
		if err != nil {
			c.JSON(400, gin.H{
				"status":    "Fail",
				"message": "Teacher ID không hợp lệ",
			})
			return
		}
		data.TeacherId = teacherID
	}
	teacherID, _ := bson.ObjectIDFromHex(teacherIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Course ID không hợp lệ",
		})
		return
	}
	
	collection := models.ClassModel()
	
	// Kiểm tra xem lớp học có bị trùng không
	isDuplicate, err := CheckDuplicateClass(collection, data.Semester, courseID, data.Name, teacherID)
	if err != nil {	
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi kiểm tra dữ liệu",
		})
		return
	}
	
	// Nếu lớp học đã tồn tại, trả về lỗi
	if isDuplicate {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Lớp học đã tồn tại",
		})
		return
	}
	
	// Thêm nếu không bị trùng lặp
	updatedBy, _ := c.Get("ID")
	data.UpdatedBy = updatedBy
	_ , err = collection.UpdateOne(context.TODO(), bson.M{"_id": classID}, bson.M{"$set": data})
	
	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi cập nhật lớp học",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Cập nhật lớp học thành công",
	})
}

// HandleDeleteClass xử lý việc xóa lớp học theo ID
func HandleDeleteClass(c *gin.Context) {
	param := c.Param("id")
	classID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "ID không hợp lệ",
		})
		return
	}
	collection := models.ClassModel()
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": classID})
	if err != nil {
		c.JSON(500, gin.H{
			"status":    "Fail",
			"message": "Lỗi khi xóa lớp học",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":    "Success",
		"message": "Xóa lớp học thành công",
	})
}