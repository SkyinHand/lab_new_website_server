package datakit

import (
	"gorm.io/gorm"
	"lib_new_website_server/app/database/model"
	"lib_new_website_server/app/schemas"
	"lib_new_website_server/app/utils"
)
/**
 * @Description: 获取用户信息
 * @param database 数据库连接对象
 * @param userID 查询所需的用户学号
 * @return *gorm.DB 数据库对象
 * @return *model.User 用户数据对象
 */
func GetUser(database *gorm.DB, userID string) (*gorm.DB, *model.User) {
	var userInfo model.User
	db := database.Where("user_id", userID).First(&userInfo)
	return db, &userInfo
}
/**
 * @Description: 用户是否存在
 * @param database 数据库连接对象
 * @param userID 查询所需的用户学号
 * @return bool
 */
func IsUserExists(database *gorm.DB, userID string) bool {
	db, _ := GetUser(database, userID)
	return !(db.Error == gorm.ErrRecordNotFound)
}
/**
 * @Description: 将用户加入到数据库中
 * @param database
 * @param customerStruct 用户数据
 * @return int
 */
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
/**
 * @Description: 查询用户状态
 * @param database
 * @param userID
 * @return int
 */
func QueryUserStatus(database *gorm.DB, userID string) int {
	_, userInfo := GetUser(database, userID)
	return userInfo.Status
}
