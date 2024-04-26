package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/rpl/initializers"
	"github.com/hardytee1/rpl/models"
	"github.com/hardytee1/rpl/utils"
)

func CreateWtb(c *gin.Context) {
	var body struct {
		Mata_pelajaran string `json:"mata_pelajaran" binding:"required"`
		Jumlah_pertemuan string `json:"jumlah_pertemuan" binding:"required"`
		Harga string `json:"harga" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Failed to create", map[string]interface{}{"error": err.Error()})
		return
	}

	//get the current user
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

	wtb := models.Wtb{
		Mata_pelajaran: body.Mata_pelajaran,
		Jumlah_pertemuan: body.Jumlah_pertemuan,
		Harga: body.Harga,
		UserID: usr.ID,
	}

	result := initializers.DB.Create(&wtb)
	if result.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create wts post", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	utils.RespondSuccess(c, wtb, "WTS posted successfully")
}
