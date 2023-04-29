package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mstomar698/GoMall/app/controllers"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login())
	router.POST("/signup", controllers.Signup())
	router.POST("/logout", controllers.Logout())
}
