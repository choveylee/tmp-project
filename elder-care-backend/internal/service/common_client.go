package service

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"github.com/mojocn/base64Captcha"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/data"
	redmodel "dev.choveylee.top/elder-care-backend/internal/model/redis"
)

var (
	store = &CaptchaStore{}

	driverDigit = &base64Captcha.DriverDigit{
		Height:   80,
		Width:    240,
		Length:   5,
		MaxSkew:  0.7,
		DotCount: 80,
	}

	captcha *base64Captcha.Captcha
)

func init() {
	rand.NewSource(time.Now().UnixNano())

	captcha = base64Captcha.NewCaptcha(driverDigit, store)
}

type CaptchaStore struct {
}

func (store CaptchaStore) Set(id string, value string) error {
	ctx := context.Background()

	errx := redmodel.CreateCaptcha(ctx, id, value)
	if errx != nil {
		tlog.E(ctx).Err(errx).Msgf("Captcha store set (id: %s, value: %s) err (red create captcha %v)",
			id, value, errx)

		return errx
	}

	return nil
}

func (store CaptchaStore) Get(id string, clear bool) string {
	ctx := context.Background()

	value, errx := redmodel.FindCaptcha(ctx, id)
	if errx != nil {
		tlog.E(ctx).Err(errx).Msgf("Captcha store get (id: %s, clear: %t) err (red find captcha %v)",
			id, clear, errx)

		return ""
	}

	if clear {
		errx := redmodel.DeleteCaptcha(ctx, id)
		if errx != nil {
			tlog.E(ctx).Err(errx).Msgf("Captcha store get (id: %s, clear: %t) err (red delete captcha %v)",
				id, clear, errx)

			return ""
		}
	}

	return value
}

func (store CaptchaStore) Verify(id, answer string, clear bool) bool {
	value := store.Get(id, clear)

	value = strings.TrimSpace(value)

	return value == strings.TrimSpace(answer)
}

func CreateCaptchaClient(ctx context.Context) (*data.CreateCaptchaClientRespData, *terror.Terror) {
	captchaId, captchaImg, answer, err := captcha.Generate()
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Create captcha client err (captcha generate %v)",
			err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return nil, errx
	}

	tlog.I(ctx).Msgf("Create captcha client (captcha id: %s, answer: %s) success",
		captchaId, answer)

	createCaptchaRespData := &data.CreateCaptchaClientRespData{
		CaptchaId:  captchaId,
		CaptchaImg: captchaImg,
	}

	return createCaptchaRespData, nil
}
