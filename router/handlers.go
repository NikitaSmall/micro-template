package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitasmall/master-service/lib"
)

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, baseJSONMessage("PONG!"))
}

func helpHandler(c *gin.Context) {
	if msg, err := lib.GetHelp(); err == nil {
		c.JSON(http.StatusOK, baseJSONMessage(msg))
	} else {
		c.JSON(http.StatusInternalServerError, baseJSONMessage(err.Error()))
	}
}
