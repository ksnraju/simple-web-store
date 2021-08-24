package item

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksnraju/simple-web-store/entity"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var item entity.Item
		err := c.BindJSON(&item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		tx := db.Create(&item)
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Couldn't Create Item",
			})
			return
		}
		c.JSON(http.StatusCreated, item)
	}
}
