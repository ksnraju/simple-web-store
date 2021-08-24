package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksnraju/simple-web-store/entity"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var user entity.User
		var cart entity.Cart
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		tx := db.Begin()
		err = tx.Create(&cart).Error
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't create user cart",
			})
			return
		}
		user.Cart = cart
		err = tx.Create(&user).Error
		fmt.Println("user Creation error")
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't create user",
			})
			return
		}
		tx.Commit()
		c.JSON(200, user)
	}
}
