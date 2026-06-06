package router

import (
	"github.com/gin-gonic/gin"

	"dev.choveylee.top/elder-care-backend/internal/handler"
)

func registerSessionAdmin(router *gin.RouterGroup) {
	router.POST("/sessions", handler.HandleCreateSessionAdmin)
}
