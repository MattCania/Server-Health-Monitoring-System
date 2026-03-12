package main

import (
	"fmt"
	"go-server/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	routes.AuthRoute(r)
	port := os.Getenv("PORT")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API Monitor Running",
		})
	})

	fmt.Println("Running on port: " + port)
	r.Run(":" + port)
}
