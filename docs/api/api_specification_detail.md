# Mô tả chi tiết API
## Tổng quan
Mô tả các tính năng chính và chức năng của API. 
## Xác thực
API này sử dụng cơ chế xác thực Bearer Token. Mọi yêu cầu đến API phải bao gồm một header Authorization có chứa token hợp lệ (Ngoại trừ api đăng nhập).
## Các Endpoints
### 1. Admin
#### 1.1. Authorization:
##### a. Login
- **Mô tả:** Được sử dụng để đăng nhập cho admin
- **URL:** `/admin/api/login`
- **Phương thức:** `POST`
- **Request body:**
  ```json
  {
    "idToken": ""
  }

- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "token": ""
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Token không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Tài khoản không tồn tại"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi lấy dữ liệu từ server database."
  }

##### b. Logout
- **Mô tả:** Được sử dụng để admin đăng xuất
- **URL:** `/admin/api/logout`
- **Phương thức:** `POST`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "message": "Đăng xuất thành công"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Token không hợp lệ"
  }

##### c. CreateAdmin
- **Mô tả:** Được sử dụng để tạo tài khoản cho admin
- **URL:** `/admin/api/create`
- **Phương thức:** `POST`
- **Xác thực**: Bearer `<token>`
-  **Request body:**
  ```json
{
    "email": "test@hcmut.edu.vn" ,
    "name" :  "admin2" ,  
    "faculty": "KHMT"  , 
    "ms":      "2222222" 
}
```
- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "message": "Tạo tài khoản admin thành công"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Token không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi tạo tài khoản admin."
  }

##### d. ProfileAdmin
- **Mô tả:** Được sử dụng để hiển thị thông tin admin đang đăng nhập.
- **URL:** `/admin/api/profile`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "message": "Thành công",
    "data": "ID": "674d498fce1b803519045d80",
	        "Email": "test@hcmut.edu.vn",
	        "Name": "test",
	        "Ms": "2222122",
	        "Faculty": "KHMT",
	        "CreatedBy": "674d498fce1b803519045d80"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Token không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Không lấy được thông tin người dùng trong dữ liệu."
  }

#### 1.2. Account (Những hoạt động liên quan tới tài khoản giảng viên và sinh viên):
##### a. CreateAccount
- **Mô tả:** Được sử dụng để tạo tài khoản giảng viên và sinh viên.
- **URL:** `/admin/api/account/create`
- **Phương thức:** `POST`
- **Xác thực**: Bearer `<token>`
-  **Request body:**
  ```json
  [
	  {
	    "email": "test22@hcmut.edu.vn",
	    "name": "test",
	    "ms": "2222222",
	    "faculty": "BDCN",
	    "role": "teacher"
	  },
	  {
	    "email": "test33@hcmut.edu.vn",
	    "name": "test",
	    "ms": "3333333",
	    "faculty": "BDCN",
	    "role": "student"
	  }
  ]
```
- **Response:**
- 200 OK:
  ```json
  {
      "invalidAccounts": null,
      "status": "Success",
      "validAccounts": [
        {
          "email": "test22@hcmut.edu.vn",
          "name": "test",
          "ms": "2222222",
          "faculty": "BDCN",
          "role": "teacher"
      },
      {
          "email": "test33@hcmut.edu.vn",
          "name": "test",
          "ms": "3333333",
          "faculty": "BDCN",
          "role": "student"
      }
      ]
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi tạo tài khoản."
  }

##### b. GetAllTeacher
- **Mô tả:** Được sử dụng để lấy tất cả thông tin tất cả giảng viên.
- **URL:** `/admin/api/account/teacher`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "message": "Tìm tài khoản thành công",
    "data": [
        {
            "ID": "672b85b8226ae67ef9aaa006",
            "Email": "anh.hoang04k22@hcmut.edu.vn",
            "Name": "ANH HOÀNG TRUNG",
            "Ms": "2210001",
            "Faculty": "KHMT",
            "Role": "teacher",
            "CreatedBy": "670de16d6fbedac9e3b8d00a",
            "ExpiredAt": "2029-11-06T15:05:28.186Z"
        },
        {
            "ID": "672b85b8226ae67ef9aaa005",
            "Email": "thuanle@hcmut.edu.vn",
            "Name": "THUẬN LÊ ĐÌNH",
            "Ms": "2210000",
            "Faculty": "KHMT",
            "Role": "teacher",
            "CreatedBy": "670de16d6fbedac9e3b8d00a",
            "ExpiredAt": "2029-11-06T15:05:28.186Z"
        }
    ],
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không có tài khoản nào trong cơ sở dữ liệu"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi lấy dữ liệu từ cơ sở dữ liệu"
  }

##### c. GetAllStudent
- **Mô tả:** Được sử dụng để lấy thông tin tất cả tài khoản của sinh viên.
- **URL:** `/admin/api/account/student`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "message": "Tìm tài khoản thành công",
    "data": [
        {
            "ID": "672b85b8226ae67ef9aaa007",
            "Email": "anh.phamviet241103@hcmut.edu.vn",
            "Name": "ANH PHẠM VIỆT",
            "Ms": "2210002",
            "Faculty": "KHMT",
            "Role": "student",
            "CreatedBy": "670de16d6fbedac9e3b8d00a",
            "ExpiredAt": "2029-11-06T15:05:28.186Z"
        },
        {
            "ID": "672b85b8226ae67ef9aaa010",
            "Email": "dang.nguyen2210737cs@hcmut.edu.vn",
            "Name": "ĐĂNG NGUYỄN HUỲNH HẢI",
            "Ms": "2210011",
            "Faculty": "KHMT",
            "Role": "student",
            "CreatedBy": "670de16d6fbedac9e3b8d00a",
            "ExpiredAt": "2029-11-06T15:05:28.186Z"
        }
    ],
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không có tài khoản nào trong cơ sở dữ liệu"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi lấy dữ liệu từ cơ sở dữ liệu"
  }

##### d. GetAccountByID
- **Mô tả:** Được sử dụng để lấy thông tin tài khoản dựa trên id của giảng viên hoặc sinh viên.
- **URL:** `/admin/api/account/{id}`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                                |
| ---- | ------------ | -------- | ------------------------------------ |
| `id` | String       | Có       | ID của người dùng cần lấy thông tin. |

- **Response:**
- 200 OK:
  ```json
  {
    "data": {
        "ID": "672b85b8226ae67ef9aaa006",
        "Email": "anh.hoang04k22@hcmut.edu.vn",
        "Name": "ANH HOÀNG TRUNG",
        "Ms": "2210001",
        "Faculty": "KHMT",
        "Role": "teacher",
        "CreatedBy": "670de16d6fbedac9e3b8d00a",
        "ExpiredAt": "2029-11-06T15:05:28.186Z"
    },
    "message": "Tìm tài khoản thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy tài khoản"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi lấy tài khoản từ cơ sở dữ liệu"
  }

##### e. DeletedAccountByID
- **Mô tả:** Được sử dụng để xóa tài khoản giảng viên và sinh viên.
- **URL:** `/admin/api/account/delete/{id}`
- **Phương thức:** `DELETE`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                                |
| ---- | ------------ | -------- | ------------------------------------ |
| `id` | String       | Có       | ID của người dùng cần lấy thông tin. |

- **Response:**
- 200 OK:
  ```json
  {
    "message": "Xóa tài khoản thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi xóa tài khoản"
  }

##### f. ChangeAccountInfoByID
- **Mô tả:** Được sử dụng để cập nhật giá trị một vài thuộc tính cho giảng viên và sinh viên
- **URL:** `/admin/api/account/change/{id}`
- **Phương thức:** `PATCH`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                                |
| ---- | ------------ | -------- | ------------------------------------ |
| `id` | String       | Có       | ID của người dùng cần lấy thông tin. |
-  **Request body:**
  ```json
  {
      "name": "Le Van A",
      "faculty": "KTHH",
      "role": "teacher"
  }
```
- **Response:**
- 200 OK:
  ```json
  {
    "message": "Cập nhật tài khoản thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
	  "status":    "Fail",
    "message": "Lỗi khi cập nhật tài khoản vào cơ sở dữ liệu"
  }

#### 1.3. Class (Những hoạt động liên quan tới lớp học):
##### a. CreateClass
- **Mô tả:** Được sử dụng để tạo ra lớp học mới.
- **URL:** `/admin/api/class/create`
- **Phương thức:** `POST`
- **Xác thực**: Bearer `<token>`
-  **Request body:**
  ```json
  { 
    "semester":       "HK242",
    "name":           "L01",
    "course_id":      "672b86fd226ae67ef9aaa045",
    "listStudent_ms": [
        "2222222",
        "2210011",
        "2210012",
        "2211417",
        "12112322313100000"
    ],
    "teacher_id":     "672b85b8226ae67ef9aaa006"
  }
```
- **Response:**
- 200 OK:
  ```json
  {
      "status": "Success",
      "message": "Tạo lớp học thành công"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi tạo lớp học"
  }

##### b. GetAllClassesByAccountID
- **Mô tả:** Được sử dụng để lấy thông tin của tất cả lớp học thông qua id của giảng viên hoặc sinh viên. Nếu là giảng viên thì lấy lớp giảng viên đó đang phụ trách, nếu là sinh viên thì lấy thông tin lớp sinh viên đó đang tham gia.
- **URL:** `/admin/api/class/account/{id}`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                                |
| ---- | ------------ | -------- | ------------------------------------ |
| `id` | String       | Có       | ID của người dùng cần lấy thông tin. |

- **Response:**
- 200 OK:
  ```json
  {
      "message": "Lấy lớp học thành công",
      "status": "Success",
      "data": [
          {
              "ID": "67375642ec72767cfd44630c",
              "Semester": "HK241",
              "Name": "L03",
              "CourseId": "672b86fd226ae67ef9aaa045",
              "ListStudentMs": null,
              "TeacherId": "672b85b8226ae67ef9aaa005",
              "CreatedBy": "6730d78984648a9d5ba4d2ec",
              "UpdatedBy": "6730d78984648a9d5ba4d2ec"
          },
          {
              "ID": "672b87af226ae67ef9aaa047",
              "Semester": "HK233",
              "Name": "L01",
              "CourseId": "672b86fd226ae67ef9aaa045",
              "ListStudentMs": [
                  "2210001",
                  "2210002"
              ],
              "TeacherId": "672b85b8226ae67ef9aaa005",
              "CreatedBy": "670de16d6fbedac9e3b8d00a",
              "UpdatedBy": "670de16d6fbedac9e3b8d00a"
          }
      ]
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi hệ thống"
  }

##### c. GetClassByCourseID
- **Mô tả:** Được sử dụng để lấy thông tin của tất cả lớp học dựa vào id khóa học.
- **URL:** `/admin/api/class/course/{id}`
- **Phương thức:** `GET`
-  **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                              |
| ---- | ------------ | -------- | ---------------------------------- |
| `id` | String       | Có       | ID của khóa học cần lấy thông tin. |
- **Response:**
- 200 OK:
  ```json
  {
      "message": "Lấy lớp học thành công",
      "status": "Success",
      "data": [
          {
              "ID": "6744928c7ac9374bba79314d",
              "Semester": "HK241",
              "Name": "L02",
              "CourseId": "672b86fd226ae67ef9aaa045",
              "ListStudentMs": [
                  "2210001",
                  "2210002"
              ],
              "TeacherId": "672b85b8226ae67ef9aaa005",
              "CreatedBy": "000000000000000000000000",
              "UpdatedBy": "000000000000000000000000"
          },
          {
              "ID": "67375815ec72767cfd44630d",
              "Semester": "HK241",
              "Name": "L04",
              "CourseId": "672b86fd226ae67ef9aaa045",
              "ListStudentMs": [
                  "2210001",
                  "2210002",
                  "2210003"
              ],
              "TeacherId": "67307b348e77d2a6ec61b4ee",
              "CreatedBy": "6730d78984648a9d5ba4d2ec",
              "UpdatedBy": "6730d78984648a9d5ba4d2ec"
          }
      ]
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy lớp học"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi hệ thống"
  }

##### d. GetClassByClassID
- **Mô tả:** Được sử dụng để lấy thông tin của lớp học theo id.
- **URL:** `/admin/api/class/{id}`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                             |
| ---- | ------------ | -------- | --------------------------------- |
| `id` | String       | Có       | ID của lớp học cần lấy thông tin. |
- **Response:**
- 200 OK:
  ```json
  {
      "message": "Lấy lớp học thành công",
      "status": "Success",
      "data": {
              "ID": "6744928c7ac9374bba79314d",
              "Semester": "HK241",
              "Name": "L02",
              "CourseId": "672b86fd226ae67ef9aaa045",
              "ListStudentMs": [
                  "2210001",
                  "2210002"
              ],
              "TeacherId": "672b85b8226ae67ef9aaa005",
              "CreatedBy": "000000000000000000000000",
              "UpdatedBy": "000000000000000000000000"
          }
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy lớp học"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi lấy dữ liệu từ cơ sở dữ liệu"
  }

##### e. DeleteClass
- **Mô tả:** Được sử dụng để xóa lớp học theo id.
- **URL:** `/admin/api/class/delete/{id}`
- **Phương thức:** `DELETE`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                   |
| ---- | ------------ | -------- | ----------------------- |
| `id` | String       | Có       | ID của lớp học cần xóa. |

- **Response:**
- 200 OK:
  ```json
  {
    "message": "Xóa lớp học thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi xóa lớp học"
  }

##### f. AddStudentToClass
- **Mô tả:** Được sử dụng để thêm sinh viên vào lớp học.
- **URL:** `/admin/api/class/add`
- **Phương thức:** `PATCH`
- **Xác thực**: Bearer `<token>`
-  **Request body:**
  ```json
  {
    "class_id": "67375642ec72767cfd44630c",
    "listStudent_ms": ["1242112", "1231311"]
  }
```
- 200 OK:
  ```json
  {
      "status": "Success",
      "message": "Thêm học sinh vào lớp học thành công"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi thêm học sinh vào lớp học"
  }

##### g. ChangeClassByID
- **Mô tả:** Được sử dụng để thay đổi dữ liệu của một lớp học.
- **URL:** `/admin/api/class/change/{id}`
- **Phương thức:** `PATCH`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                                  |
| ---- | ------------ | -------- | -------------------------------------- |
| `id` | String       | Có       | ID của lớp học cần thay đổi thông tin. |
-  **Request body:**
  ```json
  {
      "semester":       "HK241",
      "name":           "L15",
      "course_id":      "672b86fd226ae67ef9aaa045",
      "teacher_id":     "67307b348e77d2a6ec61b4ee"
  }
```
- **Response:**
- 200 OK:
  ```json
  {
        "message": "Cập nhật lớp học thành công",
        "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
      "status":    "Fail",
      "message": "Lỗi hệ thống"
  }

#### 1.4 Course
##### a. CreateCourse
- **Mô tả:** Được sử dụng để tạo khóa học
- **URL:** `/admin/api/course/create`
- **Phương thức:** `POST`
-  **Xác thực**: Bearer `<token>`
- **Request body:**
  ```json
  {
      "name": "Database system",
      "desc": "Mon nay de 10 diem",
      "credit": 4,
      "bt": 10,
      "tn": 10,
      "btl": 10,
      "gk": 20,
      "ck": 50,
      "ms": "CO3019"
  }
  
- **Response:**
- 200 OK:
  ```json
  {
      "message": "Tạo khóa học thành công",
      "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
      "status":    "Fail",
      "message": "Lỗi khi tạo khóa học"
  }

##### b. GetCourseByCourseID
- **Mô tả:** Được sử dụng để lấy thông tin của khóa học theo id.
- **URL:** `/admin/api/course/{id}`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                              |
| ---- | ------------ | -------- | ---------------------------------- |
| `id` | String       | Có       | ID của khóa học cần lấy thông tin. |
- **Response:**
- 200 OK: Phần hệ số có giá trị theo thứ tự là bt, tn, btl, gk, ck.
  ```json
  {
      "message": "Lấy khóa học thành công",
      "status": "Success",
      "data": {
              "ID": "674ac0549d65503d01c2b59e",
            "MS": "CO3019",
            "Credit": 4,
            "Name": "DBS",
            "Desc": "Mon nay de 10 diem",
            "HS": [
                10, 
                10, 
                10,
                20,
                50
            ],
            "CreatedBy": "6749bde782c0c465f12d5d11",
            "UpdatedBy": "000000000000000000000000"
          }
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy khóa học"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi lấy khóa học"
  }

##### c. GetAllCourse
- **Mô tả:** Được sử dụng để lấy thông tin của tất cả khóa học 
- **URL:** `/admin/api/course/all`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK: Phần hệ số có giá trị theo thứ tự là bt, tn, btl, gk, ck.
  ```json
  {
    "message": "Lấy ra tất cả khóa học thành công",
      "status": "Success",
      "semester": {
          "CURRENT": "HK241",
          "NEXT": "HK242",
        "PREV": "HK233"
      },
      "data": [
            {
                "ID": "672b86fd226ae67ef9aaa045",
                "MS": "CO3103",
                "Credit": 1,
                "Name": "Đồ án công nghệ phần mềm",
                "Desc": "Giúp bạn có 1 kiến thức thú dị :)))",
                "HS": [
                    10,
                    20,
                    0,
                    20,
                    50
                ],
                "CreatedBy": "670de16d6fbedac9e3b8d00a",
                "UpdatedBy": "000000000000000000000000"
            },
            {
                "ID": "6730cf4ee13cf54ebe00f3cf",
                "MS": "CO3102",
                "Credit": 4,
                "Name": "cấu trúc dữ liệu và ảo thuật",
                "Desc": "DSA dễ ẹc à",
                "HS": [
                    10,
                    20,
                    0,
                    20,
                    50
                ],
                "CreatedBy": "672b865f226ae67ef9aaa044",
                "UpdatedBy": "000000000000000000000000"
            }
          ]
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi lấy dữ liệu"
  }

##### d. DeleteCourseByID
- **Mô tả:** Được sử dụng để xóa khóa học theo id
- **URL:** `/admin/api/course/delete/{id}`
- **Phương thức:** `DELETE`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                    |
| ---- | ------------ | -------- | ------------------------ |
| `id` | String       | Có       | ID của khóa học cần xóa. |

- **Response:**
- 200 OK:
  ```json
  {
    "message": "Xóa khóa học thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi xóa khóa học"
  }

##### e. ChangeCourseByID
- **Mô tả:** Được sử dụng để thay đổi thông tin khóa học theo id
- **URL:** `/admin/api/course/change/{id}`
- **Phương thức:** `PATCH`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                                   |
| ---- | ------------ | -------- | --------------------------------------- |
| `id` | String       | Có       | ID của khóa học cần thay đổi thông tin. |
-  **Request body:**
  ```json
  {
      "ms": "CO3001",
      "credit": 3,
      "name": "Cong nghe phan mem",
      "desc": "Mon nay lay 10 diem"
  }
```
- **Response:**
- 200 OK:
  ```json
  {
      "message": "Cập nhật khóa học thành công",
      "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 500 Internal Server Error:
  ```json
  {
      "status":    "Fail",
      "message": "Lỗi khi cập nhật khóa học"
  }

#### 1.5 Result
##### a. CreateResult
- **Mô tả:** Được sử dụng để cập nhật hoặc tạo bảng điểm cho lớp học.
- **URL:** `/admin/api/result/create`
- **Phương thức:** `POST`
- **Xác thực**: Bearer `<token>`
- **Request body:**
  ```json
  {
      "score": [
          {
              "MMSV": "2215555",
              "Data": {
                  "BT": [8.5, 9.0, 7.5],
                  "TN": [7.0, 6.5, 8.0],
                  "BTL": [8.0, 7.5],
                  "GK": 7.5,
                  "CK": 8.0
              }
          },
          {
              "MMSV": "2212372",
              "Data": {
                  "BT": [9.0, 9.5, 8.5],
                  "TN": [8.0, 7.5, 9.0],
                  "BTL": [8.5, 9.0],
                  "GK": 8.0,
                  "CK": 8.5
              }
          }
      ],
      "class_id": "6732e83c95f2243e7eeb497e"
  }

- **Response:**
- 200 OK:
  ```json
  {
      "message": "Cập nhật bảng điểm thành công",
      "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy bảng điểm"
  }

- 500 Internal Server Error:
  ```json
  {
      "status":    "Fail",
      "message": "Cập nhật bảng điểm thât bại"
  }

##### b. GetResultByClassID
- **Mô tả:** Được sử dụng để lấy thông tin điểm theo id lớp học
- **URL:** `/admin/api/result/{id}`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                             |
| ---- | ------------ | -------- | --------------------------------- |
| `id` | String       | Có       | ID của lớp học cần lấy thông tin. |

- **Response:**
- 200 OK:
  ```json
  {
    "message": "Lấy bảng điểm thành công.",
    "status": "Success",
    "data": {
        "ID": "6730da8cb8b2e2dee4666a1c",
        "Semester": "HK233",
        "SCORE": [
            {
                "MSSV": "2210000",
                "Data": {
                    "BT": null,
                    "TN": null,
                    "BTL": null,
                    "GK": 8,
                    "CK": 0
                }
            },
            {
                "MSSV": "2210001",
                "Data": {
                    "BT": null,
                    "TN": null,
                    "BTL": null,
                    "GK": 10,
                    "CK": 10
                }
            }
        ],
        "ClassID": "672b87af226ae67ef9aaa047",
        "CourseID": "672b86fd226ae67ef9aaa045",
        "ExpiredAt": "2025-05-10T16:08:44.496Z",
        "CreatedBy": "672b865f226ae67ef9aaa044",
        "UpdatedBy": "672b865f226ae67ef9aaa044"
    }
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy bảng điểm"
  }

- 500 Internal Server Error:
  ```json
  {
      "status":    "Fail",
      "message": "Lỗi khi lấy bảng điểm"
  }

#### 1.6 HOF
##### a. UpdateHallOfFameForPrevSemester
- **Mô tả:** Được sử dụng để cập nhật Hall Of Fame cho kì vừa rồi.
- **URL:** `/admin/api/HOF/update`
- **Phương thức:** `POST`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "message": "Cập nhật Hall Of Fame thành công",
    "status": "Success",
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không có bản ghi nào"
  }

- 500 Internal Server Error:
  ```json
  {
      "status":    "Fail",
      "message": "Lỗi khi tính điểm trung bình"
  }

##### b. getAllHallOfFame
- **Mô tả:** Được sử dụng để đăng nhập cho admin
- **URL:** `/admin/api/HOF/all`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "message": "Lấy hall of fame thành công",
    "status": "Success",
    "data": {
        "semester": "HK233",
        "tier": [
            {
                "course_id": "6730cf4ee13cf54ebe00f3cf",
                "data": [
                    {
                        "mssv": "2111111",
                        "dtb": 8.383333
                    },
                    {
                        "mssv": "2222222",
                        "dtb": 7.7666664
                    }
                ]
            },
            {
                "course_id": "672b86fd226ae67ef9aaa045",
                "data": [
                    {
                        "mssv": "2210001",
                        "dtb": 7
                    },
                    {
                        "mssv": "2210000",
                        "dtb": 1.6
                    },
                    {
                        "mssv": "2210020",
                        "dtb": 1.5
                    }
                ]
            }
        ]
    }
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy dữ liệu cho học kỳ trước"
  }

- 500 Internal Server Error:
  ```json
  {
      "status":    "Fail",
      "message": "Đã xảy ra lỗi khi truy vấn dữ liệu"
  }

### 2. Client
#### 2.1 Student
#### 2.1.1 Index (Những hoạt động liên quan tới Student)
##### a. Login
- **Mô tả:** Được sử dụng để đăng nhập cho sinh viên
- **URL:** `/api/login`
- **Phương thức:** `POST`
- **Request body:**
  ```json
  {
    "idToken": "",
    "role": "student"
  }

- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "token": "",
    "role": "student"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Token không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Tài khoản không tồn tại"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi lấy dữ liệu từ server database."
  }

##### b. Logout
- **Mô tả:** Được sử dụng để đăng xuất cho sinh viên
- **URL:** `/api/logout`
- **Phương thức:** `POST`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "message": "Đăng xuất thành công"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Token không hợp lệ"
  }

##### c. Account
- **Mô tả:** Được sử dụng để lấy thông tin của sinh viên đang đăng nhập
- **URL:** `/api/info`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "message": "Thành công",
    "data":
		    "ID": "6749cfdaf2873c0c6c5fa7e0",
	        "Email": "kha.nguyentrong@hcmut.edu.vn",
	        "Name": "Nguyễn Trọng Kha",
	        "Ms": "2211417",
	        "Faculty": "KHMT",
	        "Role": "student",
	        "CreatedBy": "6749bde782c0c465f12d5d11",
	        "ExpiredAt": "2029-11-29T14:29:46.037Z"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Yêu cầu đăng nhập"
  }

##### d. GetTeacherByID
- **Mô tả:** Được sử dụng để lấy tên và email của giảng viên.
- **URL:** `/api/{id}`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                                |
| ---- | ------------ | -------- | ------------------------------------ |
| `id` | String       | Có       | ID của giảng viên cần lấy thông tin. |

- **Response:**
- 200 OK:
  ```json
  {
    "data": {
        "Name": "teacher1",
        "Email": "teacher1@hcmut.edu.vn"
    },
    "message": "Thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy giảng viên"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi truy vấn dữ liệu"
  }

#### 2.1.2 Class
##### a. GetAllClassJoined
- **Mô tả:** Được sử dụng để lấy thông tin của tất cả lớp học mà sinh viên đăng đăng nhập tham gia
- **URL:** `/api/class/account`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
-  **Response:**
- 200 OK:
  ```json
  {
    "data": [
        {
            "ID": "674abe469d65503d01c2b59d",
            "Semester": "HK242",
            "Name": "L01",
            "CourseId": "672b86fd226ae67ef9aaa045",
            "TeacherId": "672b85b8226ae67ef9aaa006"
        },
        {
            "ID": "623abe469d65503d01c2b50d",
            "Semester": "HK242",
            "Name": "L02",
            "CourseId": "672b86fd226ae67ef9aaa045",
            "TeacherId": "672b85b8226ae67ef9aaa006"
        }
    ],
    "status": "Success"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Sinh viên không tham gia lớp học nào"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi tìm lớp học"
  }

##### b. GetClassDetailByID
- **Mô tả:** Được sử dụng để lấy thông tin của lớp học dựa trên id lớp học
- **URL:** `/api/class/{id}`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                                |
| ---- | ------------ | -------- | ------------------------------------ |
| `id` | String       | Có       | ID của giảng viên cần lấy thông tin. |

- **Response:**
- 200 OK:
  ```json
  {
    "data": {
        "ID": "674abe469d65503d01c2b59d",
        "Semester": "HK242",
        "Name": "L01",
        "CourseId": "672b86fd226ae67ef9aaa045",
        "ListStudentMs": [
            "2222222",
            "2210011",
            "2210012"
        ],
        "TeacherId": "672b85b8226ae67ef9aaa006",
        "CreatedBy": "6749bde782c0c465f12d5d11",
        "UpdatedBy": "6749bde782c0c465f12d5d11"
    },
    "message": "Lấy lớp học thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Chỉ sinh viên mới được phép truy cập"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy lớp học"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi truy vấn dữ liệu"
  }

#### 2.1.3 Course
##### a. GetCourseByID
- **Mô tả:** Được sử dụng để lấy thông tin khóa học
- **URL:** `/api/course/{id}`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                              |
| ---- | ------------ | -------- | ---------------------------------- |
| `id` | String       | Có       | ID của khóa học cần lấy thông tin. |
- **Response:**
- 200 OK:
  ```json
  {
    "data": {
        "ID": "672b86fd226ae67ef9aaa045",
        "MS": "CO3103",
        "Credit": 1,
        "Name": "Đồ án công nghệ phần mềm",
        "Desc": "Giúp bạn có 1 kiến thức thú dị :)))",
        "HS": [
            10,
            20,
            0,
            20,
            50
        ],
        "CreatedBy": "670de16d6fbedac9e3b8d00a",
        "UpdatedBy": "000000000000000000000000"
    },
    "message": "Lấy khóa học thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy khóa học"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi lấy khóa học"
  }

#### 2.1.4 result
##### a. GetCourseResult
- **Mô tả:** Được sử dụng để lấy thông tin điểm của môn học tại một kì xác định
- **URL:** `/api/result/getmark/{mã môn học-kì}`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên             | Kiểu dữ liệu | Bắt buộc | Mô tả                                |
| --------------- | ------------ | -------- | ------------------------------------ |
| `mã môn học-kì` | String       | Có       | mã môn học-học kì cần lấy thông tin. |

- **Response:**
- 200 OK:
  ```json
  {
    "data": {
		"BT": [9, 10],
		"TN": [8, 9],
		"BTL": [10],
		"GK": 10,
		"CK": 10
    },
    "message": "Lấy điểm thành công",
    "status": "Success"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy bảng điểm"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi truy vấn dữ liệu"
  }

##### b. GetAllResult
- **Mô tả:** Được sử dụng để lấy tất cả bảng điểm của tất cả môn học
- **URL:** `/api/result/getmark`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "data": [
	    {
		    "ms": "CO3001-HK241",
		    "data":{
				"BT": [9, 10],
				"TN": [8, 9],
				"BTL": [10],
				"GK": 10,
				"CK": 10
		    }
	    },
	    {
		    "ms": "CO3005-HK241",
		    "data":{
				"BT": [9, 10],
				"TN": [8, 9],
				"BTL": [10],
				"GK": 10,
				"CK": 10
		    }
	    }
    ],
    "message": "Lấy điểm thành công",
    "status": "Success"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi truy vấn dữ liệu"
  }

#### 2.1.5 HOF
##### a. GetAllHallOfFame
- **Mô tả:** Được sử dụng để lấy thông tin của hall of fame
- **URL:** `/api/HOF/all`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "message": "Lấy hall of fame thành công",
    "status": "Success",
    "data": {
        "semester": "HK233",
        "tier": [
            {
                "course_id": "6730cf4ee13cf54ebe00f3cf",
                "data": [
                    {
                        "mssv": "2111111",
                        "dtb": 8.383333
                    },
                    {
                        "mssv": "2222222",
                        "dtb": 7.7666664
                    }
                ]
            },
            {
                "course_id": "672b86fd226ae67ef9aaa045",
                "data": [
                    {
                        "mssv": "2210001",
                        "dtb": 7
                    },
                    {
                        "mssv": "2210000",
                        "dtb": 1.6
                    },
                    {
                        "mssv": "2210020",
                        "dtb": 1.5
                    }
                ]
            }
        ]
    }
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy dữ liệu cho học kỳ trước"
  }

- 500 Internal Server Error:
  ```json
  {
      "status":    "Fail",
      "message": "Đã xảy ra lỗi khi truy vấn dữ liệu"
  }

#### 2.2 Teacher
#### 2.2.1 Index (thông tin của teacher)
##### a. Login
- **Mô tả:** Được sử dụng để đăng nhập cho giảng viên
- **URL:** `/api/login`
- **Phương thức:** `POST`
- **Request body:**
  ```json
  {
    "idToken": "",
    "role": "teacher"
  }

- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "token": "",
    "role": "teacher"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Token không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Tài khoản không tồn tại"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi lấy dữ liệu từ server database."
  }

##### b. Logout
- **URL:** `/api/logout`
- **Mô tả:** Được sử dụng để đăng xuất cho giảng viên
- **Phương thức:** `POST`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "message": "Đăng xuất thành công"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Token không hợp lệ"
  }

##### c. GetAccountInfo
- **Mô tả:** Được sử dụng để lấy thông tin của giảng viên đang đăng nhập
- **URL:** `/api/info`
- **Phương thức:** `POST`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "status": "Success",
    "message": "Thành công",
    "data":
		    "ID": "674ac897440b66d1866e3386",
	        "Email": "kha.nguyentrong@hcmut.edu.vn",
	        "Name": "kha",
	        "Ms": "2211418",
	        "Faculty": "KHMT",
	        "Role": "teacher",
	        "CreatedBy": "6749bde782c0c465f12d5d11",
	        "ExpiredAt": "2029-11-30T08:11:03.907Z"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Yêu cầu đăng nhập"
  }

#### 2.2.2 Class 
##### a. GetAllClassManaged
- **Mô tả:** Được sử dụng để lấy thông tin của lớp học mà giảng viên đang quản lý. 
- **URL:** `/api/class/account`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
-  **Response:**
- 200 OK:
  ```json
  {
    "data": [
        {
            "ID": "6732e83c95f2243e7eeb497e",
            "Semester": "HK233",
            "Name": "L01",
            "CourseId": "672b86fd226ae67ef9aaa045",
            "ListStudentMs": [
                "2210001",
                "2210002"
            ],
            "TeacherId": "674ac897440b66d1866e3386",
            "CreatedBy": "672b865f226ae67ef9aaa044",
            "UpdatedBy": "6749bde782c0c465f12d5d11"
        },
        {
            "ID": "6732e83c95f2243ere1eb497e",
            "Semester": "HK233",
            "Name": "l02",
            "CourseId": "672b86fd226ae67ef9aaa045",
            "ListStudentMs": [
                "2210001",
                "2210002",
                "2211411"
            ],
            "TeacherId": "674ac897440b66d1866e3386",
            "CreatedBy": "672b865f226ae67ef9aaa044",
            "UpdatedBy": "6749bde782c0c465f12d5d11"
        }
    ],
    "status": "Success"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Chỉ giáo viên mới được phép truy cập"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Giảng viên không quản lý lớp học nào"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi tìm lớp học"
  }

##### b. GetClassDetailByID
- **Mô tả:** Được sử dụng để lấy thông tin của class dựa trên id dành cho giảng viên
- **URL:** `/api/class/{id}`
- **Phương thức:** `POST`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                             |
| ---- | ------------ | -------- | --------------------------------- |
| `id` | String       | Có       | ID của lớp học cần lấy thông tin. |

- **Response:**
- 200 OK:
  ```json
  {
    "data": {
        "ID": "6732e83c95f2243e7eeb497e",
        "Semester": "",
        "Name": "",
        "CourseId": "672b86fd226ae67ef9aaa045",
        "ListStudentMs": [
            "2210001",
            "2210002"
        ],
        "TeacherId": "674ac897440b66d1866e3386",
        "CreatedBy": "672b865f226ae67ef9aaa044",
        "UpdatedBy": "6749bde782c0c465f12d5d11"
    },
    "message": "Lấy lớp học thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Chỉ giáo viên mới được phép truy cập"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy lớp học"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi truy vấn dữ liệu"
  }

#### 2.2.3 course
##### a. GetCourseByID
- **Mô tả:** Được sử dụng để lấy thông tin của course dựa trên id
- **URL:** `/api/course/{id}`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                              |
| ---- | ------------ | -------- | ---------------------------------- |
| `id` | String       | Có       | ID của khóa học cần lấy thông tin. |
- **Response:**
- 200 OK:
  ```json
  {
    "data": {
        "ID": "672b86fd226ae67ef9aaa045",
        "MS": "CO3103",
        "Credit": 1,
        "Name": "Đồ án công nghệ phần mềm",
        "Desc": "Giúp bạn có 1 kiến thức thú dị :)))",
        "HS": [
            10,
            20,
            0,
            20,
            50
        ],
        "CreatedBy": "670de16d6fbedac9e3b8d00a",
        "UpdatedBy": "000000000000000000000000"
    },
    "message": "Lấy khóa học thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy khóa học"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi lấy khóa học"
  }

#### 2.2.4 result
##### a. CreateResultForClassByClassId 
- **Mô tả:** Được sử dụng để tạo bảng điểm cho lớp học theo id lớp học.
- **URL:** `/api/result/create`
- **Phương thức:** `POST`
- **Request body:**
  ```json
  {
    "score": [
        {
            "MMSV": "2210001",
            "Data": {
                "BT": [8.5, 9.0, 7.5],
                "TN": [7.0, 6.5, 8.0],
                "BTL": [8.0, 7.5],
                "GK": 7.5,
                "CK": 8.0
            }
        },
        {
            "MMSV": "2210002",
            "Data": {
                "BT": [9.0, 9.5, 8.5],
                "TN": [8.0, 7.5, 9.0],
                "BTL": [8.5, 9.0],
                "GK": 8.0,
                "CK": 8.5
            }
        }
    ],
    "class_id": "6732e83c95f2243e7eeb497e"
  }

- **Response:**
- 200 OK:
  ```json
  {
    "message": "Tạo bảng điểm thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Chỉ giáo viên mới được phép truy cập"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy lớp học"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi truy vấn dữ liệu"
  }

##### b. GetResultByClassId
- **Mô tả:** Được sử dụng để lấy thông tin điểm của tất cả sinh viên trong lớp học
- **URL:** `/api/result/{id}`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- Path Parameters:

| Tên  | Kiểu dữ liệu | Bắt buộc | Mô tả                             |
| ---- | ------------ | -------- | --------------------------------- |
| `id` | String       | Có       | ID của lớp học cần lấy thông tin. |
- **Response:**
- 200 OK:
  ```json
  {
    "data": {
        "ID": "674ad1556527ef3caf7c37af",
        "Semester": "HK233",
        "SCORE": [
            {
                "MSSV": "2210001",
                "Data": {
                    "BT": [
                        8.5,
                        9,
                        9
                    ],
                    "TN": [
                        7,
                        6.5,
                        8
                    ],
                    "BTL": [
                        8,
                        7.5
                    ],
                    "GK": 7.5,
                    "CK": 8
                }
            },
            {
                "MSSV": "2210002",
                "Data": {
                    "BT": [
                        9,
                        9.5,
                        8.5
                    ],
                    "TN": [
                        8,
                        7.5,
                        9
                    ],
                    "BTL": [
                        8.5,
                        9
                    ],
                    "GK": 8,
                    "CK": 8.5
                }
            }
        ],
        "ClassID": "6732e83c95f2243e7eeb497e",
        "CourseID": "672b86fd226ae67ef9aaa045",
        "ExpiredAt": "2025-05-30T08:48:21.478Z",
        "CreatedBy": "674ac897440b66d1866e3386",
        "UpdatedBy": "674ac897440b66d1866e3386"
    },
    "status": "Success"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy bảng điểm"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Chỉ giảng viên mới được phép truy cập"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi truy vấn dữ liệu"
  }

##### c. ChangeResultForClassByClassId
- **Mô tả:** Được sử dụng để cập nhật điểm cho một lớp học
- **URL:** `/api/result/change`
- **Phương thức:** `PATCH`
- **Xác thực**: Bearer `<token>`
- **Request body:**
  ```json
  {
    "score": [
        {
            "MMSV": "2210001",
            "Data": {
                "BT": [8.5, 9.0, 9],
                "TN": [7.0, 6.5, 8.0],
                "BTL": [8.0, 7.5],
                "GK": 7.5,
                "CK": 8.0
            }
        },
        {
            "MMSV": "2210002",
            "Data": {
                "BT": [9.0, 9.5, 8.5],
                "TN": [8.0, 7.5, 9.0],
                "BTL": [8.5, 9.0],
                "GK": 8.0,
                "CK": 8.5
            }
        }
    ],
    "class_id": "6732e83c95f2243e7eeb497e"
  }

- **Response:**
- 200 OK:
  ```json
  {
    "message": "Thay đổi thành công",
    "status": "Success"
  }

- 400 Bad Request:
  ```json
  {
    "status": "Fail",
    "message": "Dữ liệu yêu cầu không hợp lệ"
  }

- 401 Unauthorized:
  ```json
  {
    "status": "Fail",
    "message": "Chỉ giảng viên mới được phép truy cập"
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy khóa học"
  }

- 500 Internal Server Error:
  ```json
  {
    "status": "Fail",
    "message": "Lỗi khi lấy khóa học"
  }

#### 2.2.5 HOF
##### a. GetAllHallOfFame
- **Mô tả:** Được sử dụng để lấy hall of fame
- **URL:** `/api/HOF/all`
- **Phương thức:** `GET`
- **Xác thực**: Bearer `<token>`
- **Response:**
- 200 OK:
  ```json
  {
    "message": "Lấy hall of fame thành công",
    "status": "Success",
    "data": {
        "semester": "HK233",
        "tier": [
            {
                "course_id": "6730cf4ee13cf54ebe00f3cf",
                "data": [
                    {
                        "mssv": "2111111",
                        "dtb": 8.383333
                    },
                    {
                        "mssv": "2222222",
                        "dtb": 7.7666664
                    }
                ]
            },
            {
                "course_id": "672b86fd226ae67ef9aaa045",
                "data": [
                    {
                        "mssv": "2210001",
                        "dtb": 7
                    },
                    {
                        "mssv": "2210000",
                        "dtb": 1.6
                    },
                    {
                        "mssv": "2210020",
                        "dtb": 1.5
                    }
                ]
            }
        ]
    }
  }

- 404 Not Found:
  ```json
  {
    "status": "Fail",
    "message": "Không tìm thấy dữ liệu cho học kỳ trước"
  }

- 500 Internal Server Error:
  ```json
  {
      "status":    "Fail",
      "message": "Đã xảy ra lỗi khi truy vấn dữ liệu"
  }
