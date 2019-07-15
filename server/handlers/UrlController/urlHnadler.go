package UrlController

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"UrlShortener/server/model"

	"fmt"
	"UrlShortener/server/ShortUrl"

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

	fmt.Println("Yo")

	var long string
	var shortUrl string
	var err error

	fmt.Println(c.Param("url"))
	fmt.Println("getLong entered "+c.Param("url"))

	shortUrl = "localhost:8882/reshmi/"+ c.Param("url")

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
	readFile(file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	c.Redirect(301 ,"/")



}

func readFile(f string){

	var longUrl string
	file, err := os.Open("saved/"+f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		longUrl = scanner.Text()

		shortURL := ShortUrl.GetShortUrl(longUrl,"default")

		fmt.Println(shortURL)
	}

	if err := scanner.Err(); err != nil {

		log.Fatal(err)
	}

}