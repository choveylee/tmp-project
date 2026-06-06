package redmodel

import (
	"context"
	"fmt"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"github.com/redis/go-redis/v9"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

const (
	CaptchaKey = "lx-one.captcha.v1.%s"
	CaptchaTtl = time.Minute * 60
)

func CreateCaptcha(ctx context.Context, captchaId, captchaCode string) *terror.Terror {
	captchaKey := fmt.Sprintf(CaptchaKey, captchaId)

	err := serverClient.Client().Set(ctx, captchaKey, captchaCode, CaptchaTtl).Err()
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("create captcha (%s, %s) err (redis set %v).",
			captchaId, captchaCode, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeRedisServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func FindCaptcha(ctx context.Context, captchaId string) (string, *terror.Terror) {
	captchaKey := fmt.Sprintf(CaptchaKey, captchaId)

	captchaCode, err := serverClient.Client().Get(ctx, captchaKey).Result()
	if err != nil {
		if err != redis.Nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("find captcha (%s) err (redis get %v).",
				captchaId, err)

			errx := terror.NewTerror(ctx, err, constant.ErrorCodeRedisServerAbnormal, errMsg)

			return "", errx
		}

		return "", nil
	}

	return captchaCode, nil
}

func DeleteCaptcha(ctx context.Context, captchaId string) *terror.Terror {
	captchaKey := fmt.Sprintf(CaptchaKey, captchaId)

	err := serverClient.Client().Del(ctx, captchaKey).Err()
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("delete captcha (%s) err (redis del %v).",
			captchaId, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeRedisServerAbnormal, errMsg)

		return errx
	}

	return nil
}
