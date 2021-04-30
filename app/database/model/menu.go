package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Menu string `gorm:"size:255"`
	Status int `gorm:"type:tinyint"`
}
