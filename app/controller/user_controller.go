package controller

import (
	"github.com/gin-gonic/gin"
	"lib_new_website_server/app/database"
	"lib_new_website_server/app/database/datakit"
	"lib_new_website_server/app/schemas"
	"lib_new_website_server/app/utils"
)

/**
 * @Description: 报名软件孵化实验室
 * @param context
 */
func PostNewCustomer(context *gin.Context) {
	// 参数验证
	var customerData schemas.NewCustomer
	// 如果参数传递错误
	if err := context.ShouldBindJSON(&customerData); err != nil {
		utils.ReturnErrorJson(context, "参数错误")
		return
	}
	// 获取数据库对象
	dataInstance := database.GetDataBase()
	// 验证报名是否开启
	if !datakit.IsSignUpOpened(dataInstance) {
		utils.ReturnErrorJson(context, "报名尚未开启")
		return
	}
	// 将用户信息加入到数据库中
	code := datakit.SendCustomerToTable(dataInstance, customerData)
	// 比 if - else 更加优雅的 enum + switch 写法
	switch code {
	case utils.PrimaryKeyInsertError:
		utils.ReturnErrorJson(context, "你已经报名, 请勿重复报名")
	case utils.OtherInsertError:
		utils.ReturnErrorJson(context, "报名失败, 请稍后再试")
	default:
		utils.ReturnSuccessJson(context, "报名成功, 等待管理员审核")
	}
}

/**
 * @Description: 查询系统
 * @param context
 */
func QueryCustomer(context *gin.Context) {
	userID := context.Param("userID")
	// 创建database实例
	base := database.GetDataBase()
	if !datakit.IsUserExists(base, userID) {
		utils.ReturnErrorJson(context, "你还没有报名, 请先报名")
		return
	}
	status := datakit.QueryUserStatus(base, userID)
	switch status {
	case utils.UserWaitting:
		utils.ReturnSuccessJson(context, "未出结果, 等待审核")
	case utils.UserPass:
		utils.ReturnSuccessJson(context, "你通过了审核, 欢迎加入软件孵化实验室 :)")
	case utils.UserNotPass:
		utils.ReturnErrorJson(context, "由于人数原因, 你未通过审核, 但是已经把您加入了人才库, 期待与你的下次相见")
	}
}
