package main

import (
	"github.com/fronomenal/go_jwt/httpd/controllers"
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
	r.GET("user/", controllers.UsersController("index", DB))
	r.POST("user/sign-up", controllers.UsersController("sign-up", DB))
	r.POST("user/login", controllers.UsersController("login", DB))
	r.POST("user/log-out", controllers.AuthMiddleware(DB), controllers.UsersController("login", DB))
	r.GET("secret/", controllers.AuthMiddleware(DB), controllers.SecretsController("index", DB))
	r.Run()
}
