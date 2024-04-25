package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/rpl/utils"
	"github.com/hardytee1/rpl/models"
)

func Validate(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.RespondError(c, http.StatusUnauthorized, "User not found in context", nil)
		return
	}

	usr, ok := user.(models.User) // Type assertion
	if !ok {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid user type in context", nil)
		return
	}

	tokenString, err := c.Cookie("Authorization")
	
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	type ResponseData struct {
		ID       uint    `json:"ID"` // Field tags ensure correct key names in JSON output
		Username string `json:"Username"`
		Email    string `json:"Email"`
		Nama	 string `json:"Nama"`
		Biodata  string `json:"Biodata"`
		Notelpon string `json:"Notelpon"`
		Role	 string `json:"Role"`
		Token    string `json:"Token"`
	}

	responseData :=  ResponseData{
		ID:       usr.ID,
		Username: usr.Username,
		Email:    usr.Email,
		Nama:	  usr.Nama,
		Biodata:  usr.Biodata,
		Notelpon: usr.Notelpon,
		Role: 	  string(usr.Role),
		Token:	  tokenString,
	}

	utils.RespondSuccess(c, responseData, "User validated successfully")
}
