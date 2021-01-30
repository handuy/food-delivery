package router

import (
	"food-delivery/common"

	"github.com/gin-gonic/gin"

	"food-delivery/module/food/foodtransport"
	"food-delivery/module/user/usertransport"
)

func SetupRouter(r *gin.Engine, provider common.AppContext) {
	v1 := r.Group("/v1")

	notes := v1.Group("foods")
	notes.GET("/", foodtransport.GetAllFood(provider))
	notes.GET("/:food-id", foodtransport.GetFoodById(provider))

	users := v1.Group("users")
	users.POST("/signup", usertransport.CreateUser(provider))
	users.POST("/login", usertransport.LoginUser(provider))
}