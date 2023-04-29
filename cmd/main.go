package main

import "github.com/mstomar698/GoMall/app"

func main() {
	app.InitializeDB()
	app.InitializeRouter()
}
