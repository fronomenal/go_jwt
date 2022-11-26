package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func getHash(text []byte) string {

	hash, err := bcrypt.GenerateFromPassword(text, 10)

	if err != nil {
		return ""
	}

	return string(hash)
}

func getToken(sub int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": sub,
		"nbf":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type ResBody interface {
	string | []UserResponse | map[string]interface{}
}

func respond[T ResBody](c *gin.Context, stat int, content T) {
	switch stat {
	case 200:
		c.JSON(http.StatusOK, gin.H{"status": 200, "body": content})
	case 400:
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": content})
	case 404:
		c.JSON(http.StatusNotFound, gin.H{"status": 404, "message": content})
	case 500:
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": content})
	default:
		c.JSON(http.StatusNoContent, gin.H{})

	}
}
