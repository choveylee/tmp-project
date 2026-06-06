// Package middleware provides HTTP middleware used by the service router.
package middleware

import (
	"strings"

	"github.com/choveylee/tlog"
	"github.com/gin-gonic/gin"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/handler"
	"dev.choveylee.top/elder-care-backend/internal/lib"
	dbmodel "dev.choveylee.top/elder-care-backend/internal/model/mysql"
)

// AdminAuthMiddleware returns a placeholder authentication middleware.
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		accessToken := strings.TrimSpace(c.Request.Header.Get("Authorization"))
		if accessToken == "" {
			errMsg := tlog.E(ctx).Msg("Admin auth middleware err (access token missing)")

			handler.SendFailResponse(c, constant.ErrorCodeAccessTokenInvalid, errMsg)

			c.Abort()
			return
		}

		if strings.HasPrefix(strings.ToLower(accessToken), "bearer ") {
			accessToken = strings.TrimSpace(accessToken[len("Bearer "):])
		}

		if accessToken == "" {
			errMsg := tlog.E(ctx).Msg("Admin auth middleware err (access token missing)")

			handler.SendFailResponse(c, constant.ErrorCodeAccessTokenInvalid, errMsg)

			c.Abort()
			return
		}

		userId, roleId, isAdmin, errx := lib.CheckJwtAccessToken(ctx, accessToken)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Admin auth middleware err (check jwt access token %v)",
				errx)

			handler.SendFailResponse(c, errx.ErrCode(), errMsg)

			c.Abort()
			return
		}

		if !isAdmin {
			errMsg := tlog.E(ctx).Msgf("Admin auth middleware (user id: %s) err (permission forbidden)",
				userId)

			handler.SendFailResponse(c, constant.ErrorCodePermissionForbidden, errMsg)

			c.Abort()
			return
		}

		path := strings.TrimSpace(c.FullPath())
		if path == "" {
			path = strings.TrimSpace(c.Request.URL.Path)
		}

		method := strings.TrimSpace(strings.ToUpper(c.Request.Method))

		// 用于判断接口是否已注册
		apiResourceDB, errx := dbmodel.FindApiResourceByPath(ctx, method, path)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Admin auth middleware (method: %s, path: %s) err (db find api resource %v)",
				c.Request.Method, path, errx)
			errx.AttachErrMsg(errMsg)

			handler.SendFailResponse(c, errx.ErrCode(), errMsg)

			c.Abort()
			return
		}

		if apiResourceDB == nil || apiResourceDB.AuthRequired != dbmodel.IsValueTrue {
			errMsg := tlog.E(ctx).Msgf("Admin auth middleware (method: %s, path: %s) err (api resource unauthorized)",
				c.Request.Method, path)

			handler.SendFailResponse(c, constant.ErrorCodePermissionForbidden, errMsg)

			c.Abort()
			return
		}

		// 用于判断接口是否有权限
		apiResource2DB, errx := dbmodel.FindApiResourceByRole(ctx, roleId, method, path)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Admin auth middleware (role id: %s, method: %s, path: %s) err (db find authorized api resource %v)",
				roleId, c.Request.Method, path, errx)
			errx.AttachErrMsg(errMsg)

			handler.SendFailResponse(c, errx.ErrCode(), errMsg)

			c.Abort()
			return
		}

		if apiResource2DB == nil {
			errMsg := tlog.E(ctx).Msgf("Admin auth middleware (role id: %s, method: %s, path: %s) err (permission forbidden)",
				roleId, c.Request.Method, path)

			handler.SendFailResponse(c, constant.ErrorCodePermissionForbidden, errMsg)

			c.Abort()
			return
		}

		c.Request = c.Request.WithContext(ctx)

		c.Request.Header.Set("user_id", userId)
		c.Request.Header.Set("role_id", roleId)

		c.Set("user_id", userId)
		c.Set("role_id", roleId)

		c.Next()
	}
}
