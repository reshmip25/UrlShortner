package main



import (

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"UrlShortener/server/model"

)


var router *gin.Engine


func main() {


	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	InitializeRoutes()
	model.Index()

	router.Run(":8882")
}




