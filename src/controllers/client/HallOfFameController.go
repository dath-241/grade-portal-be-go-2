package controller_client

import (
    "Go2/helper"
    "Go2/models"
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/v2/bson"
)

// GetAllPrevSemester lấy tất cả các bản ghi hall of fame cho học kỳ trước
func GetAllPrevSemester(c *gin.Context) {
    // Lấy collection từ model
    collection := models.HallOfFameModel()
    
    // Lấy giá trị học kỳ trước
    semester := helper.SetSemester().PREV
    
    // Khởi tạo các cấu trúc dữ liệu để lưu kết quả
    var hallOfFameData InterfaceHallOfFame
    var tierData []InterfaceTier

    // Tạo bộ lọc để truy vấn collection
    filter := bson.M{"semester": semester}

    // Truy vấn collection bằng bộ lọc
    cursor, err := collection.Find(context.TODO(), filter)
    if err != nil {
        // Trả về phản hồi lỗi nếu truy vấn thất bại
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":    "Fail",
            "message": "Đã xảy ra lỗi khi truy vấn dữ liệu"})
        return
    }
    defer cursor.Close(context.TODO())

    // Duyệt qua cursor và giải mã từng tài liệu vào cấu trúc dữ liệu
    for cursor.Next(context.TODO()) {
        var data InterfaceTier
        if err := cursor.Decode(&data); err != nil {
            // Trả về phản hồi lỗi nếu giải mã thất bại
            c.JSON(http.StatusInternalServerError, gin.H{
                "status":    "Fail",
                "message": "Đã xảy ra lỗi khi giải mã dữ liệu"})
            return
        }
        tierData = append(tierData, data)
    }

    // Kiểm tra xem có dữ liệu nào được tìm thấy không
    if len(tierData) == 0 {
        // Trả về phản hồi không tìm thấy nếu không có dữ liệu
        c.JSON(http.StatusNotFound, gin.H{
            "status":    "Fail",
            "message": "Không tìm thấy dữ liệu cho học kỳ trước"})
        return
    }

    // Điền dữ liệu hall of fame với kết quả
    hallOfFameData.Semester = semester
    hallOfFameData.Tier = tierData

    // Trả về phản hồi thành công với dữ liệu
    c.JSON(http.StatusOK, gin.H{
        "status":  "Success",
        "message": "Lấy hall of fame thành công",
        "data":    hallOfFameData,
    })
}