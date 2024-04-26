package router

import (
	controllers "github.com/hardytee1/rpl/controllers/teacher"
	middleware "github.com/hardytee1/rpl/middleware"
	models "github.com/hardytee1/rpl/models"

	"github.com/gin-gonic/gin"
)

func TeacherRouter(router *gin.Engine) {
	teacherRoutes := router.Group("")

	teacherRoutes.POST("/newteacher", middleware.RequireAuth, middleware.RequireRole(models.RoleUser), controllers.CreateTeacher)
}
