package model

import (
	"database/sql"
	"time"
)

type User struct {
	UserID string `gorm:"primarykey; size:12; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Name string `gorm:"size:28; not null"`
	Gender string `gorm:"size:2; not null"`
	QQ string `gorm:"size:20; not null"`
	ClassName string `gorm:"size: 255; not null"`
	Profile string `gorm:"type:longtext;not null"`
	Status int `gorm:"type:tinyint;not null; default: 0"`
}
