package router

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", helpHandler)
	router.GET("/ping", pingHandler)
	return router
}
