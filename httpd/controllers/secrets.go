package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SecretsController(handler string, db *gorm.DB) gin.HandlerFunc {
	defpage := func(c *gin.Context) {
		respond(c, 200, "You are Special!")
	}

	switch handler {
	default:
		return defpage
	}
}
