package middleware

import (
	"github.com/gin-gonic/gin"
	"lib_new_website_server/app/utils"
	"log"
	"runtime/debug"
)
/**
 * @Description: 服务器错误后恢复现场
 * @return gin.HandlerFunc
 */
func Recover() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 打印错误信息
				log.Printf("panic: %v\n", err)
				debug.PrintStack()
				// 返回统一错误信息格式
				utils.ReturnErrorJson(context, "服务器错误, 请稍后重试")
			}
		}()
		// 恢复现场
		context.Next()
	}
}
