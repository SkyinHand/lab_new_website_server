package datakit

import (
	"gorm.io/gorm"
	"lib_new_website_server/app/database/model"
	"lib_new_website_server/app/schemas"
	"lib_new_website_server/app/utils"
)

func GetUser(database *gorm.DB, userID string) (*gorm.DB, *model.User) {
	var userInfo model.User
	db := database.Where("user_id", userID).First(&userInfo)
	return db, &userInfo
}

func IsUserExists(database *gorm.DB, userID string) bool {
	db, _ := GetUser(database, userID)
	return !(db.Error == gorm.ErrRecordNotFound)
}

func SendCustomerToTable(database *gorm.DB, customerStruct schemas.NewCustomer) int {
	if IsUserExists(database, customerStruct.Uid) {
		return utils.PrimaryKeyInsertError
	}
	// 创建一个用户
	result := database.Create(&model.User{
		UserID:    customerStruct.Uid,
		Name:      customerStruct.Name,
		Gender:    customerStruct.Gender,
		QQ:        customerStruct.QQ,
		ClassName: customerStruct.ClassName,
		Profile:   customerStruct.Profile,
	})
	if result.Error != nil {
		return utils.OtherInsertError
	} else {
		return utils.InsertSuccess
	}
}

func QueryUserStatus(database *gorm.DB, userID string) int {
	_, userInfo := GetUser(database, userID)
	return userInfo.Status
}