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

	CourseId string

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

func CreateCourseVideo(ctx context.Context, courseId string, videoUrl string, format, language, size, duration string, uploadAt time.Time, weight, status int) (*CourseVideo, *terror.Terror) {
	courseVideoDB := &CourseVideo{
		Id: tutil.NewOid().String(),

		CourseId: courseId,

		VideoUrl: videoUrl,

		Format:   format,
		Language: language,
		Size:     size,
		Duration: duration,

		UploadAt: uploadAt,

		Weight: weight,
		Status: status,
	}

	retGorm := serverClient.DB(ctx, runMode).Create(courseVideoDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Create course video (course id: %s, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %v, weight: %d, status: %d) err (db create %v)",
			courseId, videoUrl, format, language, size, duration, uploadAt, weight, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return courseVideoDB, nil
}

func FindCourseVideo(ctx context.Context, videoId string) (*CourseVideo, *terror.Terror) {
	courseVideosDB := make([]*CourseVideo, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", videoId).Limit(1).Find(&courseVideosDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course video (video id: %s) err (db find %v)",
			videoId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(courseVideosDB) == 0 {
		return nil, nil
	}

	return courseVideosDB[0], nil
}

func FindCourseVideos(ctx context.Context, courseId string, status int) ([]*CourseVideo, *terror.Terror) {
	courseVideosDB := make([]*CourseVideo, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("course_id = ?", courseId)

	if status != -1 {
		retGorm = retGorm.Where("status = ?", status)
	}

	retGorm = retGorm.Order("weight ASC, created_at DESC").Find(&courseVideosDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course videos (course id: %s, status: %d) err (db find %v)",
			courseId, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return courseVideosDB, nil
}

func UpdateCourseVideo(ctx context.Context, videoId string, videoUrl string, format, language, size, duration string, uploadAt time.Time, weight, status int) *terror.Terror {
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

	retGorm := serverClient.DB(ctx, runMode).Model(&CourseVideo{}).Where("id = ?", videoId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update course video (video id: %s, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %v, weight: %d, status: %d) err (db updates %v)",
			videoId, videoUrl, format, language, size, duration, uploadAt, weight, status, retGorm.Error)
		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteCourseVideo(ctx context.Context, videoId string) *terror.Terror {
	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", videoId).Delete(&CourseVideo{})
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Delete course video (video id: %s) err (db delete %v)",
			videoId, retGorm.Error)
		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}
