package routes

import (
	user_controller "github.com/Thiti-Dev/tawb-service-v1/controller/user"
	"github.com/gin-gonic/gin"
)


func Initialize_Route(router *gin.Engine){
	api := router.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")
	{
		user.POST("/register", func(c *gin.Context) {
			user_controller.RegisterUser(c)
		}) // we wrapped func in func again because GIN support only func(*gin.Context) without any return types
	}
}