package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/rpl/initializers"
	"github.com/hardytee1/rpl/models"
	"github.com/hardytee1/rpl/utils"
)

func GetWts(c *gin.Context) {
	var wts []models.Wts
	result := initializers.DB.Find(&wts)

	// Check for errors in the query
	if result.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to retrieve WTS records", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	// Send the fetched Wts records in the response
	utils.RespondSuccess(c, wts, "WTS records retrieved successfully")
}
