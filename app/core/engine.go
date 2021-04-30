package core

import (
	"github.com/gin-gonic/gin"
	"lib_new_website_server/app/database"
	"lib_new_website_server/app/middleware"
	"lib_new_website_server/app/router"
	"lib_new_website_server/app/utils"
)

/**
 * @Description: 获取一个Gin引擎
 * @return *gin.Engine
 */
func GetDefaultEngine() *gin.Engine {
	// 是否是Debug
	debug := utils.GetConfig("server", "debug")
	switch debug {
	case "true":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	// 获得一个标准的gin引擎
	engine := gin.New()
	// 注册中间件
	middleware.RegisterAllMiddleware(engine)
	// 加载静态文件
	engine.Static("/static", "app/static")
	// 加载官网页模板
	engine.LoadHTMLGlob("app/views/*")
	// 注册Web路由
	router.RegisterAllRouters(engine)
	// 加载数据库文件
	database.Initdatabase()
	return engine
}
