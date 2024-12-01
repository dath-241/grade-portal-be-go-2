package controller_admin

import (
	"Go2/helper"
	"Go2/models"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Tính điểm trung bình của sinh viên trong một khóa học cụ thể.
func CalculateAvgStudentScores(semester string, courseID bson.ObjectID) ([]avgStudentScore, error) {
	coursesCollection := models.CourseModel()

	// Tìm khóa học với courseID
	var course models.InterfaceCourse
	err := coursesCollection.FindOne(context.TODO(), bson.M{"_id": courseID}).Decode(&course)
	if err != nil {
		return make([]avgStudentScore, 0), errors.New("lỗi khi tìm khóa học")
	}
	HS := course.HS

	// Tìm danh sách điểm của sinh viên trong học kỳ và khóa học cụ thể
	scoresCollection := models.ResultScoreModel()
	cursor, err := scoresCollection.Find(context.TODO(), bson.M{"course_id": courseID, "semester": semester})
	if err != nil {
		return make([]avgStudentScore, 0), errors.New("lỗi khi tìm điểm số")
	}
	defer cursor.Close(context.TODO())

	var resultScores []models.InterfaceResult
	if err = cursor.All(context.TODO(), &resultScores); err != nil {
		return make([]avgStudentScore, 0), errors.New("lỗi khi giải mã điểm số")
	}

	// Khởi tạo và gán giá trị cho slice avgScores
	totalSize := 0
	for _, result := range resultScores {
		totalSize += len(result.SCORE)
	}
	i := 0
	avgScores := make([]avgStudentScore, totalSize)
	for _, result := range resultScores {
		for _, score := range result.SCORE {
			avgScores[i].MSSV = score.MSSV
			avgScores[i].AvgScore = helper.AvgScore(score.Data, HS[:])
			i++
		}
	}
	return avgScores, nil
}

// SortAvgScores sắp xếp danh sách điểm trung bình của sinh viên.
func SortAvgScores(avgScores []avgStudentScore) []avgStudentScore {
	if len(avgScores) <= 1 {
		return avgScores
	}

	mid := len(avgScores) / 2
	left := SortAvgScores(avgScores[:mid])
	right := SortAvgScores(avgScores[mid:])

	return merge(left, right)
}

//merge 2 tập con theo giảm dần
func merge(left, right []avgStudentScore) []avgStudentScore {
	var result []avgStudentScore
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].AvgScore > right[j].AvgScore { // Sắp xếp giảm dần
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Thêm các phần tử còn lại
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// IsDuplicateHOF kiểm tra xem Hall of Fame có bị trùng không.
func IsDuplicateHOF(collection *mongo.Collection, semester string, courseID bson.ObjectID) bool {
	filter := bson.M{
		"semester":  semester,
		"course_id": courseID,
	}

	// Kiểm tra xem có bản ghi nào không
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false // Không tìm thấy bản ghi
	} else if err != nil {
		return false // Có lỗi khác
	}

	return true
}

// HandleCreateHallOfFame xử lý việc tạo Hall of Fame.
func HandleCreateHallOfFame(c *gin.Context) {
	scoresCollection := models.ResultScoreModel()
	var results []models.InterfaceResult
	semester := helper.SetSemester()
	cursor, err := scoresCollection.Find(context.TODO(), bson.M{
		"semester": semester.PREV,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"status":    "Fail",
			"message": "Lỗi tìm kiếm bản ghi",
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &results); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":    "Fail",
			"message": "Lỗi tìm kiếm bản ghi",
		})
		return
	}
	processed := make(map[string]bool)
	collection := models.HallOfFameModel()
	for _, result := range results {
		key := result.Semester + "-" + result.CourseID.Hex()
		if found := processed[key]; !found {
			processed[key] = true
			avgStudentScores, err := CalculateAvgStudentScores(result.Semester, result.CourseID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":    "Fail",
					"message": err.Error(),
				})
				return
			}
			studentHOF := SortAvgScores(avgStudentScores)
			var data bson.A
			length := min(10, len(studentHOF))
			for i := 0; i < length; i++ {
				student := studentHOF[i]
				data = append(data, bson.M{"mssv": student.MSSV, "dtb": student.AvgScore})
			}
			if !IsDuplicateHOF(collection, result.Semester, result.CourseID) {
				collection.InsertOne(context.TODO(), bson.M{
					"semester":  result.Semester,
					"course_id": result.CourseID,
					"data":      data,
				})
			} else {
				filter := bson.M{
					"semester":  result.Semester,
					"course_id": result.CourseID,
				}
				update := bson.M{
					"$set": bson.M{
						"data": data},
				}
				collection.UpdateOne(context.TODO(), filter, update)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    "Success",
		"message": "Cập nhật Hall Of Fame thành công",
	})
}

// HandleGetPrevSemesterAllHallOfFame xử lý việc lấy Hall of Fame của học kỳ trước.
func HandleGetPrevSemesterAllHallOfFame(c *gin.Context) {
	collection := models.HallOfFameModel()
	semester := helper.SetSemester().PREV
	var hallOfFameData InterfaceHallOfFame

	var tierData []InterfaceTier
	filter := bson.M{
		"semester": semester,
	}

	// Sử dụng Find để lấy tất cả các tài liệu khớp với filter
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":    "Fail",
			"message": "Đã xảy ra lỗi khi truy vấn dữ liệu",
		})
		return
	}
	defer cursor.Close(context.TODO())

	// Duyệt qua các tài liệu và thêm chúng vào results
	for cursor.Next(context.TODO()) {
		var data InterfaceTier
		if err := cursor.Decode(&data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":    "Fail",
				"message": "Đã xảy ra lỗi khi giải mã dữ liệu"})
			return
		}
		tierData = append(tierData, data)
	}

	// Kiểm tra xem có kết quả nào không
	if len(tierData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":    "Fail",
			"message": "Không tìm thấy dữ liệu cho học kỳ trước"})
		return
	} else {
		hallOfFameData.Semester = semester
		hallOfFameData.Tier = tierData
	}

	// Trả về tất cả các bản ghi nếu tìm thấy
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Lấy hall of fame thành công",
		"data":    hallOfFameData,
	})
}
