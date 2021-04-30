package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description: 跳转到软件孵化实验室首页
 * @param context: Gin上下文
 */
func GetIndexWebPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}
/**
 * @Description: 跳转到报名页
 * @param context： Gin上下文
 */
func GetNewsPage(context *gin.Context) {
	context.HTML(http.StatusOK, "signup.html", gin.H{})
}
