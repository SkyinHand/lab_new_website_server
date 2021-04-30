package middleware

import "github.com/gin-gonic/gin"

// 注册所有的中间件
func RegisterAllMiddleware(engine *gin.Engine) {
	// 引擎加入中间件
	engine.Use(
		// 跨域中间件
		Cors(),
		// 日志中间件
		Logger(),
		// 恢复现场中间件
		Recover(),
	)
}
