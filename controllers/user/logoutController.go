package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hardytee1/rpl/utils"
)

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	utils.RespondSuccess(c, nil, "Logout successful")
}
