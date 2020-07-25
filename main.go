package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"runtime/pprof"

	"github.com/andmorefine/golang-exchange/config"
	"github.com/andmorefine/golang-exchange/models"
	"github.com/andmorefine/golang-exchange/utils"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	// プロファイルデータの取得
	f, _ := os.Create("test.profile")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// db 接続
	db, err := sql.Open("mysql", "root:password@tcp(db_mysql5.7:3306)/my_database?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ログ生成
	utils.LoggingSettings(config.Config.LogFile)

	// success
	log.Println("success")

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
	r.GET("/create", func(c *gin.Context) {
		user := models.User{UUID: 1, Name: "moge", DeleteFlag: false}
		result, _ := user.Create()
		LastInsertId, _ := result.LastInsertId()
		log.Println("LastInsertId: ", LastInsertId)
		c.JSON(200, gin.H{
			"LastInsertId": LastInsertId,
			"test":         "item",
			"array":        []int{1, 2, 3, 4, 5},
			"map":          map[string]int{"apple": 150, "banana": 300, "lemon": 300},
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	r.POST("/user/:name/*action", func(c *gin.Context) {
		c.FullPath()
	})

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.JSON(200, gin.H{
			"status":    http.StatusOK,
			"matched":   "get",
			"firstname": firstname,
			"lastname":  lastname,
		})

		// c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.Run(":8080")
}
