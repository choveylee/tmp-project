package dbmodel

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"gorm.io/gorm"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

const (
	CourseDetailLenLimit  = 16777215
	CourseSummaryLenLimit = 16777215

	CourseObjectiveLenLimit  = 16777215
	CourseOutlineLenLimit    = 16777215
	CourseReferencesLenLimit = 16777215
)

type CourseDetail struct {
	Id string

	Detail string

	Summary    string
	Objective  string
	Outline    string
	References string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateCourseDetail(ctx context.Context, tx *gorm.DB, courseId string, detail, summary, objective, outline, references string) (*CourseDetail, *terror.Terror) {
	courseDetailDB := &CourseDetail{
		Id: courseId,

		Detail: detail,

		Summary:    summary,
		Objective:  objective,
		Outline:    outline,
		References: references,
	}

	retGorm := tx.Create(courseDetailDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Create course detail (course id: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s) err (db create %v)",
			courseId, detail, summary, objective, outline, references, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return courseDetailDB, nil
}

func FindCourseDetail(ctx context.Context, courseId string) (*CourseDetail, *terror.Terror) {
	courseDetailsDB := make([]*CourseDetail, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", courseId).Limit(1).Find(&courseDetailsDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course detail (course id: %s) err (db find %v)",
			courseId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(courseDetailsDB) == 0. {
		return nil, nil
	}

	return courseDetailsDB[0], nil
}

func UpdateCourseDetail(ctx context.Context, tx *gorm.DB, courseId string, detail, summary, objective, outline, references string) *terror.Terror {
	params := map[string]interface{}{
		"detail": detail,

		"summary":    summary,
		"objective":  objective,
		"outline":    outline,
		"references": references,

		"updated_at": time.Now(),
	}

	retGorm := tx.Model(&CourseDetail{}).Where("id = ?", courseId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update course detail (course id: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s) err (db update %v)",
			courseId, detail, summary, objective, outline, references, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteCourseDetail(ctx context.Context, tx *gorm.DB, courseId string) *terror.Terror {
	retGorm := tx.Where("id = ?", courseId).Delete(&CourseDetail{})
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Delete course detail (course id: %s) err (db delete %v)",
			courseId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}
