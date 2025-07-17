package main

import (
	"net/http"	
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.LoadHTMLFiles("site/index.html")
	
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.InitRoutes(&r.RouterGroup)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
