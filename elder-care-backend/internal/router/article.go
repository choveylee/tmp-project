package router

import (
	"github.com/gin-gonic/gin"

	"dev.choveylee.top/elder-care-backend/internal/handler"
)

func registerArticleClient(router *gin.RouterGroup) {
	router.GET("/articles/categories", handler.HandleListArticleCategoriesClient)

	router.GET("/articles", handler.HandleListArticlesClient)

	router.GET("/articles/:id", handler.HandleGetArticleClient)
}

func registerArticleAdmin(router *gin.RouterGroup) {
	router.GET("/articles/categories", handler.HandleListArticleCategoriesAdmin)

	router.POST("/articles/categories", handler.HandleCreateArticleCategoryAdmin)
	router.GET("/articles/categories/:id", handler.HandleGetArticleCategoryAdmin)
	router.PUT("/articles/categories/:id", handler.HandleUpdateArticleCategoryAdmin)
	router.DELETE("/articles/categories/:id", handler.HandleDeleteArticleCategoryAdmin)

	router.GET("/articles", handler.HandleListArticlesAdmin)

	router.POST("/articles", handler.HandleCreateArticleAdmin)
	router.GET("/articles/:id", handler.HandleGetArticleAdmin)
	router.PUT("/articles/:id", handler.HandleUpdateArticleAdmin)
	router.DELETE("/articles/:id", handler.HandleDeleteArticleAdmin)
}
