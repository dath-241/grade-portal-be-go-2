package main

import (
    "Go2/config"
    routes_admin "Go2/routes/admin"
    routes_client "Go2/routes/client"
    "fmt"
    "os"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://grade-portal-service.vercel.app")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Nếu là request OPTIONS (preflight), trả về 200 luôn
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}


func main() {
    gin.SetMode(gin.ReleaseMode)
    // Load các biến môi trường từ file .env
    godotenv.Load()
    config.ConnectMongoDB(os.Getenv("MONGO_URL"))


    app := gin.Default()

    // Cấu hình CORS để hạn chế các request từ các domain khác
    app.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080","http://localhost:5500","http://127.0.0.1:5500"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH", "HEAD", "CONNECT"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // Đăng ký các route
    routes_admin.MainRoute(app)
    routes_client.MainRoute(app)

    // Chạy server
    fmt.Println("Server đang chạy trên cổng", os.Getenv("PORT"))
    app.Run(":" + os.Getenv("PORT"))
}