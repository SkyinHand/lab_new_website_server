package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
/**
 * @Description: 无路由跳转到404页面
 * @param context
 */
func NoRouterController(context *gin.Context) {
	context.HTML(http.StatusOK, "404.html", gin.H{})
}
