package handler

import (
	"path"

	"github.com/choveylee/tlog"
	"github.com/gin-gonic/gin"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/service"
)

func HandleCreateImageAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	file, err := c.FormFile("file")
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle create image admin err (form file %v).",
			err)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	filename := file.Filename

	fileExt := path.Ext(filename)
	if fileExt != ".img" && fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		errMsg := tlog.E(ctx).Msgf("Handle create image admin (file ext: %s) err (file ext invalid).",
			filename)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	createImageRespData, errx := service.CreateImageAdmin(ctx, userId, c, file, filename, fileExt)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle create image admin (user id: %s, filename: %s, file ext: %s) err (create image admin %v).",
			userId, filename, fileExt, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, createImageRespData)
}
