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
	CourseVideoUrlLen      = 1024
	CourseVideoFormatLen   = 32
	CourseVideoLanguageLen = 32
	CourseVideoSizeLen     = 32
	CourseVideoDurationLen = 32
)

const (
	CourseVideoStatusNormal   = 0
	CourseVideoStatusDisabled = 1
)

var (
	CourseVideoStatusesMap = map[int]bool{
		CourseVideoStatusNormal:   true,
		CourseVideoStatusDisabled: true,
	}
)

type CourseVideo struct {
	Id string

	CatalogId string

	VideoUrl string

	Format   string
	Language string
	Size     string
	Duration string

	UploadAt time.Time

	Weight int
	Status int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CreateCourseVideo(ctx context.Context, tx *gorm.DB, catalogId string, videoUrl string, format, language, size, duration string, uploadAt time.Time, weight, status int) (*CourseVideo, *terror.Terror) {
	courseVideoDB := &CourseVideo{
		Id: tutil.NewOid().String(),

		CatalogId: catalogId,

		VideoUrl: videoUrl,

		Format:   format,
		Language: language,
		Size:     size,
		Duration: duration,

		UploadAt: uploadAt,

		Weight: weight,
		Status: status,
	}

	retGorm := tx.Create(courseVideoDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Create course video (catalog id: %s, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %v, weight: %d, status: %d) err (db create %v)",
			catalogId, videoUrl, format, language, size, duration, uploadAt, weight, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return courseVideoDB, nil
}

func FindCourseVideoByCatalog(ctx context.Context, catalogId string) (*CourseVideo, *terror.Terror) {
	courseVideosDB := make([]*CourseVideo, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("catalog_id = ?", catalogId).Limit(1).Find(&courseVideosDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course video by catalog (catalog id: %s) err (db find %v)",
			catalogId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(courseVideosDB) == 0 {
		return nil, nil
	}

	return courseVideosDB[0], nil
}

func FindCourseVideosByCatalog(ctx context.Context, catalogIds []string, status int) ([]*CourseVideo, *terror.Terror) {
	courseVideosDB := make([]*CourseVideo, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("catalog_id IN (?)", catalogIds)

	if status != -1 {
		retGorm = retGorm.Where("status = ?", status)
	}

	retGorm = retGorm.Order("weight ASC, created_at DESC").Find(&courseVideosDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course videos by catalog (catalog ids: %v, status: %d) err (db find %v)",
			catalogIds, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return courseVideosDB, nil
}

func UpdateCourseVideo(ctx context.Context, tx *gorm.DB, videoId string, videoUrl string, format, language, size, duration string, uploadAt time.Time, weight, status int) *terror.Terror {
	params := map[string]interface{}{
		"video_url": videoUrl,

		"format":   format,
		"language": language,
		"size":     size,

		"duration":  duration,
		"upload_at": uploadAt,

		"weight": weight,
		"status": status,

		"updated_at": time.Now(),
	}

	retGorm := tx.Model(&CourseVideo{}).Where("id = ?", videoId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update course video (video id: %s, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %v, weight: %d, status: %d) err (db updates %v)",
			videoId, videoUrl, format, language, size, duration, uploadAt, weight, status, retGorm.Error)
		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteCourseVideo(ctx context.Context, tx *gorm.DB, videoId string) *terror.Terror {
	retGorm := tx.Where("id = ?", videoId).Delete(&CourseVideo{})
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Delete course video (video id: %s) err (db delete %v)",
			videoId, retGorm.Error)
		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteCourseVideosByCatalog(ctx context.Context, tx *gorm.DB, catalogIds []string) *terror.Terror {
	retGorm := tx.Where("catalog_id IN (?)", catalogIds).Delete(&CourseVideo{})
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Delete course videos by catalog (catalog ids: %v) err (db delete %v)",
			catalogIds, retGorm.Error)
		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}
