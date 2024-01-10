package utils

import (
	"github.com/gin-gonic/gin"
)

type GeneralResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Result(c *gin.Context, code int, message string, data interface{}) {
	response := GeneralResponse{
		Data:    data,
		Message: message,
	}

	c.JSON(code, response)
}
