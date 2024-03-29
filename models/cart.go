package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint
	User      User `gorm:"foreignkey:UserID"`
	ProductID uint
	Product   User `gorm:"foreignkey:ProductID"`
	Quantity  uint `json:"quantity" gorm:"int"`
}
