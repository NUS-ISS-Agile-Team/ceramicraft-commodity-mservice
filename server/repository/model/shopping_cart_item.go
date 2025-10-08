package model

import "time"

const (
	CartItemStatusUnselected = 1
	CartItemStatusSelected   = 2
)

type ShoppingCartItem struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	UserID       int       `gorm:"not null;index:idx_user_product,unique"`
	ProductID    int       `gorm:"not null;index:idx_user_product,unique"`
	Quantity     int       `gorm:"not null;default:1"`
	SelectStatus int       `gorm:"not null;default:0"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func (ShoppingCartItem) TableName() string {
	return "shopping_cart_items"
}
