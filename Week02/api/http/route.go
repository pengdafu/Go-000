package http

import (
	"github.com/gin-gonic/gin"
	"week02/service"
)

func Route() *gin.Engine {
	r := gin.New()

	userGroup := r.Group("/users")
	{
		userGroup.GET("/:id", service.FindUser)
	}

	return r
}
