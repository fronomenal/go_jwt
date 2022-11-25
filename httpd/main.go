package main

import (
	"github.com/fronomenal/go_jwt/httpd/inits"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	inits.SetupEnv()
	DB = inits.Connect()
	inits.Sync(DB)
}

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "200", "message": "success",
		})
	})
	r.Run()
}
