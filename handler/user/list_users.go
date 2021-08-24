package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ksnraju/simple-web-store/entity"
	"github.com/ksnraju/simple-web-store/handler"
	"gorm.io/gorm"
)

func ListUsers(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var users []entity.UserInfo
		if result := db.Model(&entity.User{}).Find(&users); result.Error == nil {
			c.JSON(200, users)
		} else {
			handler.AppError(500, "Couldn't fetch user details")(c)
		}
	}
}
