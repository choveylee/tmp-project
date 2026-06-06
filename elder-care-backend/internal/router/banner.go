package router

import (
	"github.com/gin-gonic/gin"

	"dev.choveylee.top/elder-care-backend/internal/handler"
)

func registerBannerClient(router *gin.RouterGroup) {
	router.GET("/banners", handler.HandleListBannersClient)
}

func registerBannerAdmin(router *gin.RouterGroup) {
	router.GET("/banners", handler.HandleListBannersAdmin)

	router.POST("/banners", handler.HandleCreateBannerAdmin)
	router.GET("/banners/:id", handler.HandleGetBannerAdmin)
	router.PUT("/banners/:id", handler.HandleUpdateBannerAdmin)
	router.DELETE("/banners/:id", handler.HandleDeleteBannerAdmin)
}
