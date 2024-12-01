package main

import (
<<<<<<< HEAD
	"be-go-2/config"
	routes_admin "be-go-2/routes/admin"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	godotenv.Load()
	config.ConnectMongoDB(os.Getenv("MONGODB_URL"))

	app := gin.Default()

	routes_admin.MainRoute(app)

	fmt.Println("Server is running on port :", os.Getenv("PORT"))
	app.Run(":" + os.Getenv("PORT"))
}
=======
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

func main() {
    gin.SetMode(gin.ReleaseMode)

    // Load các biến môi trường từ file .env
    godotenv.Load()
    config.ConnectMongoDB(os.Getenv("MONGO_URL"))


    app := gin.Default()

    // Cấu hình CORS để hạn chế các request từ các domain khác
    app.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080","http://localhost:5500"},
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
>>>>>>> 6bfebbc4c1f06c7bc871229c29617513338352cd
