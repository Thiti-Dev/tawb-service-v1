package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func Initialize_Route(router *gin.Engine){
	v1 := router.Group("/v1")
	{
		v1.GET("/test",func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{"status": "success"})

		})
	}
}