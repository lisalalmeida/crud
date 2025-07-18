package controller

import (
	"api/src/controller/modules/request"
	"api/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"fmt"
)

func CreatUser(c *gin.Context){
	
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBandRequestError(fmt.Sprintf("error=%s", err))
		
		c.JSON(restErr.Code, restErr)
		return
	}
	
	fmt.Println(userRequest)

}
