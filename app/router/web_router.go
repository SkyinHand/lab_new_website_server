package router

import (
	"github.com/gin-gonic/gin"
	"lib_new_website_server/app/controller"
)

/**
 * @Description: 注册Web相关路由
 * @param engine
 */
func RegisterWebRouter(engine *gin.Engine) {
	// 获取软件孵化实验室首页内容
	engine.GET("/", controller.GetIndexWebPage)
	engine.GET("/news/", controller.GetNewsPage)
	engine.GET("/news/query/", controller.GetNewsPage)
	engine.GET("/news/about/", controller.GetNewsPage)
	engine.GET("/news/contact/", controller.GetNewsPage)
	engine.GET("/news/new/", controller.GetNewsPage)

	// 未找到路由
	engine.NoRoute(controller.NoRouterController)
}
