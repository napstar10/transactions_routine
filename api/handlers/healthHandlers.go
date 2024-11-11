package handlers

import (
	"github.com/gin-gonic/gin"
)

func HelloHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
