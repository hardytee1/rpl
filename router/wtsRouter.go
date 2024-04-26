package router

import (
	controllers "github.com/hardytee1/rpl/controllers/wts"
	middleware "github.com/hardytee1/rpl/middleware"
	models "github.com/hardytee1/rpl/models"

	"github.com/gin-gonic/gin"
)

func WtsRouter(router *gin.Engine) {
	wtsRoutes := router.Group("")

	wtsRoutes.GET("/allwts", controllers.GetWts)
	wtsRoutes.POST("/newwts", middleware.RequireAuth, middleware.RequireRole(models.RoleTeacher), controllers.CreateWts)
}
