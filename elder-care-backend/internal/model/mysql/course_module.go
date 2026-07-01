package dbmodel

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"gorm.io/gorm"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

type CourseModule struct {
	Id string

	Code string
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func FindCourseModuleByCode(ctx context.Context, code string) (*CourseModule, *terror.Terror) {
	courseModulesDB := make([]*CourseModule, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("code = ?", code).Limit(1).Find(&courseModulesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course module by code (code: %s) err (db find %v)",
			code, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(courseModulesDB) == 0 {
		return nil, nil
	}

	return courseModulesDB[0], nil
}
