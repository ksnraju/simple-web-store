package cart

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksnraju/simple-web-store/entity"
	"gorm.io/gorm"
)

func AddToCart(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var item entity.Item
		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var user entity.User
		tx := db.Where(&entity.User{Token: c.Request.Header.Get("X-User-Token")}).First(&user)
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't Fetch User Information",
			})
			return
		}

		cartItem := entity.CartItem{CartID: user.CartID, ItemID: item.ID}
		tx = db.Create(&cartItem)
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to Add item to cart",
			})
			return
		}
		c.JSON(http.StatusCreated, cartItem)
	}
}
