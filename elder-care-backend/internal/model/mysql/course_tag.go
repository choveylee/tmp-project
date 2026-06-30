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

const (
	CourseTagNameLenLimit = 1024
)

type CourseTag struct {
	Id string

	CourseId string
	Name     string

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CreateCourseTags(ctx context.Context, tx *gorm.DB, courseId string, names []string) (*CourseTag, *terror.Terror) {
	courseTagsDB := make([]*CourseTag, 0)

	for _, name := range names {
		courseTagDB := &CourseTag{
			Id: tutil.NewOid().String(),

			CourseId: courseId,
			Name:     name,
		}

		courseTagsDB = append(courseTagsDB, courseTagDB)
	}

	retGorm := tx.Create(courseTagsDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Create course tags (course id: %s, names: %v) err (db create %v)",
			courseId, names, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}
}

func FindCourseTags(ctx context.Context, courseId string) ([]*CourseTag, *terror.Terror) {
	courseTagsDB := make([]*CourseTag, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", courseId).Find(&courseTagsDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course tags (course id: %s) err (db find %v)",
			courseId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return courseTagsDB, nil
}
