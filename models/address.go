package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Value  string `json:"value" gorm:"type:text"`
	UserID uint
	User   User `gorm:"foreignkey:UserID"`
}
