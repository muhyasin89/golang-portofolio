package utils

import "github.com/gin-gonic/gin"

func Result(code, message, data) c.JSON {
	return c.JSON(code, gin.H{
		"data":    data,
		"message": message
	})
}