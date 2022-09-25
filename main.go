package main

import (
	router "TODO/api/Handlers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title TODO APIs
// @version 1.0
// @description This is a sample server server.
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
// @BasePath /
// @schemes http
func main() {

	r := gin.Default()
	r = router.Router()

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	fmt.Println("this is port", port)

	fmt.Print("Listening on port")
	r.Run(":" + port)
}
