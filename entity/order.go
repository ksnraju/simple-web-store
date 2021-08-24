package entity

import "time"

type Order struct {
	ID         uint        `json:"id" gorm:"PrimaryKey;AutoIncrement"`
	CreatedAt  time.Time   `json:"created_at" gorm:"AutoCreateTime"`
	UpdatedAt  time.Time   `json:"-" gorm:"AutoUpdateTime"`
	UserId     uint64      `json:"user_id" binding:"required,gte=0"`
	OrderItems []OrderItem `josn:"-" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	OrderID  uint `binding:"required,gte=0"`
	ItemID   uint `binding:"required,gte=0"`
	Item     Item
	Order    Order
	Quantity uint `binding:"required,gte=0"`
}
