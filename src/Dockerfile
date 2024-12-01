FROM golang:1.23.0 AS builder
WORKDIR /app
COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -tags netgo -ldflags '-s -w' -o main .
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
# B1 : cd vào thư mục chứa Dockerfile
# B2 : Build image -> docker build -t go2 .
# B3 : Run container -> docker run -p 8080:8080 go2
# B4 : Kiểm tra container đang chạy -> docker ps