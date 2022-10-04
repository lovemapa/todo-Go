package main

import (
	router "TODO/api/Handlers"
	"fmt"
	"os"

	_ "TODO/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title TODO APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @securityDefinitions.apiKey JWT
// @in header
// @name token

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath  /api/v1

// @schemes http
func main() {

	r := gin.Default()
	r = router.Router()

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	fmt.Println("this is port", port)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	fmt.Print("Listening on port")
	r.Run(":" + port)
}
