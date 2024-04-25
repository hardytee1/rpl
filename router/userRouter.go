package router

import (
	controllers "github.com/hardytee1/rpl/controllers/user"
	middleware "github.com/hardytee1/rpl/middleware"
	models "github.com/hardytee1/rpl/models"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userRoutes := router.Group("")

	// userRoutes.GET("/validate",middleware.RequireAuth, middleware.RequireRole(models.RoleAdmin), controllers.Validate)
	// contoh buat yg akses cuman boleh role tertentu

	userRoutes.POST("/register", controllers.Register)
	userRoutes.POST("/login", controllers.Login)
	userRoutes.GET("/validate", middleware.RequireAuth, controllers.Validate)
	userRoutes.POST("/logout", middleware.RequireAuth, controllers.Logout)
	userRoutes.POST("/createacc", controllers.CreateAccount)
	userRoutes.DELETE("/deleteacc/:id", middleware.RequireAuth, middleware.RequireRole(models.RoleAdmin), controllers.DeleteAccount)
}
