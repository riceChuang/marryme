package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func Router(router *gin.Engine) {

	//h:= &Handler{}
	v1 := router.Group("/v1")
	{
		list := v1.Group("/user")
		{
			list.POST("/")
		}
	}
}
