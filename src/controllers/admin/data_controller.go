package controllers_admin

import (
	"time"
)

type User struct {
	ID          string       `json:"_id" bson:"_id"`
	Name        Name         `json:"name"`
	Email       string       `json:"email"`
	LastLogin   time.Time    `json:"LastLogin"`
	Role        string       `json:"role"`
	UpdatedAt   time.Time    `json:"updated_at"`
	CreatedAt   time.Time    `json:"created_at"`
	AdminInfo   *AdminInfo   `json:"admin_info,omitempty"`   // tồn tại nếu là admin
	StudentInfo *StudentInfo `json:"student_info,omitempty"` // tồn tại nếu là student
	TeacherInfo *TeacherInfo `json:"teacher_info,omitempty"` // tồn tại nếu là teacher
}

type Name struct {
	LastName string `json:"LastName"`
	MFName   string `json:"MFName"` // minit name + first name
}

type AdminInfo struct {
	AdminID string `json:"admin_id" bson:"admin_id"`
}

type StudentInfo struct {
	StudentID string   `json:"student_id" bson:"student_id"`
	ClassID   []string `json:"class_id" bson:"class_id"` // danh sách ID lớp học
}

type TeacherInfo struct {
	TeacherID     string   `json:"teacher_id" bson:"teacher_id"`
	ManageClassID []string `json:"manage_class_id" bson:"manage_class_id"` // danh sách ID lớp quản lý
}

type Class struct {
	ID        string    `json:"_id" bson:"_id"`
	ClassName string    `json:"class_name" bson:"class_name"`
	TeacherID string    `json:"teacher_id" bson:"teacher_id"`
	StudentID []string  `json:"student_id" bson:"student_id"` // danh sách ID sinh viên
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	Status    string    `json:"status" bson:"status"` // open, close
	Semester  string    `json:"semester" bson:"semester"`
	CourseID  string    `json:"course_id" bson:"course_id"`
}

type Course struct {
	ID         string   `json:"_id" bson:"_id"`
	CourseName string   `json:"course_name" bson:"course_name"`
	Credit     float32  `json:"credit" bson:"credit"`
	ClassID    []string `json:"class_id" bson:"class_id"` // danh sách ID lớp học
}

type GradeSheet struct {
	ID        string    `json:"_id" bson:"_id"`
	StudentID string    `json:"student_id" bson:"student_id"`
	ClassID   string    `json:"class_id" bson:"class_id"`
	Status    string    `json:"status" bson:"status"` // open, close
	Grade     Grade     `json:"grade" bson:"grade"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Grade struct {
	Final      float32   `json:"final" bson:"final"`
	Midterm    float32   `json:"midterm" bson:"midterm"`
	Assignment []float32 `json:"assignment" bson:"assignment"`
	Exercise   []float32 `json:"exercise" bson:"exercise"`
}
