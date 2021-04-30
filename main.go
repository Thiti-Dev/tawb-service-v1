package main

import (
	route_utils "github.com/Thiti-Dev/tawb-service-v1/routes/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	route_utils.Initialize_Route(router)
	router.Run(":8080")
}