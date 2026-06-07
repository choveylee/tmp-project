package router

import (
	"github.com/gin-gonic/gin"

	"dev.choveylee.top/elder-care-backend/internal/handler"
)

func registerUserAdmin(router *gin.RouterGroup) {
	router.GET("/users", handler.HandleListUsersAdmin)

	router.POST("/users", handler.HandleCreateUserAdmin)
	router.PUT("/users/:id", handler.HandleUpdateUserAdmin)

	router.PUT("/users/:id/password", handler.HandleUpdateUserPasswordAdmin)
	router.PUT("/users/password", handler.HandleUpdateOwnPasswordAdmin)
}
