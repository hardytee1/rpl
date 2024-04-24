package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hardytee1/rpl/initializers"
	"github.com/hardytee1/rpl/models"
	"github.com/hardytee1/rpl/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	//get the email pass
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//look up requested user
	var user models.User
	initializers.DB.Where("username = ?", body.Username).First(&user)

	if user.ID == 0 {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid username or password", gin.H{"username": "Invalid credentials"})
		return
	}
	//compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid username or password", gin.H{"password": "Invalid credentials"})
		return
	}

	//generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to generate token", nil)
		return
	}

	//Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	utils.RespondSuccess(c, gin.H{"token": tokenString, "user": user}, "Login successful")
}
