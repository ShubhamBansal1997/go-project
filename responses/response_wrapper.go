package responses

import (
	"github.com/gin-gonic/gin"
)

//	type ErrorResponse struct {
//		Message string `json:"message"`
//	}
type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, statusCode int, data interface{}) {
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// nolint // context.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
	c.JSON(statusCode, data)
	return
}

func MessageResponse(c *gin.Context, statusCode int, message string) {
	Response(c, statusCode, Data{
		Code:    statusCode,
		Message: message,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	Response(c, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}
