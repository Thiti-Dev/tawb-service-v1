package main

import (
	"github.com/Thiti-Dev/tawb-service-v1/database"
	route_utils "github.com/Thiti-Dev/tawb-service-v1/routes/api/v1"
	"github.com/Thiti-Dev/tawb-service-v1/validator"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	validator.InitializeTranslator() // init validator translation (CUSTOM-DECLARED)
	database.ConnectMongoDatabase()
	route_utils.Initialize_Route(router) // init routes
	router.Run(":8080")
}