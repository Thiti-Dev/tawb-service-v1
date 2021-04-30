package helpers

import "github.com/gin-gonic/gin"

type responseMessage struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseMsg(c *gin.Context, code int, msg string, data interface{}) error {
	responseMsg := &responseMessage{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	c.JSON(code,responseMsg)
	return nil
}