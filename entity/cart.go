package entity

import "time"

type Cart struct {
	ID        uint       `json:"id" gorm:"PrimaryKey;AutoIncrement"`
	CreatedAt time.Time  `json:"-" gorm:"AutoCreateTime"`
	UpdatedAt time.Time  `json:"-" gorm:"AutoUpdateTime"`
	CartItems []CartItem `gorm:"foreignKey:CartID"`
}

type CartItem struct {
	CartID   uint `binding:"omitempty,gt=0"`
	ItemID   uint `binding:"required,gt=0"`
	Item     Item `json:"-"`
	Cart     Cart `json:"-"`
	Quantity uint `binding:"omitempty,gte=0"`
}
