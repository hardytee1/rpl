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

func CreateAccount(c *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email string `json:"email" binding:"required"`
		Nama string `json:"nama" binding:"required"`
		Biodata string `json:"biodata" binding:"required"`
		Notelpon string `json:"notelpon" binding:"required"`
		Role string `json:"role" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Failed to create", map[string]interface{}{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to hash password", map[string]interface{}{"error": err.Error()})
		return
	}

	user := models.User{
		Username: body.Username,
		Password: string(hash),
		Email:    body.Email,
		Nama:     body.Nama,
		Biodata:  body.Biodata,
		Notelpon: body.Notelpon,
		Role: models.Role(body.Role),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create user", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"roles": string(user.Role), // Convert Role to string for JWT claims
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to generate JWT token", map[string]interface{}{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	utils.RespondSuccess(c, nil, "Admin registered successfully")
}
