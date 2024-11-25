package controllers_admin

import (
	"be-go-2/models"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateClass(c *gin.Context) {
	var data  struct{
		Semester      string   `json:"semester"`
		Name          string   `json:"name"` // nhom lop
		CourseId      string   `json:"course_id"`
		ListStudentId []string `json:"listStudent_id"`
		TeacherId     string   `json:"teacher_id"`
	}

	//binding data xem co loi khong
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"position": "binding data error",
		})
		return
	}
	// kiem tra student id có hợp lệ không
	studentCollection := models.StudentModel()
	ListStudentObjID := []bson.ObjectID{}
	for _, id := range data.ListStudentId {
		stuID, err := bson.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
                "message": err.Error(),
                "position": "encoding student id error",
            })
            return
		}
		check, err := CheckUserExists(studentCollection, stuID, "student")
		if !check {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
                "position": "checking student id error",
			})
			return
		}
        ListStudentObjID = append(ListStudentObjID, stuID)
    }

    // kiem tra teacher id co ton tai khong
    teacherCollection := models.TeacherModel()

	teacherID, err := bson.ObjectIDFromHex(data.TeacherId)
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"position": "encoding teacher id error",
		})
		return
	}
	check, err := CheckUserExists(teacherCollection, teacherID, "teacher")
	if !check {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
            "position": "checking teacher id error",
		})
		return
	}

	// kiem tra courseID co ton tai khong
	courseID, err := bson.ObjectIDFromHex(data.CourseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"position": "encoding course id error",
		})
		return
	}
	var course bson.M
	courseCollection := models.CourseModle()
	err = courseCollection.FindOne(context.Background(), bson.M{
		"_id": courseID,
	}).Decode(&course)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{
            "message": "Course id not found",
            "position": "checking course id error",
        })
        return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "message": err.Error(),
            "position": "checking course id error",
        })
        return
	}


	classCollection := models.ClassModel()
	
	//kiem tra lop da ton tai hay chua
	check, err = CheckClassExists(classCollection, courseID, data.Name, data.Semester)
	if !check {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
            "position": "checking class error",
		})
		return
	}

	//kiem tra adminID da gui post
	id, exist := c.Get("ID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID is required, need set ID in middleware",
            "position": "adminid sent request error",
		})
		return
	}
	//tao lop moi
	result, err := classCollection.InsertOne(context.Background(), bson.M{
		"Name": data.Name,
		"CourseId": courseID,
		"Semester": data.Semester,
		"ListStudentId": ListStudentObjID,
		"TeacherId": data.TeacherId,
		"CreatedBy": id,
		"UpdatedBy": id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"position": "inserting class error",
		})
	}
	// return success
	c.JSON(http.StatusOK, gin.H{
		"message": "Created class successfully",
		"class_id": result.InsertedID,
	})
}

func CheckUserExists(col *mongo.Collection, id bson.ObjectID, role string) (bool, error){
	if role == "student" || role == "teacher"{
		var user bson.M
		err := col.FindOne(context.Background(), bson.M{
			"_id": id,
		}).Decode(&user)
		if err == mongo.ErrNoDocuments {
			return false, errors.New(role + " id not found")
		}else if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, errors.New("invalid role")
}

func CheckClassExists(col *mongo.Collection, courseID bson.ObjectID, classID string, semester string) (bool, error) {
	var class bson.M
    err := col.FindOne(context.Background(), bson.M{
        "ClassName": classID,
        "CourseId": courseID,
        "semester": semester,
    }).Decode(&class)
    if err == mongo.ErrNoDocuments {
        return true, nil
    } else if err!= nil {
        return false, err
    }
    return false, errors.New("class adready exists")
}