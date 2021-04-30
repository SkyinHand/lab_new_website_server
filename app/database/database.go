package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lib_new_website_server/app/utils"
	"log"
)

var DBInstance *gorm.DB

func Initdatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.GetConfig("database", "user"),
		utils.GetConfig("database", "pass"),
		utils.GetConfig("database", "host"),
		utils.GetConfig("database", "port"),
		utils.GetConfig("database", "database"))
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		DefaultStringSize: 256,

	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Connect Database Error: ", err.Error())
	}

	DBInstance = db

	// 迁移文件
	//_ = db.AutoMigrate(model.User{})
	//_ = db.AutoMigrate(model.Menu{})

	//db.Model(model.Menu{}).Create(&model.Menu{
	//	Menu: "报名",
	//	Status: 1,
	//})
	//
	//db.Model(model.Menu{}).Create(&model.Menu{
	//	Menu: "结果查询",
	//	Status: 1,
	//})
}

func GetDataBase() *gorm.DB {
	return DBInstance
}
