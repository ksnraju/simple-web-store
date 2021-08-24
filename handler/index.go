package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
		c.Abort()
	}
}
