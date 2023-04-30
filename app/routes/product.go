package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mstomar698/GoMall/app/controllers"
	"github.com/mstomar698/GoMall/app/utils"
)

func ProductRoutes(router *gin.Engine) {
	router.POST("/createProduct", controllers.CreateProduct())
	router.Use(utils.Authenticate())
	router.GET("/product", controllers.ProductScreen())
	router.GET("/dashboard", controllers.DashBoardScreen())
}
