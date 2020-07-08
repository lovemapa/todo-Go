package main

import (
	router "TODO/api/Handlers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()
	r = router.Router()

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	fmt.Println("this is port", port)

	fmt.Print("Listening on port")
	r.Run(":" + port)
}
