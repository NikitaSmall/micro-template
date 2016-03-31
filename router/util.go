package router

import (
	"github.com/gin-gonic/gin"
)

func baseJSONMessage(msg string) gin.H {
	return gin.H{"message": msg}
}
