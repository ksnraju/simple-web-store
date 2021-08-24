package middleware

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ksnraju/simple-web-store/handler"
)

func ExtractToken(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	token := ""
	if authorization != "" {
		valid, err := regexp.Match(`(?i)^Bearer [a-z0-9\-]*$`, []byte(authorization))
		if err != nil || !valid {
			handler.AppError(400, "Invalid Authrization Header")(c)
			return
		} else {
			token = strings.Split(authorization, " ")[1]
		}
	}
	c.Request.Header.Add("X-User-Token", token)
	c.Next()
}
