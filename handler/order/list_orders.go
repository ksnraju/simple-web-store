package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksnraju/simple-web-store/entity"
	"gorm.io/gorm"
)

func ListOrders(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var orders []entity.Order
		var user entity.User
		token := c.Request.Header.Get("X-User-Token")
		if token != "" {
			tx := db.Where(&entity.User{Token: token}).First(&user)
			if tx.Error != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Invalid user token",
				})
				return
			}
		}
		var result *gorm.DB
		if user.ID != 0 {
			result = db.Where(&entity.Order{UserId: uint64(user.ID)}).Find(&orders)
		} else {
			result = db.Find(&orders)
		}
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't fetch order detials",
			})
			return
		}
		c.JSON(200, orders)
	}
}
