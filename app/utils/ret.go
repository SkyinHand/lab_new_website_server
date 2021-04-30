package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnSuccessJson(context *gin.Context, message string) {
	context.JSON(
		http.StatusOK,
		gin.H {
			"code": 1,
			"message": message,
		},
	)
}

func ReturnErrorJson(context *gin.Context, message string) {
	context.JSON(
		http.StatusOK,
		gin.H {
			"code": 0,
			"message": message,
		},
	)
}

