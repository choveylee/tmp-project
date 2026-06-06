package handler

import (
	"github.com/choveylee/tlog"
	"github.com/gin-gonic/gin"

	"dev.choveylee.top/elder-care-backend/internal/service"
)

func HandleCreateCaptchaClient(c *gin.Context) {
	ctx := c.Request.Context()

	createCaptchaRespData, errx := service.CreateCaptchaClient(ctx)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle create captcha client err (create captcha client %v).",
			errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, createCaptchaRespData)
}
