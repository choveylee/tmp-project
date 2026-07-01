package router

import (
	"github.com/gin-gonic/gin"

	"dev.choveylee.top/elder-care-backend/internal/handler"
)

func registerCourseClient(router *gin.RouterGroup) {
	router.GET("/courses/categories", handler.HandleListCourseCategoriesClient)

	router.GET("/courses", handler.HandleListCoursesClient)

	router.GET("/courses/:id", handler.HandleGetCourseClient)
	router.GET("/courses/:id/catalogs", handler.HandleListCourseCatalogsClient)
}

func registerCourseAdmin(router *gin.RouterGroup) {
	router.GET("/courses/categories", handler.HandleListCourseCategoriesAdmin)

	router.POST("/courses/categories", handler.HandleCreateCourseCategoryAdmin)
	router.GET("/courses/categories/:id", handler.HandleGetCourseCategoryAdmin)
	router.PUT("/courses/categories/:id", handler.HandleUpdateCourseCategoryAdmin)
	router.DELETE("/courses/categories/:id", handler.HandleDeleteCourseCategoryAdmin)

	router.GET("/courses", handler.HandleListCoursesAdmin)

	router.POST("/courses", handler.HandleCreateCourseAdmin)
	router.GET("/courses/:id", handler.HandleGetCourseAdmin)
	router.PUT("/courses/:id", handler.HandleUpdateCourseAdmin)
	router.DELETE("/courses/:id", handler.HandleDeleteCourseAdmin)

	router.GET("/courses/:id/catalogs", handler.HandleListCourseCatalogsAdmin)

	router.POST("/courses/:id/catalogs", handler.HandleCreateCourseCatalogAdmin)
	router.PUT("/courses/catalogs/:id", handler.HandleUpdateCourseCatalogAdmin)
	router.DELETE("/courses/catalogs/:id", handler.HandleDeleteCourseCatalogAdmin)

}
