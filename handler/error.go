package handler

import "github.com/gin-gonic/gin"

func AppError(code int, message string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(code, gin.H{"message": message})
		return
	}
}
