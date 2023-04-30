package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mstomar698/GoMall/app/controllers"
)

func ScreenRoute(router *gin.Engine) {
	router.GET("/login", controllers.LoginScreen())
	router.GET("/signup", controllers.SignupScreen())
	router.GET("/createuser", controllers.CreateUserScreen())
	router.GET("/edituser/:userId", controllers.EditUserScreen())
	router.GET("/deleteuser/:userId", controllers.DeleteUserScreen())
	router.GET("/createproduct", controllers.CreateProductScreen())
}
