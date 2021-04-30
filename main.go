package main

import (
	"github.com/Thiti-Dev/tawb-service-v1/database"
	route_utils "github.com/Thiti-Dev/tawb-service-v1/routes/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectMongoDatabase()
	route_utils.Initialize_Route(router)
	router.Run(":8080")
}