package router

import (
	"github.com/gin-gonic/gin"

	"dev.choveylee.top/elder-care-backend/internal/handler"
)

func registerCommonClient(router *gin.RouterGroup) {
	router.POST("/captcha", handler.HandleCreateCaptchaClient)
}

func registerCommonAdmin(router *gin.RouterGroup) {
	router.POST("/captcha", handler.HandleCreateCaptchaClient)

	router.POST("/image", handler.HandleCreateImageAdmin)
}
