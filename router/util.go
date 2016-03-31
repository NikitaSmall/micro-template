package router

import (
	"github.com/gin-gonic/gin"
)

func baseJsonMessage(msg string) gin.H {
	return gin.H{"message": msg}
}
