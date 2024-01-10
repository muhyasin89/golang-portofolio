package utils

import (
	"github.com/gin-gonic/gin"
)

func Result(c *gin.Context, code int, message string, data any) {
	return c.JSON(code, gin.H{
		"data":    data,
		"message": message
	})
}
