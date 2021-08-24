package order

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ksnraju/simple-web-store/entity"
	"github.com/ksnraju/simple-web-store/handler"
	"gorm.io/gorm"
)

func CreateOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var user entity.User
		tx := db.Where(&entity.User{Token: c.Request.Header.Get("X-User-Token")}).First(&user)
		if tx.Error != nil {
			handler.AppError(500, "Couldn't Fetch User Information")(c)
			return
		}
		CartID, err := strconv.ParseUint(c.Param("CartID"), 10, 32)
		if err != nil {
			handler.AppError(400, "CartID has to be integer")(c)
			return
		}
		if CartID != uint64(user.CartID) {
			handler.AppError(401, "Unauthorized Cart Access")(c)
			return
		}

		tx = db.Begin()

		if tx.Error != nil {
			handler.AppError(500, "Database Connection Error")(c)
			return
		}
		var cartItems []entity.CartItem
		tx.Where(&entity.CartItem{CartID: user.CartID}).Find(&cartItems)
		if tx.Error != nil {
			tx.Rollback()
			handler.AppError(500, "Unable to Fetch Cart Items")(c)
			return
		}
		if len(cartItems) < 1 {
			tx.Rollback()
			handler.AppError(400, "No Items in cart to create order")(c)
			return
		}

		order := entity.Order{
			UserId: uint64(user.ID),
		}
		tx.Create(&order)
		if tx.Error != nil {
			tx.Rollback()
			handler.AppError(500, "Unable to Create Order")(c)
			return
		}

		var orderItems []entity.OrderItem
		for _, item := range cartItems {
			orderItems = append(orderItems, entity.OrderItem{
				OrderID:  order.ID,
				ItemID:   item.ItemID,
				Quantity: item.Quantity,
			})
		}
		tx.Create(&orderItems)
		if tx.Error != nil {
			tx.Rollback()
			handler.AppError(500, "Unable to add Order Items")(c)
			return
		}
		tx.Where(&entity.CartItem{CartID: user.CartID}).Delete(&entity.CartItem{})
		if tx.Error != nil {
			tx.Rollback()
			handler.AppError(500, "Unable to delete cart Items")(c)
			return
		}
		tx.Commit()
		c.JSON(201, "Order Created Successfully")
	}
}
