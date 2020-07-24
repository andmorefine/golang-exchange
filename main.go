package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/andmorefine/golang-exchenge/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// プロファイルデータの取得
	f, _ := os.Create("test.profile")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	fmt.Println("success")
	r := gin.Default()
	user := models.User{UUID: 1, Name: "moge"}
	user.Create()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.GET("/health_check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test":  "item",
			"array": []int{1, 2, 3, 4, 5},
			"map":   map[string]int{"apple": 150, "banana": 300, "lemon": 300},
		})
	})
	r.Run(":8080")
}
