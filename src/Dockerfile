# Sử dụng image Go chính thức
FROM golang:1.20 as builder

# Đặt thư mục làm việc trong container
WORKDIR /app

# Copy go.mod và go.sum để cài đặt dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ mã nguồn vào container
COPY . .

# Biên dịch ứng dụng
RUN go build -o main ./src/main.go

# Tạo một image nhỏ hơn để chạy ứng dụng từ kết quả biên dịch
FROM debian:bullseye-slim

WORKDIR /app

# Copy file thực thi từ bước builder
COPY --from=builder /app/main /app/main

# Copy các file cấu hình và tài nguyên cần thiết (nếu có)
COPY --from=builder /app/.env /app/.env

# Thiết lập biến môi trường (nếu cần)
ENV GIN_MODE=release

# Expose cổng mà ứng dụng sẽ chạy
EXPOSE 8080

# Chạy ứng dụng
CMD ["/app/main"]
