package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/rpl/initializers"
	"github.com/hardytee1/rpl/models"
	"github.com/hardytee1/rpl/utils"
)

func CreateWts(c *gin.Context) {
	var body struct {
		Mata_pelajaran string `json:"mata_pelajaran" binding:"required"`
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

	//gets the teacher id
	var teacher models.Teacher
	err := initializers.DB.Where("user_id = ?", usr.ID).First(&teacher).Error
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Teacher record not found", nil)
		return
	}

	wts := models.Wts{
		Mata_pelajaran: body.Mata_pelajaran,
		TeacherID:      teacher.ID,
	}

	result := initializers.DB.Create(&wts)
	if result.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create wts post", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	utils.RespondSuccess(c, wts, "WTS posted successfully")
}
