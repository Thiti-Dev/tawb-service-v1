package routes

import (
	user_controller "github.com/Thiti-Dev/tawb-service-v1/controller/user"
	"github.com/Thiti-Dev/tawb-service-v1/middlewares/protected_route"
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
		user.POST("/login", func(c *gin.Context) {
			user_controller.SignInWithCredential(c)
		})
		user.GET("/authcheck", protected_route.ProtectedRoute(), func(c *gin.Context){
			user_controller.GetUserCredential(c)
		})
	}
}