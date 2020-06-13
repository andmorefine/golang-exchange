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
			"message": "ok",
			"age":     40,
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "item",
		})
	})
	r.Run(":8080")
}
