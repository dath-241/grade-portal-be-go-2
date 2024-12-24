<a id="readme-top"></a>

# ĐỒ ÁN TỔNG HỢP HK241

### TÊN NHÓM :  ≽^• GO GO GO • ྀི≼
# Hướng dẫn sử dụng
## Cài đặt dự án
### Tải mã 
```bash
git clone https://github.com/dath-241/grade-portal-be-go-2.git
cd grade-portal-be-go-2.git
```
### Thiết lập biến môi trường
Bạn có thể tự tạo file .env trong folder và điền các đường dẫn tài nguyên cần sử dụng theo định dạng: 
```bash
MONGO_URL=<mongodb+srv>
PORT=8080
YOUR_CLIENT_ID=<Your client ID>
JWT_SECRET=<JWT Secret>
```
## Chạy dự án
### Không dùng docker
Nếu bạn không dùng docker bạn cần tải các dependencies sau:
```bash
go mod download
go install github.com/air-verse/air@latest
```
Sau khi chạy xong, bạn mở terminal của folder grade-portal-be-go-2 và thực hiện lệnh:
```bash
cd src
air
```
Bạn chờ hệ thống build và run đến khi hiện 2 dòng chữ sau là thành công:
```bash
2024/12/08 00:25:54 Connected to MongoDB
Server đang chạy trên cổng 8080
```
### Nếu bạn sử dụng docker và docker compose
Bạn có thể cài docker ([tại đây](https://docs.docker.com/get-docker/)) và docker compose ([tại đây](https://docs.docker.com/compose/install/))
Sau khi cấu hình môi trường, bạn có thể sử dụng lệnh sau để build và chạy dự án:
```bash
docker compose up --build -d
```
## Kiểm tra ứng dụng
Bạn có thể check health bằng cách vào trình duyệt và truy cập đường dẫn [localhost:8080](localhost:8080).
</br>
**Good Luck!**
