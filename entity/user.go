package entity

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"PrimaryKey;AutoIncrement"`
	CreatedAt time.Time `json:"-" gorm:"AutoCreateTime" `
	UpdatedAt time.Time `json:"-" gorm:"AutoUpdateTime"`
	Name      string    `json:"name" gorm:"not null" binding:"required,min=1,max=50,alphaspace"`
	Username  string    `json:"username" gorm:"UniqueIndex" binding:"required,min=1,max=20,username"`
	Password  string    `json:"password" gorm:"not null" binding:"required,min=1,max=20,password"`
	Token     string    `json:"token" binding:"omitempty,uuid4"`
	CartID    uint      `json:"cart_id"`
	Cart      Cart      `json:"-" gorm:"foreignKey:CartID"`
	Orders    []Order   `json:"-"`
}

type UserInfo struct {
	ID       uint
	Name     string
	Username string
	CartID   uint
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
