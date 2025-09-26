package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name      string `gorm:"not null"`
	Price     int64  `gorm:"not null"`
	Desc      string `gorm:"type:text"`
	Stock     int64  `gorm:"not null"`
	PicInfo   string `gorm:"type:text"`
	Status    int32  `gorm:"not null"` // 0: 未上架, 1: 已上架
}

func (Product) TableName() string {
	return "products"
}
