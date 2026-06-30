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
	CourseAuthorLenLimit = 1024
	CourseSourceLenLimit = 1024

	CourseTitleLenLimit    = 1024
	CourseAbstractLenLimit = 1024

	CourseCoverUrlLenLimit = 1024
	CourseLinkUrlLenLimit  = 1024
)

const (
	CourseSortByPublish   = "publish"
	CourseSortByView      = "view"
	CourseSortByFavourite = "favourite"
)

var (
	CourseSortBysMap = map[string]bool{
		CourseSortByPublish:   true,
		CourseSortByView:      true,
		CourseSortByFavourite: true,
	}
)

const (
	CourseTypeNormal = 0
	CourseTypeVideo  = 1
)

var (
	CourseTypesMap = map[int]bool{
		CourseTypeNormal: true,
		CourseTypeVideo:  true,
	}
)

const (
	CourseStatusNormal   = 0
	CourseStatusDisabled = 1
)

var (
	CourseStatusesMap = map[int]bool{
		CourseStatusNormal:   true,
		CourseStatusDisabled: true,
	}
)

type Course struct {
	Id string

	CategoryId string
	CourseType int

	Author string
	Source string

	Title    string
	Abstract string

	CoverUrl string
	LinkUrl  string

	PublishAt *time.Time

	FavouriteCount int
	ViewCount      int

	Status int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CreateCourse(ctx context.Context, tx *gorm.DB, categoryId string, courseType int, author, source, title, abstract, coverUrl, linkUrl string, publishAt *time.Time, status int) (*Course, *terror.Terror) {
	courseDB := &Course{
		Id: tutil.NewOid().String(),

		CategoryId: categoryId,

		CourseType: courseType,

		Author: author,
		Source: source,

		Title:    title,
		Abstract: abstract,

		CoverUrl: coverUrl,
		LinkUrl:  linkUrl,

		PublishAt: publishAt,

		FavouriteCount: 0,
		ViewCount:      0,

		Status: status,
	}

	retGorm := serverClient.DB(ctx, runMode).Create(courseDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Create course (category id: %s, course type: %d, author: %s, source: %s, title: %s, abstract: %s, cover url: %s, link url: %s, publish at: %v, status: %d) err (db create %v)",
			categoryId, courseType, author, source, title, abstract, coverUrl, linkUrl, publishAt, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return courseDB, nil
}

func FindCourse(ctx context.Context, courseId string) (*Course, *terror.Terror) {
	coursesDB := make([]*Course, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", courseId).Limit(1).Find(&coursesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course (course id: %s) err (db find %v)",
			courseId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(coursesDB) == 0. {
		return nil, nil
	}

	return coursesDB[0], nil
}

func FindCourses(ctx context.Context, categoryId string, courseType int, status int, sortBy string, pageNum, pageSize int) (int64, []*Course, *terror.Terror) {
	query := serverClient.DB(ctx, runMode)

	if categoryId != "" {
		query = query.Where("course_type = ?", courseType)
	}

	if courseType != -1 {
		query = query.Where("course_type = ?", courseType)
	}

	if status != -1 {
		query = query.Where("status = ?", status)
	}

	total := int64(0)

	retGorm := query.Model(&Course{}).Count(&total)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find courses (category id: %s, course type: %d, status: %d, sort by: %s, page num: %d, page size: %d) err (db count %v)",
			categoryId, courseType, status, sortBy, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return -1, nil, errx
	}

	if total == 0 {
		return 0, nil, nil
	}

	coursesDB := make([]*Course, 0)

	retGorm = query

	if pageNum != -1 && pageSize != -1 {
		retGorm = retGorm.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	switch sortBy {
	case CourseSortByPublish:
		retGorm = retGorm.Order("publish_at DESC, created_at DESC")
	case CourseSortByView:
		retGorm = retGorm.Order("view_count DESC, created_at DESC")
	case CourseSortByFavourite:
		retGorm = retGorm.Order("favourite_count DESC, created_at DESC")
	default:
		retGorm = retGorm.Order("publish_at DESC, created_at DESC")
	}

	retGorm = query.Find(&coursesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find courses (category id: %s, course type: %d, status: %d, sort by: %s, page num: %d, page size: %d) err (db find %v)",
			categoryId, courseType, status, sortBy, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return -1, nil, errx
	}

	return total, coursesDB, nil
}

func FindCourseCountByCategory(ctx context.Context, categoryId string) (int64, *terror.Terror) {
	total := int64(0)

	retGorm := serverClient.DB(ctx, runMode).Model(&Course{}).Where("category_id = ?", categoryId).Count(&total)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course count by category (category id: %s) err (db count %v)",
			categoryId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return -1, errx
	}

	return total, nil
}

func UpdateCourse(ctx context.Context, tx *gorm.DB, courseId string, author, source, title, abstract, coverUrl, linkUrl string, publishAt time.Time, status int) *terror.Terror {
	params := map[string]interface{}{
		"author": author,
		"source": source,

		"title":    title,
		"abstract": abstract,

		"cover_url": coverUrl,
		"link_url":  linkUrl,

		"publish_at": publishAt,

		"status": status,
	}

	retGorm := tx.Model(&Course{}).Where("id = ?", courseId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update course (course id: %s, author: %s, source: %s, title: %s, abstract: %s, cover url: %s, link url: %s, publish at: %v, status: %d) err (db update %v)",
			courseId, author, source, title, abstract, coverUrl, linkUrl, publishAt, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}
