package router

import (
	controllers "github.com/hardytee1/rpl/controllers/wtb"
	middleware "github.com/hardytee1/rpl/middleware"
	models "github.com/hardytee1/rpl/models"

	"github.com/gin-gonic/gin"
)

func WtbRouter(router *gin.Engine) {
	wtbRoutes := router.Group("")

	wtbRoutes.GET("/allwtb", controllers.GetWtb)
	wtbRoutes.POST("/newwtb", middleware.RequireAuth, middleware.RequireRole(models.RoleUser), controllers.CreateWtb)
}
