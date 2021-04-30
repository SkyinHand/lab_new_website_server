package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
/**
 * @Description: 返回正确的JSON
 * @param context
 * @param message
 */
func ReturnSuccessJson(context *gin.Context, message string) {
	context.JSON(
		http.StatusOK,
		gin.H{
			"code":    1,
			"message": message,
		},
	)
}
/**
 * @Description: 返回错误的JSON
 * @param context
 * @param message
 */
func ReturnErrorJson(context *gin.Context, message string) {
	context.JSON(
		http.StatusOK,
		gin.H{
			"code":    0,
			"message": message,
		},
	)
}
