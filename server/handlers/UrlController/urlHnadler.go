package UrlController

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"urlProject/server/model"

	"fmt"
	"urlProject/server/ShortUrl"
	_ "urlProject/server/ShortUrl"
)

func Create(c *gin.Context) {

	longUrl := c.PostForm("Url")

	shortURL := ShortUrl.GetShortUrl(longUrl,c.PostForm("custom"))

	fmt.Println(shortURL)

	c.HTML(200, "result.html" ,gin.H{
		"success": shortURL,

	})

}

func GetLong(c *gin.Context){


	var long string

	var shortUrl string
	var err error

	fmt.Println("getLong entered"+c.Param("url1"))

	shortUrl = "localhost:8882/short/"+ c.Param("url1")

	fmt.Println("after appending"+shortUrl)

	long,err  = model.ShowLong(shortUrl)


	if(err!=nil){

		fmt.Println(err)

	} else {

		fmt.Println(long)

		c.Redirect(http.StatusMovedPermanently, long)

	}

}


func FileForm(c *gin.Context){

	c.HTML(200, "getfile.html" ,gin.H{
	})
}

func Getfile(c *gin.Context){

	file, err := c.FormFile("uploadfile")

	if err != nil {
		log.Fatal(err)
	}
	log.Println(file.Filename)

	err = c.SaveUploadedFile(file, "saved/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))




}