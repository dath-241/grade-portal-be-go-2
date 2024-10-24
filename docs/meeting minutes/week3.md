# MEETING MINUTES

### Week : `3`

### LƯU Ý :

- Push code lên `branch` và merge code tại nhánh `develop`, không push code vào `main`.
- Ai đảm nhiệm phần nào phải viết `Docs` (`.md`), và comment vào source code để giải thích.
- Tạo 1 trang `.html `đơn giản để kiểm thử chức năng.
- `Thứ 3 (21/10)` sẽ có hàm mẫu để mọi người implement code.
- `Tối thứ 3 (21/10)` - `9h` : họp.

# NỘI DUNG

- Lên kế hoạch cho tuần làm việc tiếp theo `week 4`, (`21/10/2024`).

# CÔNG VIỆC

- Hoàn Thiện Documents: `Deadline : 24/10/2024`
  - ERD.
  - Database.
  - Diagram _(Usecase, Sequence)_.
  - RESTfuls API _(POST, GET, DELETE, PUT)_.
- Hoàn thiện chức năng Admin (_POST_): `Deadline : 27/10/2024`.

| STT | Tên                 | MSSV    | Công việc          |
| --- | ------------------- | ------- | ------------------ |
| 1   | Chu Minh Tâm        | 2213009 | APIs               |
| 2   | Nguyễn Tiến Phát    | 2212527 | ERD                |
| 3   | Nguyễn Ngọc Diệu Hà | 2210846 | ERD                |
| 4   | Nguyễn Trọng Kha    | 2211417 | Database           |
| 5   | Lê Hoàng Bảo Hân    | 2210935 | Database           |
| 6   | Nguyễn Phúc Hưng    | 2211368 | Diagram (Use-case) |
| 7   | Huỳnh Thanh Tâm     | 2213012 | Diagram (Sequence) |

# HOÀN THIỆN CHỨC NĂNG ADMIN

### Admin:

- `Đăng nhập`: Cookie, Session, OAuth2.
- `Middleware` :

  - `Authentication`
  - `Validate`

  | STT | Tên       | Công việc                  | Hàm tương tứng        |
  | --- | --------- | -------------------------- | --------------------- |
  | 1   | Hân       | Thêm `Teacher`             | `createTeacher`       |
  | 2   | Hà        | Thêm `Student`             | `createStudent`       |
  | 3   | Thanh Tâm | Thêm `Course`              | `createCourse`        |
  | 4   | Phát      | Thêm `Admin`               | `createAdmin`         |
  | 5   | Kha       | Thêm `Class`               | `createClass`         |
  | 6   | Minh Tâm  | Thêm `Student` vào `Class` | `addStudent_to_Class` |
  | 7   | Hưng      | Thêm `Grade`               | `createGrade`         |

  **Sau khi hoàn thành chức năng `POST` của `Admin` sẽ tiếp tục thực hiện các chức năng còn lại (PUT, GET, DELETE), sau đó sẽ đến `Student`, `Teacher`.**

---

# TÀI LIỆU

- [Golang](https://go.dev/)
- [REST APIs](https://viblo.asia/p/restful-api-la-gi-1Je5EDJ4lnL)
- [BE-GO-1](https://github.com/dath-241/grade-portal-be-go-1/tree/main)
