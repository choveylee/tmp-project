package dbmodel

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"github.com/choveylee/tutil"
	"gorm.io/gorm"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

type AccessLog struct {
	Id string

	UserId string

	Lang   string
	RealIp string

	CreatedAt time.Time
}

func CreateAccessLog(ctx context.Context, tx *gorm.DB, userId string, lang, realIp string) (*AccessLog, *terror.Terror) {
	accessLogDB := &AccessLog{
		Id: tutil.NewOid().String(),

		UserId: userId,

		Lang:   lang,
		RealIp: realIp,
	}

	retGorm := tx.Create(accessLogDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Create access log (user id: %s, lang: %s, real ip: %s) err (db create %v).",
			userId, lang, realIp, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return accessLogDB, nil
}
