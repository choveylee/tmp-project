package handler

import (
	"github.com/choveylee/tlog"
	"github.com/gin-gonic/gin"

	"dev.choveylee.top/elder-care-backend/internal/service"
)

func HandleListBannersClient(c *gin.Context) {
	ctx := c.Request.Context()

	listBannersRespData, errx := service.ListBannersClient(ctx)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list banners client err (list banners client %v)",
			errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, listBannersRespData)
}
