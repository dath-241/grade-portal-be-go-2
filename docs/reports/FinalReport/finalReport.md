<div align="center">

<img src="img/hcmutlogo.png" alt="alt text" width="150"/>

## ĐỒ ÁN TỔNG HỢP: HƯỚNG CÔNG NGHỆ PHẦN MỀM  

# BÁO CÁO ĐỒ ÁN  

## CHỦ ĐỀ: GRADE PORTAL BACKEND GO 2  

| **STT** | **Họ và Tên**          | **MSSV**  | **Vai Trò**            |
|---------|-------------------------|-----------|------------------------|
| 1       | Chu Minh Tâm            | 2213009   | Product Owner          |
| 2       | Nguyễn Trọng Kha        | 2211417   | Developer              |
| 3       | Nguyễn Tiến Phát        | 2212527   | Developer              |
| 4       | Nguyễn Ngọc Diệu Hà     | 2210846   | Developer              |
| 5       | Lê Hoàng Ngọc Hân       | 2210935   | Developer              |
| 6       | Nguyễn Phúc Hưng        | 2211368   | Developer              |
| 7       | Huỳnh Thanh Tâm         | 2213012   | Developer              |

</div>

---


# Functional & Non-Functional Requirements

## Functional Requirements

- **Đối với Admin**
  - Đăng nhập vào/đăng xuất khỏi hệ thống.
  - Tạo tài khoản admin mới, cấp quyền quản trị cho các tài khoản khác khi cần thiết.
  - Xem thông tin hồ sơ của mình, bao gồm các thông tin cá nhân và quyền hạn trong hệ thống.
  - Tạo, cập nhật, xóa và xem thông tin chi tiết của tài khoản giảng viên/sinh viên.
  - Xem danh sách tài khoản của tất cả giảng viên/sinh viên.
  - Tạo, cập nhật, xóa, và xem thông tin chi tiết của lớp học/khoá học.
  - Xem danh sách các khoá học/lớp học của giảng viên giảng dạy/sinh viên theo học.
  - Thêm sinh viên vào lớp học.
  - Cập nhật danh sách Hall of Fame cũng như xem danh sách Hall of Fame của các học kì.
  - Tạo và xem kết quả điểm số của các lớp học.

- **Đối với Giảng viên**
  - Đăng nhập vào/đăng xuất khỏi hệ thống.
  - Xem thông tin hồ sơ của mình, bao gồm các thông tin cá nhân và quyền hạn trong hệ thống.
  - Xem thông tin chi tiết lớp học/khoá học mà giảng viên giảng dạy.
  - Xem danh sách Hall of Fame của khoá học mà giảng viên giảng dạy của các học kì.
  - Tạo, cập nhật và xem kết quả điểm số của cả lớp học hoặc một sinh viên cụ thể mà giảng viên giảng dạy.

- **Đối với Sinh viên**
  - Đăng nhập vào/đăng xuất khỏi hệ thống.
  - Xem thông tin hồ sơ của mình, bao gồm các thông tin cá nhân và quyền hạn trong hệ thống.
  - Xem thông tin chi tiết lớp học/khoá học mà sinh viên theo học.
  - Xem danh sách Hall of Fame các khoá học mà sinh viên theo học của các học kì.
  - Chỉ được xem kết quả điểm số của chính mình.

## Non-Functional Requirements

- **Performance (Hiệu năng)**
  - Hệ thống phải phản hồi nhanh chóng trong điều kiện tải thông thường.
  - Hệ thống phải xử lý được nhiều yêu cầu đồng thời từ người dùng mà không bị suy giảm hiệu năng.

- **Security (Bảo mật)**
  - Áp dụng xác thực dựa trên token cho tất cả các route được bảo vệ.
  - Phân quyền rõ ràng: Admin chỉ được truy cập route `/admin`, giảng viên và sinh viên bị giới hạn trong route `/client`.
  - Sử dụng HTTPS để đảm bảo giao tiếp an toàn.

- **Scalability (Khả năng mở rộng)**
  - Hệ thống phải hỗ trợ mở rộng theo chiều ngang để xử lý lưu lượng tăng cao.
  - Cơ sở dữ liệu phải xử lý hiệu quả khi số lượng bản ghi (tài khoản, lớp học, kết quả) tăng lên.

- **Reliability (Độ tin cậy)**
  - Hệ thống cần đảm bảo 99.9% thời gian hoạt động với cơ chế xử lý lỗi mạnh mẽ.
  - Dữ liệu phải được được backup và khôi phục nhanh chóng trong trường hợp xảy ra sự cố.

- **Maintainability (Tính dễ bảo trì)**
  - Cấu trúc module rõ ràng giúp dễ dàng bảo trì và cập nhật.
  - Tuân thủ tiêu chuẩn coding và tài liệu hóa đầy đủ cho các API.

- **Usability (Tính thân thiện Người dùng)**
  - Cung cấp thông báo lỗi rõ ràng khi thao tác thất bại (ví dụ: đăng nhập không thành công, truy cập bị từ chối).
  - Đảm bảo các API RESTful với tên gọi nhất quán để dễ tích hợp.

- **Compatibility (Tính tương thích)**
  - Đảm bảo tương thích với các framework frontend phổ biến (ReactJS, Angular, v.v.).
  - API hỗ trợ định dạng JSON cho các request và response payload.

---

# Sơ đồ thành phần

![alt text](img/component.png)


### Các thành phần chính
- View: Giao diện sinh viên, giảng viên và admin, lớp học, khóa học và danh sách sinh viên có điểm số cao nhất. Riêng admin có giao diện tạo mới, sửa, xóa tài khoản, lớp học, khóa học. Khi người dùng gửi yêu cầu từ giao diện sẽ gọi đến dịch vụ ở lớp dưới.
- Controller: Tiếp nhận và xử lý các yêu cầu từ component View, sau đó chuyển yêu cầu tới các dịch vụ ở component tiếp theo.
- Model: Thực hiện các dịch vụ được yêu cầu và tương tác với kho dữ liệu để đáp ứng dịch vụ đó.
- Repository: Cung cấp dữ liệu cho các dịch vụ và thực hiện các thao tác trên cơ sở dữ liệu.
Database: Chứa các dữ liệu "tài khoản", "lớp học", "khóa học", "môn học", "danh sách sinh viên cao điểm nhất", các dữ liệu này phụ thuộc vào thông tin từ người dùng.
---

### Thiết kế CSDL
![alt text](img/erd.jpg)

1. Trong layered architecture, database sẽ là lớp ở dưới cùng và sẽ chịu trách nhiệm lưu trữ toàn bộ data và xử lý chúng. Các dữ liệu của ứng dụng sẽ  được lưu trữ tại đây và các thao tác như search, insert, update and delete sẽ được thực hiện thường xuyên để thao tác với dữ liệu thông qua hệ quản trị cơ sở dữ liệu.
2. Đối với đồ án lần này, nhóm sẽ sử dụng kiến trúc lớp với database layer lưu trữ dữ liệu bằng MongoDB, một hệ cơ sở dữ liệu NoSQL, các kiểu thực thể cần thiết sẽ có các thuộc tính như trên hình.
