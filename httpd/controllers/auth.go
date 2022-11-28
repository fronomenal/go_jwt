package controllers

import (
	"net/http"

	"github.com/fronomenal/go_jwt/httpd/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	auth := func(c *gin.Context) {

		authToken, err := c.Cookie("Authorization")
		if err != nil || authToken == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if id, ok := validateToken(authToken); ok {
			var user models.User
			db.First(&user, id)

			if user.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			c.Next()

		}

	}

	return auth
}
