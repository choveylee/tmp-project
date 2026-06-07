// Package router configures the HTTP routing used by the generated service.
package router

import (
	"context"

	"github.com/choveylee/tcfg"
	"github.com/choveylee/tserver"
	tmiddleware "github.com/choveylee/tserver/middleware"
	"github.com/gin-gonic/gin"

	"dev.choveylee.top/elder-care-backend/internal/router/middleware"
)

// NewRouter constructs the service HTTP router and registers application routes.
func NewRouter(ctx context.Context) *gin.Engine {
	appName := tcfg.DefaultString("APP_NAME", "unknown")

	router := tserver.NewRouter(appName)

	router.Use(tmiddleware.CorsMiddleware())

	// Register monitoring routes.
	registerMonitor(router)

	clientRouter := router.Group("/api/v1/public")

	registerCommonClient(clientRouter)

	registerBannerClient(clientRouter)
	registerArticleClient(clientRouter)

	clientAuthRouter := router.Group("/api/v1/public")
	clientAuthRouter.Use(middleware.ClientAuthMiddleware())

	adminRouter := router.Group("/api/v1/admin")

	registerCommonAdmin(adminRouter)
	registerSessionAdmin(adminRouter)

	adminAuthRouter := router.Group("/api/v1/admin")
	adminAuthRouter.Use(middleware.AdminAuthMiddleware())

	registerUserAdmin(adminAuthRouter)

	registerBannerAdmin(adminAuthRouter)
	registerArticleAdmin(adminAuthRouter)

	return router
}
