package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("success")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.GET("/mogemoge", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "mogetta",
		})
	})
	r.Run(":8080")
}
