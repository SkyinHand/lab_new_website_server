package core

import (
	"github.com/gin-gonic/gin"
	"lib_new_website_server/app/database"
	"lib_new_website_server/app/middleware"
	"lib_new_website_server/app/router"
)

// 获取一个标准的引擎
func GetDefaultEngine() *gin.Engine {
	// 获得一个标准的gin引擎
	engine := gin.Default()
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
