package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitasmall/master-service/lib"
)

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PONG!"})
}

func helpHandler(c *gin.Context) {
	if msg, err := lib.GetHelp(); err == nil {
		c.JSON(http.StatusOK, baseJsonMessage(msg))
	} else {
		c.JSON(http.StatusInternalServerError, baseJsonMessage(err.Error()))
	}
}
