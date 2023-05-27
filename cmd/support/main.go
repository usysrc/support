package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define your application routes here
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start the server
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
