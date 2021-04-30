package router

import (
	"github.com/gin-gonic/gin"
)
/**
 * @Description: 注册所有路由
 * @param engine
 */
func RegisterAllRouters(engine *gin.Engine) {

	// 注册Web相关路由
	RegisterWebRouter(engine)

	// 注册API相关路由
	RegisterApiRouter(engine)

}
