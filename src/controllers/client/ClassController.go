package controller_client

import (
	"Go2/models"
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var TIME_INTERVAL = 3 * time.Second
var TIME_MONITOR = 10 * time.Second

// HandleTeacherClasses xử lý việc lấy danh sách lớp học của giáo viên.
func HandleTeacherClasses(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	if user.Role != "teacher" {
		c.JSON(401, gin.H{
			"status":  "Fail",
			"message": "Chỉ giáo viên mới được phép truy cập",
		})
		return
	}
	var classTeacherAll []models.InterfaceClass
	collection := models.ClassModel()
	cursor, err := collection.Find(context.TODO(), bson.M{
		"teacher_id": user.ID,
	})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"status":  "Fail",
				"message": "Giảng viên không quản lý lớp học nào",
			})
			return
		}
		c.JSON(500, gin.H{
			"status":  "Fail",
			"message": "Lỗi khi tìm lớp học",
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &classTeacherAll); err != nil {
		c.JSON(500, gin.H{
			"status":  "Fail",
			"message": "Lỗi khi đọc dữ liệu lớp học",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "Success",
		"data":   classTeacherAll,
	})
}

// HandleStudentClasses xử lý việc lấy danh sách lớp học của sinh viên.
func HandleStudentClasses(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	var classStudentAll []models.InterfaceClassStudent
	collection := models.ClassModel()
	fmt.Println(user)
	cursor, err := collection.Find(context.TODO(), bson.M{
		"listStudent_ms": user.Ms,
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"status":  "Fail",
				"message": "Sinh viên không tham gia lớp học nào",
			})
			return
		}
		c.JSON(500, gin.H{
			"status":  "Fail",
			"message": "Lỗi khi tìm lớp học",
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err := cursor.All(context.TODO(), &classStudentAll); err != nil {
		c.JSON(500, gin.H{
			"status":  "Fail",
			"message": "Lỗi khi đọc dữ liệu lớp học",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "Success",
		"data":   classStudentAll,
	})
}

// HandleUserClasses xử lý việc lấy danh sách lớp học của người dùng.
func HandleUserClasses(c *gin.Context) {
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	if user.Role == "teacher" {
		HandleTeacherClasses(c)
		return
	} else if user.Role == "student" {
		HandleStudentClasses(c)
		return
	}
	c.JSON(400, gin.H{
		"status":  "Fail",
		"message": "Không tìm thấy người dùng",
	})
}

// HandleClassDetail xử lý việc lấy chi tiết lớp học.
func HandleClassDetail(c *gin.Context) {
	paramID := c.Param("id")
	id, err := bson.ObjectIDFromHex(paramID)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	data, _ := c.Get("user")
	user := data.(models.InterfaceAccount)
	var classDetail models.InterfaceClass
	collection := models.ClassModel()
	err = collection.FindOne(context.TODO(), bson.M{
		"_id": id,
	}).Decode(&classDetail)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"status":  "Fail",
				"message": "Không tìm thấy lớp học",
			})
			return
		}
		c.JSON(500, gin.H{
			"status":  "Fail",
			"message": "Lỗi khi tìm lớp học",
		})
		return
	}
	if user.Role == "student" {
		var listStudent = classDetail.ListStudentMs
		for _, studentMs := range listStudent {
			if studentMs == user.Ms {
				c.JSON(200, gin.H{
					"status":  "Success",
					"message": "Lấy lớp học thành công",
					"data":    classDetail,
				})
				return
			}
		}
		c.JSON(401, gin.H{
			"status":  "Fail",
			"message": "Chỉ sinh viên mới được phép truy cập",
		})
		return
	} else if user.Role == "teacher" {
		if user.ID != classDetail.TeacherId {
			c.JSON(401, gin.H{
				"status":  "Fail",
				"message": "Chỉ giáo viên mới được phép truy cập",
			})
			return
		}
		c.JSON(200, gin.H{
			"status":  "Success",
			"message": "Lấy lớp học thành công",
			"data":    classDetail,
		})
		return
	}
	c.JSON(401, gin.H{
		"status":  "Fail",
		"message": "Chỉ sinh viên và giáo viên mới được phép truy cập",
	})
}

// HandleCountDocuments xử lý việc đếm số lượng lớp học của một môn học.
func HandleCountDocuments(c *gin.Context) {
	param := c.Param("id")
	courseID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Không tìm thấy môn học",
		})
		return
	}
	collection := models.ClassModel()
	count, err := collection.CountDocuments(context.TODO(), bson.M{"course_id": courseID})
	if err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Lỗi khi lấy môn học",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "Success",
		"data":   count,
	})
}

func HandleAddClass(c *gin.Context) {
	var data Class4Teacher
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Dữ liệu không hợp lệ",
		})
		return
	}
	courseID, err := bson.ObjectIDFromHex(data.CourseId)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Course ID không hợp lệ",
		})
		return
	}
	user, _ := c.Get("user")
	teacher := user.(models.InterfaceAccount)

	collection := models.ClassModel()

	// Kiểm tra xem lớp học có bị trùng không
	isDuplicate, err := CheckDuplicateClass(collection, data.Semester, courseID, data.Name, teacher.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "Fail",
			"message": "Lỗi khi kiểm tra dữ liệu",
		})
		return
	}

	// Nếu lớp học đã tồn tại, trả về lỗi
	if isDuplicate {
		c.JSON(400, gin.H{
			"status":  "Fail",
			"message": "Lớp học đã tồn tại",
		})
		return
	}

	if isDirectLink(data.CsvURL) {
	} else {
		directLink, err := convertToDirectLink(data.CsvURL)
		if err != nil {
			c.JSON(400, gin.H{
				"status": "fail",
				"message": err.Error(),
			})
			return
		}
		data.CsvURL = directLink;
	}
	
	resp, err := http.Head(data.CsvURL)
	if err != nil {
		fmt.Println("Failed to check file:", err)
	}
	lastModified := resp.Header.Get("Last-Modified")
	result, err := collection.InsertOne(context.TODO(), bson.M{
		"semester":       data.Semester,
		"name":           data.Name,
		"course_id":      courseID,
		"teacher_id":     teacher.ID,
		"createdBy":      teacher.ID,
		"updatedBy":      teacher.ID,
		"csv_url":        data.CsvURL,
		"last_mod":       lastModified,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"status": "fail",
			"message": "fail when insert class",
		})
	}
	Rescollection := models.ResultScoreModel()

	var dataResult InterfaceResultScoreController
	// Tạo mới bản ghi result
	Resresult, err := Rescollection.InsertOne(context.TODO(), bson.M{
		"semester":  data.Semester,
		"course_id": data.CourseId,
		"score":     dataResult.SCORE,
		"class_id":  result.InsertedID,
		"monitor_valid": time.Now().Add(TIME_MONITOR),
		"expired_at": time.Now().AddDate(0, 6, 0),
		"createdBy": teacher.ID,
		"updatedBy": teacher.ID,
		"status": "active",
	})
	fmt.Println(result.InsertedID)
	// monitorAndDownload(c, data.CsvURL, checkInterval, collection, result.InsertedID.(primitive.ObjectID))
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "Fail",
			"message": "Lỗi khi tạo lớp học",
		})
		return
	}
	go monitorAndDownload(c, TIME_INTERVAL, Rescollection, collection, Resresult.InsertedID, result.InsertedID)
	c.JSON(200, gin.H{
		"status":  "success",
		"message": result.InsertedID,
	})
}

func HandleUpdateClassCsvURL(c *gin.Context){
	user, _ := c.Get("user")
	teacher := user.(models.InterfaceAccount)
	param := c.Param("id")
	classID, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	var data InterfaceChangeClassController
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Dữ liệu yêu cầu không hợp lệ",
		})
		return
	}
	fmt.Println()
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
	collection := models.ClassModel()
	var class models.InterfaceClass

	if err := collection.FindOne(context.TODO(), bson.M{"_id": classID}).Decode(&class); err != nil {
		if err == mongo.ErrNoDocuments {
      c.JSON(404, gin.H{
        "status":    "Fail",
        "message": "Không tìm thấy lớp học",
      })
      return
    }
    c.JSON(500, gin.H{
      "status":    "Fail",
      "message": "Lỗi hệ thống",
    })
    return
	}
	if (class.CourseId == courseID && class.Semester == data.Semester && class.Name == data.Name){

	}else{
		isDuplicate, err := CheckDuplicateClass(collection, data.Semester, courseID, data.Name, teacher.ID)
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
	}
	// Kiểm tra xem lớp học có bị trùng không
	if isDirectLink(data.CsvURL) {
		} else {
			directLink, err := convertToDirectLink(data.CsvURL)
			if err != nil {
				c.JSON(400, gin.H{
					"status": "fail",
					"message": err.Error(),
				})
				return
			}
			data.CsvURL = directLink;
		}
	
	// Thêm nếu không bị trùng lặp
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


func CheckDuplicateClass(collection *mongo.Collection, semester string, courseID bson.ObjectID, name string, teacherID bson.ObjectID) (bool, error) {

	// Sử dụng FindOne để kiểm tra xem có bản ghi nào không
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{
		"semester":  semester,
		"course_id": courseID,
		"name":      name,
	}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, nil // Không tìm thấy bản ghi
	} else if err != nil {
		return false, err // Có lỗi khác
	}

	return true, nil // Tìm thấy bản ghi trùng
}

func monitorAndDownload(c *gin.Context, interval time.Duration, Rescollection *mongo.Collection, collection *mongo.Collection, id interface{}, classId interface{}) {
	var lastModified string
	for {
		var res models.InterfaceResult
		err := Rescollection.FindOne(context.TODO(), bson.M{
			"_id": id,
		}).Decode(&res)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "false",
				"message": "Lỗi khi lấy thông tin class !",
			})
		}
		if res.Status == "inactive" {
			fmt.Println("user dont want monitor csv link")
			return;
		}
		if res.MonitorValid.Before(time.Now()){
			_, err := Rescollection.UpdateOne(context.TODO(), bson.M{"_id": res.ID}, bson.M{"$set": bson.M{
				"status": "inactive",
			}},)
			if err != nil {
				fmt.Println("Error")
				return;
			}
			fmt.Println("Time out")
			return;
		}
		var classDetail models.InterfaceClass
		err = collection.FindOne(context.TODO(), bson.M{
			"_id": classId,
		}).Decode(&classDetail)
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "false",
				"message": "Lỗi khi lấy thông tin class !",
			})
		}
		// Gửi yêu cầu HEAD để kiểm tra thông tin file
		resp, err := http.Head(classDetail.CsvURL)
		if err != nil {
			fmt.Println("Failed to check file:", err)
			time.Sleep(interval)
			continue
		}
		modified := resp.Header.Get("Last-Modified")
		if modified == "" {
			time.Sleep(interval)
			continue
		}
		// modified = strings.TrimSpace(modified)
		// Kiểm tra nếu file đã được cập nhật
		if modified != lastModified {
			lastModified = modified
			// collection.UpdateByID(context.TODO(), id, bson.M{"last_mod": lastModified})
			// Tải và phân tích file mới
			records, err := parseCSV(classDetail.CsvURL)
			if err != nil {
				fmt.Println("Failed to parse CSV:", err)
				time.Sleep(interval)
				continue
			}
			_, err = Rescollection.UpdateOne(
				context.TODO(),
				bson.M{"class_id": classId},
				bson.M{"$set": bson.M{
					"score":     records,
					"updatedBy": classDetail.TeacherId,
					"monitor_valid": time.Now().Add(TIME_MONITOR),
					"expired_at": time.Now().AddDate(0, 6, 0),
				}},
			)
			if err != nil {
				fmt.Println("err", err.Error())
				continue
			}
			var list_student []string
			for _, item := range records {
				list_student = append(list_student, item.MSSV)
			}
			_, err = collection.UpdateOne(
				context.TODO(),
				bson.M{"_id": classId},
				bson.M{"$set": bson.M{
					"listStudent_ms":  list_student,
				}},
			)
			if err != nil {
				fmt.Println("err", err.Error())
				continue
			}
			// Hiển thị dữ liệu mới
			fmt.Println("Updating...")
			// for _, record := range records {
			// 	fmt.Printf("MSSV: %s\n", record.MSSV)
			// 	fmt.Printf("Result: %v\n", record.Data)
			// 	fmt.Println("---")
			// }
		} else {
			fmt.Println("No updates found.")
		}

		// Đợi trước khi kiểm tra lại
		time.Sleep(interval)
	}
}

func parseCSV(url string) ([]StudentRecord, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch CSV: %w", err)
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read headers: %w", err)
	}

	var studentRecords []StudentRecord
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		student := StudentRecord{}
		for i, value := range record {
			header := headers[i]
			switch {
			case header == "MSSV":
				student.MSSV = value
			case strings.HasPrefix(header, "BT") && !strings.HasPrefix(header, "BTL"):
				// student.Data.BT = append(student.Data.BT, value)
				if value != "" {
					float64Value, err := strconv.ParseFloat(value, 32)
					float32Value := float32(float64Value)
					if err != nil {
						fmt.Println("Error converting string to float:", err)
						return nil, nil
					}

					student.Data.BT = append(student.Data.BT, float32Value)
				}else {
					student.Data.BT = append(student.Data.BT, 0)
				}
			case strings.HasPrefix(header, "TN"):
				// student.Data.BT = append(student.Data.TN, value)
				if value != "" {
					float64Value, err := strconv.ParseFloat(value, 32)
					float32Value := float32(float64Value)
					if err != nil {
						fmt.Println("Error converting string to float:", err)
						return nil, nil
					}
					student.Data.TN = append(student.Data.TN, float32Value)
				} else {
					student.Data.TN = append(student.Data.TN, 0)
				}

			case strings.HasPrefix(header, "BTL"):
				// student.Data.BT = append(student.Data.BTL, value)
				if value != "" {
					float64Value, err := strconv.ParseFloat(value, 32)
					float32Value := float32(float64Value)
					if err != nil {
						fmt.Println("Error converting string to float:", err)
						return nil, nil
					}
					student.Data.BTL = append(student.Data.BTL, float32Value)
				} else{
					student.Data.BTL = append(student.Data.BTL, 0)
				}
			case header == "GK":
				// student.Data.GK = value
				if value != "" {
					float64Value, err := strconv.ParseFloat(value, 32)
					float32Value := float32(float64Value)
					if err != nil {
						fmt.Println("Error converting string to float:", err)
						return nil, nil
					}
					student.Data.GK = float32Value
				}else{
					student.Data.GK = 0
				}

			case header == "CK":
				// student.Data.CK = value
				if value != "" {
					float64Value, err := strconv.ParseFloat(value, 32)
					float32Value := float32(float64Value)
					if err != nil {
						fmt.Println("Error converting string to float:", err)
						return nil, nil
					}
					student.Data.CK = float32Value
				}else{
					student.Data.CK = 0
				}
			}
		}

		studentRecords = append(studentRecords, student)
	}

	return studentRecords, nil
}

func convertToDirectLink(shareLink string) (string, error) {
	if !strings.Contains(shareLink, "drive.google.com") {
		return "", fmt.Errorf("liên kết không phải của Google Drive")
	}

	var fileID string
	if strings.Contains(shareLink, "/file/d/") {
		parts := strings.Split(shareLink, "/file/d/")
		if len(parts) > 1 {
			fileIDParts := strings.Split(parts[1], "/")
			fileID = fileIDParts[0]
		}
	} else if strings.Contains(shareLink, "id=") {
		parts := strings.Split(shareLink, "id=")
		fileID = strings.Split(parts[1], "&")[0]
	} else {
		return "", fmt.Errorf("không tìm thấy ID file trong liên kết")
	}

	directLink := fmt.Sprintf("https://drive.google.com/uc?export=download&id=%s", fileID)
	return directLink, nil
}

func isDirectLink(link string) bool {
	// Phân tích cú pháp URL
	parsedURL, err := url.Parse(link)
	if err != nil {
		return false
	}

	// Kiểm tra host có phải drive.google.com không
	if !strings.Contains(parsedURL.Host, "drive.google.com") {
		return false
	}

	// Kiểm tra đường dẫn và các tham số truy vấn
	path := parsedURL.Path
	query := parsedURL.Query()

	if strings.HasPrefix(path, "/uc") {
		// Kiểm tra tham số export=download
		if query.Get("export") == "download" && query.Get("id") != "" {
			return true
		}
	}

	return false
}