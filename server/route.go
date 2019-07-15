package main

import (
	"UrlShortener/server/handlers/HomeController"

	"UrlShortener/server/handlers/UrlController"

)


func InitializeRoutes() {



	router.LoadHTMLGlob("templates/*")

	router.GET("/reshmi/:url", UrlController.GetLong)

	router.GET("/", HomeController.Index)

	router.POST("/form_post", UrlController.Create)

	router.GET("/file" ,UrlController.FileForm)

	router.POST("/uploadfile",UrlController.Getfile)





}
