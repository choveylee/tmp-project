/**
 * @Author: lidonglin
 * @Description:
 * @File:  jwt.go
 * @Version: 1.0.0
 * @Date: 2024/7/18 09:15:59
 */

package lib

import (
	"context"

	"github.com/choveylee/tcfg"
	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
)

var (
	jwtTokenSecret string
)

func initJwt(ctx context.Context) *terror.Terror {
	jwtTokenSecret = tcfg.DefaultString(tcfg.LocalKey("JWT_TOKEN_SECRET"), "")
	if jwtTokenSecret == "" {
		errMsg := tlog.E(ctx).Msgf("Init jwt token err (jwt token secret invalid)")

		errx := terror.NewRawTerror(ctx, terror.ErrConfInvalid("jwt token secret"), errMsg)

		return errx
	}

	return nil
}
