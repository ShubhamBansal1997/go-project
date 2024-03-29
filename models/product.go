package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(200);"`
	Description string  `json:"description" gorm:"type:varchar(500);"`
	Image       string  `json:"image" gorm:"type:varchar(500);"`
	Price       float64 `json:"price" gorm:"float;"`
	Category    string  `json:"category" gorm:"type:varchar(50);"`
	Sku         string  `json:"sku" gorm:"type:varchar(10);"`
}
