package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mstomar698/GoMall/app/controllers"
	"github.com/mstomar698/GoMall/app/utils"
)

func HomeRoute(router *gin.Engine) {
	router.GET("/", controllers.Home())
	router.Use(utils.Authenticate())
	router.GET("/home", controllers.HomeSecured())
}
