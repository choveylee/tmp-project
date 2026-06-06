/**
 * @Author: lidonglin
 * @Description:
 * @File:  password.go
 * @Version: 1.0.0
 * @Date: 2024/8/28 11:57:10
 */

package redmodel

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"github.com/redis/go-redis/v9"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

const (
	PasswordRetryKey = "config.rcrai.com.password.retry.%s"
	PasswordRetryTtl = 24 * time.Hour
)

func FindPasswordRetry(ctx context.Context, mobile string) (int, *terror.Terror) {
	passwordRetryKey := fmt.Sprintf(PasswordRetryKey, mobile)

	retryCount, err := serverClient.Client().Get(ctx, passwordRetryKey).Int()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			errMsg := tlog.E(ctx).Err(err).Msgf("Find password retry (mobile: %s) err (redis get %v).",
				mobile, err)

			errx := terror.NewTerror(ctx, err, constant.ErrorCodeRedisServerAbnormal, errMsg)

			return -1, errx
		}

		return -1, nil
	}

	return retryCount, nil
}

func UpdatePasswordRetry(ctx context.Context, mobile string) *terror.Terror {
	passwordRetryKey := fmt.Sprintf(PasswordRetryKey, mobile)

	err := serverClient.Client().Get(ctx, passwordRetryKey).Err()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			errMsg := tlog.E(ctx).Err(err).Msgf("update password retry (%s, %s) err (redis get %v).",
				mobile, passwordRetryKey, err)

			errx := terror.NewTerror(ctx, err, constant.ErrorCodeRedisServerAbnormal, errMsg)

			return errx
		} else {
			err := serverClient.Client().Set(ctx, passwordRetryKey, 0, PasswordRetryTtl).Err()
			if err != nil {
				errMsg := tlog.E(ctx).Err(err).Msgf("create captcha (%s, %s) err (redis set %v).",
					mobile, passwordRetryKey, err)

				errx := terror.NewTerror(ctx, err, constant.ErrorCodeRedisServerAbnormal, errMsg)

				return errx
			}
		}
	}

	err = serverClient.Client().Incr(ctx, passwordRetryKey).Err()
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("update password retry (%s) err (redis get %v).",
			mobile, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeRedisServerAbnormal, errMsg)

		return errx

	}

	return nil
}

func DeletePasswordRetry(ctx context.Context, mobile string) *terror.Terror {
	passwordRetryKey := fmt.Sprintf(PasswordRetryKey, mobile)

	err := serverClient.Client().Del(ctx, passwordRetryKey).Err()
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("delete password retry (%s) err (redis del %v).",
			mobile, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeRedisServerAbnormal, errMsg)

		return errx
	}

	return nil
}
