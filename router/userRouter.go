package router

import (
	controllers "github.com/hardytee1/rpl/controllers/user"
	middleware "github.com/hardytee1/rpl/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userRoutes := router.Group("")

	userRoutes.POST("/register", controllers.Register)
	userRoutes.POST("/login", controllers.Login)
	userRoutes.GET("/validate", middleware.RequireAuth, controllers.Validate)
	userRoutes.POST("/logout", middleware.RequireAuth, controllers.Logout)
}
