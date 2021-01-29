package router

import (
	"github.com/gin-gonic/gin"
	"food-delivery/common"
)

func SetupRouter(r *gin.Engine, provider common.AppContext) {
	v1 := r.Group("/v1")
	notes := v1.Group("foods")

	notes.GET("/")
	notes.GET("/:food-id")
	notes.POST("/")
}