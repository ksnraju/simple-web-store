package cart

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksnraju/simple-web-store/entity"
	"gorm.io/gorm"
)

func ListCarts(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var carts []entity.Cart
		if result := db.Find(&carts); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't fetch cart detials",
			})
			return
		}
		c.JSON(http.StatusOK, carts)
	}
}
