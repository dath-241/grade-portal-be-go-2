package main

import (
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
