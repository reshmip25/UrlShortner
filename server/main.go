package main



import (

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"UrlShortener/server/model"
	"log"
	"os"
)


var router *gin.Engine

func main() {

	f, err := os.Create("log.txt")
	if err != nil {

		log.Println(err)
	}


	log.SetOutput(f)

	f.Close()

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	InitializeRoutes()

	model.Index()

	router.Run(":8882")
}




