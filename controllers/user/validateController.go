package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/rpl/utils"
)

func Validate(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.RespondError(c, http.StatusUnauthorized, "User not found in context", nil)
		return
	}

	utils.RespondSuccess(c, user, "User validated successfully")
}
