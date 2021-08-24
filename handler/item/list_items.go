package item

import (
	"github.com/gin-gonic/gin"
	"github.com/ksnraju/simple-web-store/entity"
	"github.com/ksnraju/simple-web-store/handler"
	"gorm.io/gorm"
)

func ListItems(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var items []entity.Item
		if result := db.Find(&items); result.Error == nil {
			c.JSON(200, items)
		} else {
			handler.AppError(500, "Couldn't fetch item details")(c)
		}
	}
}
