package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRouterController(context *gin.Context) {
	context.HTML(http.StatusOK, "404.html", gin.H{})
}
