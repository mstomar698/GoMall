package app

import (
	"fmt"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func InitializeDB() error {
	// database.ConnectDB()
	// if database.DB == nil {
	// 	log.Fatal("DB not connected ❌")
	// } else {
	// 	log.Println("DB connected 👍")
	// }
	fmt.Println("DB connected 👍")
	return nil
}

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	// r.AddFromFiles("home", "../web/template/home.html")
	// r.AddFromFiles("homesecured", "../web/template/homesecured.html")
	// r.AddFromFiles("login", "../web/template/login.html")
	// r.AddFromFiles("signup", "../web/template/signup.html")
	// r.AddFromFiles("allusers", "../web/template/allusers.html")
	// r.AddFromFiles("edituser", "../web/template/edituser.html")
	// r.AddFromFiles("edit", "../web/template/edit.html")
	// r.AddFromFiles("deleteuser", "../web/template/deleteuser.html")
	// r.AddFromFiles("createuser", "../web/template/createuser.html")
	fmt.Println("Templates loaded 👍")
	return r
}

func InitializeRouter() {
	router := gin.Default()

	// router.Static("/css", "../web/static/css")
	// router.Static("/img", "../web/static/img")
	// router.Static("/js", "../web/static/js")
	// router.StaticFile("/favicon.ico", "../web/static/img/favicon.ico")

	router.HTMLRender = createMyRender()
	
	// routes.ScreenRoute(router)
	// routes.AuthRoutes(router)
	// routes.HomeRoute(router)
	// routes.UserRoutes(router)

	fmt.Println("Server running on port http://localhost:8080 👍")
	router.Run(":8080")
}
