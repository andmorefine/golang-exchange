package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	// プロファイルデータの取得
	f, _ := os.Create("test.profile")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// db 接続
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3336)/my_database?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("success")
	r := gin.Default()
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
		// user := models.User{UUID: 1, Name: "moge", DeleteFlag: false}
		// user.Create()
		c.JSON(200, gin.H{
			"test":  "item",
			"array": []int{1, 2, 3, 4, 5},
			"map":   map[string]int{"apple": 150, "banana": 300, "lemon": 300},
		})
	})
	r.Run(":8080")
}
