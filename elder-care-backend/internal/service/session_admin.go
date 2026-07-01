package service

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/data"
	"dev.choveylee.top/elder-care-backend/internal/lib"
	dbmodel "dev.choveylee.top/elder-care-backend/internal/model/mysql"
	redmodel "dev.choveylee.top/elder-care-backend/internal/model/redis"
)

func CreateSessionAdmin(ctx context.Context, realIp string, mobile, password string, captchaId, captchaCode string) (*data.CreateSessionAdminRespData, *terror.Terror) {
	if runMode != constant.RunModeDebug {
		isMatch := captcha.Store.Verify(captchaId, captchaCode, true)
		if !isMatch {
			errMsg := tlog.E(ctx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (captcha code invalid)",
				realIp, mobile, password, captchaId, captchaCode)

			errx := terror.NewTerror(ctx, terror.ErrParamInvalid("captcha code"), constant.ErrorCodeCaptchaInvalid, errMsg)

			return nil, errx
		}
	}

	userDB, errx := dbmodel.FindUserByMobile(ctx, mobile)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (db find user by mobile %v)",
			realIp, mobile, password, captchaId, captchaCode, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if userDB == nil {
		errMsg := tlog.E(ctx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (mobile not exist)",
			realIp, mobile, password, captchaId, captchaCode)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("mobile"), constant.ErrorCodeUserNotExist, errMsg)

		return nil, errx
	}

	roleId := userDB.RoleId

	roleDB, errx := dbmodel.FindRole(ctx, roleId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s, role id: %s) err (db find role %v)",
			realIp, mobile, password, captchaId, captchaCode, roleId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if roleDB == nil {
		errMsg := tlog.E(ctx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s, role id: %s) err (role not exist)",
			realIp, mobile, password, captchaId, captchaCode, roleId)

		errx := terror.NewTerror(ctx, terror.ErrSvcAbnormal("role id"), constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return nil, errx
	}

	// 校验密码重试次数
	retryCount, errx := redmodel.FindPasswordRetry(ctx, mobile)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (red find captcha retry %v)",
			realIp, mobile, password, captchaId, captchaCode, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if retryCount >= constant.PasswordRetryCount {
		errMsg := tlog.E(ctx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (password retry limit)",
			realIp, mobile, password, captchaId, captchaCode)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("retry limit"), constant.ErrorCodePasswordRetryLimit, errMsg)

		return nil, errx
	}

	srcPassword := userDB.Password

	err := bcrypt.CompareHashAndPassword([]byte(srcPassword), []byte(password))
	if err != nil {
		errx = redmodel.UpdatePasswordRetry(ctx, mobile)
		if errx != nil {
			tlog.E(ctx).Err(errx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (red update password retry %v)",
				realIp, mobile, password, captchaId, captchaCode, errx)
		}

		errMsg := tlog.E(ctx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (password not match)",
			realIp, mobile, password, captchaId, captchaCode)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("password"), constant.ErrorCodePasswordNotMatch, errMsg)

		return nil, errx
	}

	errx = redmodel.DeleteCaptcha(ctx, mobile)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (red delete captcha %v)",
			realIp, mobile, password, captchaId, captchaCode, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	status := userDB.Status
	if status != dbmodel.UserStatusNormal {
		errMsg := tlog.E(ctx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (status invalid)",
			realIp, mobile, password, captchaId, captchaCode)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("user status"), constant.ErrorCodeUserStatusInvalid, errMsg)

		return nil, errx
	}

	userId := userDB.Id

	isAdmin := roleDB.IsAdmin

	curTime := time.Now()

	accessToken, errx := lib.CreateJwtAccessToken(ctx, userId, roleId, isAdmin, curTime)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s, user id: %s) err (create jwt access token %v)",
			realIp, mobile, password, captchaId, captchaCode, userId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	err = dbmodel.DB(ctx).Transaction(func(tx *gorm.DB) error {
		errx = dbmodel.UpdateUserLoginAt(ctx, tx, userId, curTime)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (db update user login at %v)",
				realIp, mobile, password, captchaId, captchaCode, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		_, errx = dbmodel.CreateAccessLog(ctx, tx, userId, "", realIp)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (db create access log %v)",
				realIp, mobile, password, captchaId, captchaCode, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		return nil
	})
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Create session admin (real ip: %s, mobile: %s, password: %s, captcha id: %s, captcha code: %s) err (db transaction %v)",
			realIp, mobile, password, captchaId, captchaCode, err)

		errx = terror.NewTerror(ctx, err, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	createSessionRespData := &data.CreateSessionAdminRespData{
		AccessToken: accessToken,

		ExpiresIn: lib.JwtAccessTokenTtl,

		/*
			User: &data.UserBriefAdminData{
				UserId: userDB.Id,

				Name:   userDB.Name,
				Mobile: userDB.Mobile,
			},
		*/
	}

	return createSessionRespData, nil
}
