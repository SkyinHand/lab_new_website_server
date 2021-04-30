package router

import (
	"github.com/gin-gonic/gin"
	"lib_new_website_server/app/controller"
)

// 注册API相关路由, engine使用指针传递, 无需返回

func RegisterApiRouter(engine *gin.Engine) {
	apiRouter := engine.Group("/api/v1")
	{
		// 用户相关路由
		userRouter := apiRouter.Group("/user")
		{
			// 添加新成员，报名
			userRouter.POST("/post", controller.PostNewCustomer)
			userRouter.GET(":userID", controller.QueryCustomer)
		}
	}
}
