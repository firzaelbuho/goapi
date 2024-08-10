package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	username := os.Getenv("USERNAME")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": username,
		})
	})
	r.Run() // Default runs on :8080
}
