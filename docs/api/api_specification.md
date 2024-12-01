# TỔNG HỢP API HỆ THỐNG
```bash
routes
  ├── /admin
        ├── AccountRoute.go
        ├── AuthRoute.go
        ├── ClassRoute.go
        ├── CourseRoute.go
        ├── ResultRoute.go
        ├── HallOfFameRoute.go
        └── MainRoute.go # Tổng hợp 
  ├── /client
        ├── AccountRoute.go
        ├── ClassRoute.go
        ├── CourseRoute.go
        ├── ResultRoute.go
        ├── HomeRoute.go
        ├── HallOfFameRoute.go
        └── MainRoute.go # Tổng hợp 
```

### **AuthRoute**
- `POST /admin/api/login` - `HandleLogin()`: Xử lý đăng nhập.
- `POST /admin/api/logout` - `HandleRequireAuth()`, `HandleLogout()`: Xử lý đăng xuất, yêu cầu xác thực.
- `POST /admin/api/create` - `HandleRequireAuth()`, `ValidateDataAdmin()`, `HandleCreateAdmin()`: Tạo admin mới, yêu cầu xác thực, và kiểm tra dữ liệu.
- `GET /admin/api/profile` - `HandleRequireAuth()`, `HandleProfile()`: Lấy thông tin profile của admin, yêu cầu xác thực.

### **AccountRoute**
- `POST /admin/api/account/create` - `HandleCreateAccount()`: Tạo tài khoản mới.
- `GET /admin/api/account/:id` - `HandleGetAccountByID()`: Lấy thông tin tài khoản theo ID.
- `GET /admin/api/account/teacher` - `HandleGetTeacherAccounts()`: Lấy danh sách tài khoản giáo viên.
- `GET /admin/api/account/student` - `HandleGetStudentAccounts()`: Lấy danh sách tài khoản sinh viên.
- `DELETE /admin/api/account/delete/:id` - `HandleDeleteAccount()`: Xóa tài khoản theo ID.
- `PATCH /admin/api/account/change/:id` - `HandleUpdateAccount()`: Cập nhật thông tin tài khoản.

### **ClassRoute**
- `POST /admin/api/class/create` - `HandleCreateClass()`: Tạo lớp học mới.
- `GET /admin/api/class/:id` - `HandleGetClassByID()`: Lấy thông tin lớp học theo ID.
- `GET /admin/api/class/account/:id` - `HandleGetAllClassesByAccountID()`: Lấy danh sách lớp học của tài khoản.
- `GET /admin/api/class/course/:id` - `HandleGetClassesByCourseID()`: Lấy danh sách lớp học theo khóa học.
- `PATCH /admin/api/class/add` - `HandleAddStudentsToClass()`: Thêm sinh viên vào lớp học.
- `PATCH /admin/api/class/change/:id` - `HandleUpdateClass()`: Cập nhật thông tin lớp học.
- `DELETE /admin/api/class/delete/:id` - `HandleDeleteClass()`: Xóa lớp học theo ID.

### **CourseRoute**
- `POST /admin/api/course/create` - `HandleCreateCourse()`: Tạo khóa học mới.
- `GET /admin/api/course/:id` - `HandleGetCourseByID()`: Lấy thông tin khóa học theo ID.
- `GET /admin/api/course/all` - `HandleGetAllCourses()`: Lấy danh sách tất cả khóa học.
- `PATCH /admin/api/course/change/:id` - `HandleUpdateCourse()`: Cập nhật thông tin khóa học.
- `DELETE /admin/api/course/delete/:id` - `HandleDeleteCourse()`: Xóa khóa học theo ID.

### **HallOfFameRoute**
- `POST /admin/api/HOF/update` - `HandleCreateHallOfFame()`: Cập nhật bảng vinh danh.
- `GET /admin/api/HOF/all` - `HandleGetPrevSemesterAllHallOfFame()`: Lấy danh sách Hall of Fame của các học kỳ trước.

### **ResultRoute**
- `POST /admin/api/result/create` - `HandleCreateResult()`: Tạo kết quả điểm mới.
- `GET /admin/api/result/:id` - `HandleGetResult()`: Lấy thông tin kết quả điểm theo ID.

---

## **Client Routes**

### **HomeRouter**
- `GET /` - `HomePage()`: Trang chủ.

### **AccountRoute**
- `POST /api/login` - `HandleLogin()`: Đăng nhập.
- `POST /api/logout` - `RequireUser()`, `HandleLogout()`: Đăng xuất, yêu cầu người dùng xác thực.
- `GET /api/info` - `RequireUser()`, `HandleAccount()`: Lấy thông tin tài khoản người dùng, yêu cầu xác thực.
- `GET /api/:id` - `HandleGetInfoByID()`: Lấy thông tin tài khoản theo ID.

### **ClassRoute**
- `GET /api/class/account` - `HandleUserClasses()`: Lấy danh sách lớp học của người dùng.
- `GET /api/class/:id` - `HandleClassDetail()`: Lấy thông tin chi tiết lớp học theo ID.
- `GET /api/class/count/:id` - `HandleCountDocuments()`: Đếm số tài liệu trong lớp học.

### **CourseRoute**
- `GET /api/course/:id` - `HandleGetCourseByID()`: Lấy thông tin chi tiết khóa học theo ID.

### **HallOfFameRoute**
- `GET /api/HOF/all` - `GetAllPrevSemester()`: Lấy danh sách Hall of Fame của các học kỳ trước.

### **ResultRoute**
- `POST /api/result/create` - `RequireTeacher()`, `HandleCreateResult()`: Tạo kết quả điểm, yêu cầu giáo viên xác thực.
- `GET /api/result/getmark` - `HandleAllResults()`: Lấy tất cả kết quả điểm.
- `GET /api/result/getmark/:ms` - `HandleCourseResult()`: Lấy kết quả điểm của sinh viên theo mã số.
- `GET /api/result/:id` - `HandleResult()`: Lấy kết quả điểm theo ID.
- `PATCH /api/result/:id` - `RequireTeacher()`, `HandlePatchResult()`: Cập nhật kết quả điểm, yêu cầu giáo viên xác thực.
