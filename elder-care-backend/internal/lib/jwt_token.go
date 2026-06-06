/**
 * @Author: lidonglin
 * @Description:
 * @File:  jwt_token.go
 * @Version: 1.0.0
 * @Date: 2024/7/18 09:16:06
 */

package lib

import (
	"context"
	"errors"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"github.com/golang-jwt/jwt/v4"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

const (
	JwtAccessTokenTtl = 60 * 60 * 48
)

type CustomClaims struct {
	jwt.RegisteredClaims

	UserId string `json:"user_id"`
	RoleId string `json:"role_id"`

	IsAdmin bool `json:"is_admin"`
}

// CreateJwtAccessToken 生成access token
func CreateJwtAccessToken(ctx context.Context, userId string, roleId string, isAdmin bool, issueTime time.Time) (string, *terror.Terror) {
	customClaims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: userId,

			ExpiresAt: jwt.NewNumericDate(issueTime.Add(time.Second * JwtAccessTokenTtl)),
		},

		UserId: userId,
		RoleId: roleId,

		IsAdmin: isAdmin,
	}

	// 1. 生成jwt的RS256签名
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	// 2. 生成access token
	accessToken, err := tokenClaims.SignedString([]byte(jwtTokenSecret))
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Create jwt access token (user id: %s, role id: %s) err (signed string %v)",
			userId, roleId, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return "", errx
	}

	return accessToken, nil
}

// CheckJwtAccessToken 校验access token
func CheckJwtAccessToken(ctx context.Context, accessToken string) (string, string, bool, *terror.Terror) {
	tokenClaims, err := jwt.ParseWithClaims(accessToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtTokenSecret), nil
	})
	if err != nil {
		var validationErr *jwt.ValidationError

		if errors.As(err, &validationErr) {
			if validationErr.Errors&(jwt.ValidationErrorExpired) != 0 {
				// access token 过期, 重新生成token
				errMsg := tlog.E(ctx).Msgf("Check jwt access token (access token: %s) err (access token expired)",
					accessToken)

				errx := terror.NewTerror(ctx, terror.ErrParamInvalid("access token"), constant.ErrorCodeAccessTokenExpired, errMsg)

				return "", "", false, errx
			}
		}

		errMsg := tlog.E(ctx).Err(err).Msgf("Check jwt access token (access token: %s) err (parse with claims %v)",
			accessToken, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeAccessTokenInvalid, errMsg)

		return "", "", false, errx
	}

	if tokenClaims == nil {
		errMsg := tlog.E(ctx).Msgf("Check jwt access token (access token: %s) err (token claims illegal)",
			accessToken)

		errx := terror.NewTerror(ctx, terror.ErrSvcAbnormal("token claims"), constant.ErrorCodeAccessTokenInvalid, errMsg)

		return "", "", false, errx
	}

	customClaims, ok := tokenClaims.Claims.(*CustomClaims)
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Check jwt access token (access token: %s) err (access token illegal)", accessToken)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("access token"), constant.ErrorCodeAccessTokenInvalid, errMsg)

		return "", "", false, errx
	}

	if !tokenClaims.Valid {
		errMsg := tlog.E(ctx).Msgf("Check jwt access token (access token: %s) err (access token illegal)", accessToken)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("access token"), constant.ErrorCodeAccessTokenInvalid, errMsg)

		return "", "", false, errx
	}

	return customClaims.UserId, customClaims.RoleId, customClaims.IsAdmin, nil
}
