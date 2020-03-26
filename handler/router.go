package handler

import "github.com/gin-gonic/gin"

func Router(router *gin.Engine) {

	h := &Handler{}
	v1 := router.Group("/v1")
	{
		list := v1.Group("/user")
		{
			list.POST("/", h.createUser)
			list.GET("/", h.getUser)
			list.GET("/:id", h.getUserByID)
			list.DELETE("/:id", h.deleteUser)
			list.PATCH("/:id", h.updateUser)
		}
	}
}
