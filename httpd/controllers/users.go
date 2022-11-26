package controllers

import (
	"net/http"
	"strings"

	"github.com/fronomenal/go_jwt/httpd/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UsersController(handler string, db *gorm.DB) gin.HandlerFunc {

	index := func(c *gin.Context) {
		var allusrs []models.User

		if rez := db.Find(&allusrs); rez.Error != nil {
			respond(c, 500, "Error retrieving users")
			return
		} else if rez.RowsAffected == 0 {
			respond(c, 404, "No users registered")
			return
		}

		var resusrs []UserResponse

		for _, usr := range allusrs {
			resusrs = append(resusrs, UserResponse{usr.Name, usr.Email})
		}

		respond(c, 200, resusrs)

	}

	signup := func(c *gin.Context) {

		reqBody := UserRequest{}

		if c.Bind(&reqBody) != nil {
			respond(c, 400, "Problem with request body")
			return
		}

		if len(strings.Trim(reqBody.Email, " ")) < 3 || len(strings.Trim(reqBody.Name, " ")) < 3 || len(strings.Trim(reqBody.Pass, " ")) < 3 {
			respond(c, 400, "Must provide a name, email and password")
			return
		}

		hashedPass := getHash([]byte(reqBody.Pass))

		if len(hashedPass) == 0 {
			respond(c, 400, "Invalid password provided")
			return
		}

		newuser := models.User{Name: reqBody.Name, Email: reqBody.Email, Pass: hashedPass}

		if rez := db.Create(&newuser); rez.Error != nil {
			respond(c, 400, "User sign up failed due to duplicate mail")
			return
		}

		respond(c, 200, "Account created for "+newuser.Email)

	}

	login := func(c *gin.Context) {
		reqBody := UserRequest{}

		if c.Bind(&reqBody) != nil {
			respond(c, 400, "Problem with request body")
			return
		}

		if len(strings.Trim(reqBody.Email, " ")) < 3 || len(strings.Trim(reqBody.Pass, " ")) < 3 {
			respond(c, 400, "Must provide an email and password")
			return
		}

		result := map[string]interface{}{}

		if rez := db.Table("users").Where("email = ?", reqBody.Email).Select([]string{"id", "pass"}).Take(&result); rez.Error != nil {
			respond(c, 404, "Couldn't find requested user")
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(result["pass"].(string)), []byte(reqBody.Pass)); err != nil {
			respond(c, 400, "Wrong password provided")
			return

		}

		usrtoken, err := getToken(result["id"].(int64))
		if err != nil {
			respond(c, 500, "Error tokenizing user")
			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", "bearer "+usrtoken, 3600*24*30, "", "", false, true)

		respond(c, 200, usrtoken)

	}

	defpage := func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "Route handler mismatch"})
	}

	switch handler {
	case "index":
		return index
	case "sign-up":
		return signup
	case "login":
		return login
	default:
		return defpage
	}

}
