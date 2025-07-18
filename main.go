package main

import (
	"log"
	"net/http"
	"api/src/routes"
	"github.com/gin-gonic/gin"
)

func main(){ 	
	r := gin.Default()

	r.LoadHTMLFiles("site/index.html")
	
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	
	routes.InitRoutes(&r.RouterGroup)
	
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
