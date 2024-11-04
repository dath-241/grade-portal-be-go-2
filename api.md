### API

### Authentication

- `POST /auth/login` : đăng nhập bằng oauth2

```json
{
  "idToken": "string",
  "role": "string" // admin, student, teacher
}
```

- `POST /auth/logout` : đăng xuất, xóa session.

```json
{
  "sessionId": "string"
}
```
### Lưu ý : 
- Sau khi người dùng đăng nhập thì chỉ cò 2 role là `admin` và `client` , sau khi `client` được phân quyền thì sẽ chia ra thành `teacher` và `client` còn lại chính là `student`.

## Admin

### Teacher Management

- `POST /admin/teacher` : Tạo tài khoản giáo viên mới.

```json
{
  "email": "string",
  "name": "string",
  "teacherID": "string",
  "faculty": "string",
  "role": "string" //teacher
}
```

- `PUT /admin/teacher/:teacherID` : Cập nhật quyền cho giáo viên.

```json
{
  "permissions": ["string", "string"] // ["create_class","upload_class"]
}
```

- `DELETE /admin/teacher/:teacherID`: Xóa tài khoản giáo viên.

```json
{
  "message": "Teacher with ID ... has been deleted !"
}
```

- `GET /admin/teachers` : Lấy danh sách tất cả giáo viên.

```json
[
  {
    "teacherID": "string",
    "name": "string",
    "email": "string",
    "faculty": "string",
    "role": "string", // teacher
    "permissions": ["string", "string"] // ["create_class", "update_class"]
  },
  {
    "teacherID": "string",
    "name": "string",
    "email": "string",
    "faculty": "string",
    "role": "string", // teacher
    "permissions": ["string", "string"] // ["create_class", "update_class"]
  }
]
```

- `GET /admin/teacher/:teacherID` : Lấy chi tiết giáo viên theo ID.

### Course & Class Management

- `POST /admin/course` : Tạo khóa học mới.

```json
{
  "courseID": "string", // CO2003
  "courseName": "string", // DSA
  "credit": 3, // integer
  "semester": "string",
  "description": "string"
}
```

- `PUT /admin/course/:courseID` : Cập nhật thông tin khóa học.

```json
{
  "courseName": "string", //new maybe
  "credit": 4, // integer
  "description": "string" // new maybe
}
```

- `DELETE /admin/course/:courseID` : Xóa khóa học.

```json
{
  "message": "string" //Course with ID CO2003 has been deleted
}
```

- `GET /admin/courses` : Lấy danh sách tất cả khóa học.

```json
[
  {
    "courseID": "string", // PH1003
    "courseName": "string", // Physics
    "credit": 4,
    "description": "string"
  },
  {
    "courseID": "string", // PH1003
    "courseName": "string", // Physics
    "credit": 4,
    "description": "string"
  }
]
```

- `POST /admin/class`: Tạo lớp học mới.

```json
{
  "classID": "string",
  "courseID": "string",
  "teacherID": "string",
  "semester": "string",
  "listStudentID": [] "string" // ["studentCode1","studentCode2"]
}
```

- `PUT /admin/class/:classID` : Cập nhật thông tin lớp học.

```json
{
    "className" : "string",
    "listStudentID" :[] "string"
}
```

- `DELETE /admin/class/:classID` : Xóa lớp học.
- `GET /admin/classes` : Lấy danh sách tất cả lớp học.

## Teacher

### Course Management

- `GET /teacher/courses` : Lấy danh sách các môn học được phân công.

```json
{
  "semester": "241",
  "courses": [
    {
      "courseID": "string",
      "courseName": "string", // Môn A
      "classID": "string", // L01, L02, ...
      "credit": "int",
      "description": "string"
    },
    {
      "courseID": "string",
      "courseName": "string", // Môn B
      "classID": "string",
      "credit": "int",
      "description": "string"
    }
  ]
}
```

- `GET /teacher/course/:courseID` : Xem thông tin chi tiết môn học.
- `GET /teacher/course/:courseID/class/:classID` : Xem thông tin lớp của môn học.

### Grade Management

- `POST /teacher/course/:courseID/class/:classID/upload` : Tải lên file CSV điểm.

```json
{
  "file": "base_64_encoded_csv_content"
}
```

- `GET /teacher/course/:courseID/class/:classID/gradeSheet` : Lấy danh sách điểm.

```json
{
  "gradeSheet": [
    {
      "studentID": "string", // Mã số sinh viên
      "studentName": "string",
      "grades": {
        "Homework": [],
        "Lab": [],
        "Assignment": [],
        "Midterm": ,
        "Final":
      },
      "averageScore":
    },
    {
      "studentID": "string", // Mã số sinh viên
      "studentName": "string",
      "grades": {
        "Homework": [],
        "Lab": [],
        "Assignment": [],
        "Midterm": ,
        "Final":
      },
      "averageScore":
    }
  ]
}
```

- `PATCH /teacher/course/:courseID/class/:classID/grade/:studentID` : Cập nhật điểm cho sinh viên.

```json
"grades": {
    "Homework": "float",
    "Lab": "float",
    "Assignment": "float",
    "Midterm": "float",
    "Final": "float"
  }
```

- `DELETE /teacher/course/:courseID/class/:classID/grade/:studentID` : Xóa điểm của sinh viên trong lớp thuộc môn học.

## Student

### Course Enrollment

- `GET /student/courses/` : Lấy danh sách môn học đã đăng ký.

```json
{
  "semesters": [
    {
      "semester": "241",
      "courses": [
        {
          "courseId": "string",
          "courseName": "string",
          "teacherName": "string",
          "classID": "string",
          "credit": 3
        },
        {
          "courseId": "string",
          "courseName": "string",
          "teacherName": "string",
          "classID": "string",
          "credit": 3
        }
      ]
    },
    {
      "semester": "242",
      "courses": [
        {
          "courseId": "string",
          "courseName": "string",
          "teacherName": "string",
          "classID": "string",
          "credit": 3
        }
      ]
    }
  ]
}

```

### Grade Viewing

- `GET /student/course/:courseID/grades` : Lấy bảng điểm của môn học đó.

```json
{
  "grades": {
    "Homework": "float",
    "Lab": "float",
    "Assignment": "float",
    "Midterm": "float",
    "Final": "float",
    "averageScore": "float"
  }
}
```

- `GET /student/grades` : Lấy tất cả điểm của sinh viên trong môn học đó.
```json 
{
  "grades": [
    {
      "courseID": "string",
      "averageScore": "float"
    },
    {
      "courseID": "string",
      "averageScore": "float"
    }
  ]
}
```

### Hall of Fame

- `GET /halloffame` : Lấy danh sách top sinh viên theo điểm số.

```json
{
  "topStudents": [
    {
      "studentID": "string",
      "name": "string",
      "averageScore": 9.6
    },
    {
      "studentID": "string",
      "name": "string",
      "averageScore": 9.1
    },
    {
      "studentID": "string",
      "name": "string",
      "averageScore": 8.8
    }
  ]
}
```
