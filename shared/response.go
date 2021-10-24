package shared

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, gin.H{
		"message": msg,
		"data":    data,
	})
}

func ErrorResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, gin.H{
		"message": msg,
		"errors":    data,
	})
}