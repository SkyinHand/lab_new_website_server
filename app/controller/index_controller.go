package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndexWebPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

func GetNewsPage(context *gin.Context) {
	context.HTML(http.StatusOK, "signup.html", gin.H{})
}
