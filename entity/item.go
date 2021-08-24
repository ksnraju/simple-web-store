package entity

import "time"

type Item struct {
	ID        uint      `json:"id" gorm:"PrimaryKey;AutoIncrement"`
	CreatedAt time.Time `json:"-" gorm:"AutoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"AutoUpdateTime"`
	Name      string    `json:"name" gorm:"Unique" binding:"required,alphanumspace,min=1,max=50"`
	Price     float32   `json:"price" binding:"required,gt=0"`
}
