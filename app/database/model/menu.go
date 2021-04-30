package model

import "gorm.io/gorm"
/**
 * @Description: 面板服务ORM
 */
type Menu struct {
	gorm.Model
	Menu   string `gorm:"size:255"`
	Status int    `gorm:"type:tinyint"`
}
