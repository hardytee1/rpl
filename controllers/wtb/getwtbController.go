package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/rpl/initializers"
	"github.com/hardytee1/rpl/models"
	"github.com/hardytee1/rpl/utils"
)

func GetWtb(c *gin.Context) {
	var wtb []models.Wtb
	result := initializers.DB.Find(&wtb)

	// Check for errors in the query
	if result.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to retrieve WTB records", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	// Send the fetched Wts records in the response
	utils.RespondSuccess(c, wtb, "WTB records retrieved successfully")
}
