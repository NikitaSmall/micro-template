package router

import (
	"github.com/gin-gonic/gin"
)

// Function returns router (pointer to engine object) with default middlewares enabled.
// You may redefine these middlewares for router by changing the code below.
func newRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return r
}

func CreateRouter() *gin.Engine {
	router := newRouter()

	router.GET("/", helpHandler)
	router.GET("/ping", pingHandler)

	return router
}
