package HomeController

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,

		// Use the index.html template
		"index.html",

		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Home Page",
		},

	)

}