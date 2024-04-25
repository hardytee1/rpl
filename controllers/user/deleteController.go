package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/rpl/initializers"
	"github.com/hardytee1/rpl/models"
	"github.com/hardytee1/rpl/utils"
)

func DeleteAccount(c *gin.Context) {
	//get id url
	id := c.Param("id")
	fmt.Println("this is id:", id);

	//delete account
	initializers.DB.Delete(&models.User{}, id)

	//respond
	utils.RespondSuccess(c, nil, "Account deleted successfully")
}