package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ksnraju/simple-web-store/entity"
	"gorm.io/gorm"
)

func UserLogin(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var user entity.User
		var login entity.UserLogin
		if err := c.BindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		tx := db.Where(&entity.User{Username: login.Username}).First(&user)
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't fetch user detials",
			})
			return
		}
		if user.Password != login.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Credentials",
			})
			return
		}
		user.Token = uuid.New().String()
		tx = db.Save(&user)
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate token",
			})
			return
		}
		c.JSON(200, gin.H{
			"id":      user.ID,
			"name":    user.Name,
			"token":   user.Token,
			"cart_id": user.CartID,
		})

	}
}
