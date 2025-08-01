package routes

import (
	"api/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup){
	
	r.GET("/getUserById/:userId", controller.FindUserID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/postInfoclient/:info", controller.InfoUserClient)
}
