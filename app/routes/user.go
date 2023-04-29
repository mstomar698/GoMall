package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mstomar698/GoMall/app/controllers"
)

func UserRoutes(router *gin.Engine) {
	// router.Use(utils.Authenticate())
	// router.Use(utils.CheckUserType())
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetUser())
	router.PUT("/user/:userId", controllers.UpdateUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
	router.GET("/users", controllers.GetAllUsers())
}
