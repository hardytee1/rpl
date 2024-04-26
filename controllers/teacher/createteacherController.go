package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/rpl/initializers"
	"github.com/hardytee1/rpl/models"
	"github.com/hardytee1/rpl/utils"
)

func CreateTeacher(c *gin.Context) {
	var body struct {
		Pendidikan string `json:"pendidikan" binding:"required"`
		Bukti_1 string `json:"bukti_1" binding:"required"`
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

	teacher := models.Teacher{
		Pendidikan: body.Pendidikan,
		Bukti_1: body.Bukti_1,
		UserID: usr.ID,
	}

	result := initializers.DB.Create(&teacher)
	if result.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create teacher", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	updateRole := initializers.DB.Model(&usr).Update("Role", "TEACHER")
	if updateRole.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update role", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	utils.RespondSuccess(c, nil, "Teacher created successfully")
}
