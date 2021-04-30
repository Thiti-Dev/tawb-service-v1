package protected_route

import (
	"net/http"
	"strings"

	"github.com/Thiti-Dev/tawb-service-v1/packages/jwt"
	"github.com/gin-gonic/gin"
)


func ProtectedRoute() gin.HandlerFunc{
	return func(c *gin.Context){
		auth_header := c.Request.Header.Get("Authorization")
		if auth_header == ""{
			c.String(http.StatusForbidden, "No Authorization header provided")
			c.Abort()
			return
		}
		token := strings.TrimPrefix(auth_header, "Bearer ")
		if token == auth_header {
			c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
			c.Abort()
			return
		}
		isValid, claimData := jwt.DecodeJwt(token)
		if !isValid{
			c.String(http.StatusForbidden, "Invalid Authorization header")
			c.Abort()
			return
		}
		
		// If valid
		c.Set("user",claimData)
		c.Next()
	}
}