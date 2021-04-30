package datakit

import (
	"gorm.io/gorm"
	"lib_new_website_server/app/database/model"
)
/**
 * @Description: 报名服务是否开启
 * @param database: 数据库连接指针
 * @return bool: 是否开启
 */
func IsSignUpOpened(database *gorm.DB) bool {
	var signUpInfo model.Menu
	db := database.Where("menu", "报名").First(&signUpInfo)
	// 如果没有记录(报名没有添加到数据库内)或已经被关闭(状态为0)则返回假, 否则返回真
	return !(db.Error == gorm.ErrRecordNotFound || signUpInfo.Status == 0)
}
