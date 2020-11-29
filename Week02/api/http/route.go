package http

import (
	"github.com/gin-gonic/gin"
	"week02/service"
)

func Route() *gin.Engine {
	r := gin.New()
	{
		userSrv := service.New()
		userGroup := r.Group("/users")
		userGroup.GET("/:id", userSrv.FindUser)
	}

	return r
}
