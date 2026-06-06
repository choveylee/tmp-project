package handler

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/choveylee/thttp"
	"github.com/choveylee/tlog"
	"github.com/gin-gonic/gin"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/data"
	"dev.choveylee.top/elder-care-backend/internal/service"
)

func HandleCreateSessionAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	realIp := thttp.GetRealIP(c.Request)

	createSessionRequest := &data.CreateSessionAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, createSessionRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle create session admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	mobile := strings.TrimSpace(createSessionRequest.Mobile)
	if mobile == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create session admin err (mobile invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	retMatches := constant.ChnMobileReg.FindAllString(mobile, -1)
	if len(retMatches) == 0 {
		errMsg := tlog.E(ctx).Msgf("Handle create session admin (mobile: %s) err (mobile invalid)",
			mobile)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	password := strings.TrimSpace(createSessionRequest.Password)
	if password == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create session admin err (password invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	captchaId := strings.TrimSpace(createSessionRequest.CaptchaId)
	if captchaId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create session admin err (captcha id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	captchaCode := strings.TrimSpace(createSessionRequest.CaptchaCode)
	if captchaCode == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create session admin err (captcha code invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	createSessionRespData, errx := service.CreateSessionAdmin(ctx, realIp, mobile, password, captchaId, captchaCode)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle create token admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (create token admin %v)",
			realIp, mobile, password, captchaId, captchaCode, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, createSessionRespData)
}
