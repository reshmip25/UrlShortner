package main

import (
	"urlProject/server/handlers/HomeController"
	_ "urlProject/server/handlers/HomeController"
	"urlProject/server/handlers/UrlController"
	_ "urlProject/server/handlers/UrlController"
)


func InitializeRoutes() {

	router.LoadHTMLGlob("templates/*")

	router.GET("/short/:url1", UrlController.GetLong)

	router.GET("/", HomeController.Index)

	router.POST("/form_post", UrlController.Create)

	router.GET("/file" ,UrlController.FileForm)

	router.POST("/file",UrlController.Getfile)





}
