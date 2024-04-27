package controllers

import (
	"net/http"
	"mime/multipart"
	"github.com/gin-gonic/gin"
	"github.com/hardytee1/rpl/initializers"
	"github.com/hardytee1/rpl/models"
	"github.com/hardytee1/rpl/utils"
)

type CreateTeacherRequest struct {
	Pendidikan string                `json:"pendidikan" binding:"required"`
	Bukti_1    *multipart.FileHeader `form:"bukti_1" binding:"required"`
}


func CreateTeacher(c *gin.Context) {
	var req CreateTeacherRequest

	// Bind JSON data
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Failed to create", map[string]interface{}{"error": err.Error()})
		return
	}

	// Get the file
	file, err := req.Bukti_1.Open()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to open file", map[string]interface{}{"error": err.Error()})
		return
	}
	defer file.Close()

	
	filepath := "C:/Users/Hp/Desktop/tugas skuull/website/rpl/file.pdf"
	err = c.SaveUploadedFile(req.Bukti_1, filepath)
	if err != nil {
	    utils.RespondError(c, http.StatusInternalServerError, "Failed to save file", map[string]interface{}{"error": err.Error()})
	    return
	}

	// Get the current user
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

	// Create teacher object with file path or other necessary data
	teacher := models.Teacher{
		Pendidikan: req.Pendidikan,
		Bukti_1:    filepath, // Use the file path here
		UserID:     usr.ID,
	}

	// Persist teacher data to the database
	result := initializers.DB.Create(&teacher)
	if result.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create teacher", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	// Update user role to TEACHER
	updateRole := initializers.DB.Model(&usr).Update("Role", "TEACHER")
	if updateRole.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update role", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	utils.RespondSuccess(c, nil, "Teacher created successfully")
}

