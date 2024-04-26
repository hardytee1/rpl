package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hardytee1/rpl/initializers"
	"github.com/hardytee1/rpl/router"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	router.UserRouter(r)
	router.TeacherRouter(r)
	router.WtsRouter(r)
	router.WtbRouter(r)
	r.Run()
}