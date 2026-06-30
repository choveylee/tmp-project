package service

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/data"
	dbmodel "dev.choveylee.top/elder-care-backend/internal/model/mysql"
)

func ListCourseCategoriesClient(ctx context.Context) (*data.ListCourseCategoriesClientRespData, *terror.Terror) {
	_, courseCategoriesDB, errx := dbmodel.FindCourseCategories(ctx, dbmodel.CourseCategoryStatusNormal, -1, -1)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course categories client err (db find course categories %v)",
			errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listCourseCategoriesRespData := &data.ListCourseCategoriesClientRespData{
		Categories: make([]*data.CourseCategoryClientData, 0),
	}

	for _, courseCategoryDB := range courseCategoriesDB {
		courseCategoryData := &data.CourseCategoryClientData{
			CategoryId: courseCategoryDB.Id,

			Name: courseCategoryDB.Name,
		}

		listCourseCategoriesRespData.Categories = append(listCourseCategoriesRespData.Categories, courseCategoryData)
	}

	return listCourseCategoriesRespData, nil
}

func ListCoursesClient(ctx context.Context, categoryId string, courseType int, sortBy string, pageNum, pageSize int) (*data.ListCoursesClientRespData, *terror.Terror) {
	total, coursesDB, errx := dbmodel.FindCourses(ctx, categoryId, courseType, dbmodel.CourseStatusNormal, sortBy, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List courses client (category id: %s, course type: %d, sort by: %s, page num: %d, page size: %d) err (db find courses %v)",
			categoryId, courseType, sortBy, pageNum, pageSize, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listCoursesRespData := &data.ListCoursesClientRespData{
		Courses: make([]*data.CourseClientData, 0),

		Total: total,
	}

	for _, courseDB := range coursesDB {
		courseData := &data.CourseClientData{
			CourseId: courseDB.Id,

			Title: courseDB.Title,

			Abstract: courseDB.Abstract,

			CoverUrl: courseDB.CoverUrl,
			LinkUrl:  courseDB.LinkUrl,

			PublishAt: courseDB.PublishAt.Format(time.RFC3339),
		}

		listCoursesRespData.Courses = append(listCoursesRespData.Courses, courseData)
	}

	return listCoursesRespData, nil
}

func GetCourseClient(ctx context.Context, courseId string) (*data.GetCourseClientRespData, *terror.Terror) {
	courseDB, errx := dbmodel.FindCourse(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get course client (course id: %s) err (db find course %v)",
			courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseDB == nil {
		return nil, nil
	}

	if courseDB.Status != dbmodel.CourseStatusNormal {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get course client (course id: %s) err (course status invalid)",
			courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseInvalid, errMsg)

		return nil, errx
	}

	getCourseRespData := &data.GetCourseClientRespData{
		CourseId: courseDB.Id,

		Author: courseDB.Author,
		Source: courseDB.Source,

		Title: courseDB.Title,

		Tags: make([]string, 0),

		Abstract: courseDB.Abstract,

		CoverUrl: courseDB.CoverUrl,
		LinkUrl:  courseDB.LinkUrl,

		FavouriteCount: courseDB.FavouriteCount,
		ViewCount:      courseDB.ViewCount,

		PublishAt: courseDB.PublishAt.Format(time.RFC3339),
	}

	courseTagsDB, errx := dbmodel.FindCourseTags(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get course client (course id: %s) err (db find course tags %v)",
			courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	for _, courseTagDB := range courseTagsDB {
		getCourseRespData.Tags = append(getCourseRespData.Tags, courseTagDB.Name)
	}

	courseDetail, errx := dbmodel.FindCourseDetail(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get course client (course id: %s) err (db find course detail %v)",
			courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseDetail != nil {
		getCourseRespData.Detail = courseDetail.Detail

		getCourseRespData.Summary = courseDetail.Summary
		getCourseRespData.Objective = courseDetail.Objective
		getCourseRespData.Outline = courseDetail.Outline
		getCourseRespData.References = courseDetail.References
	}

	return getCourseRespData, nil
}

func ListCourseVideosClient(ctx context.Context, courseId string) (*data.ListCourseVideosClientRespData, *terror.Terror) {
	courseDB, errx := dbmodel.FindCourse(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course videos client (course id: %s) err (db find course %v)",
			courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseDB == nil {
		return nil, nil
	}

	if courseDB.Status != dbmodel.CourseStatusNormal {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course videos client (course id: %s) err (course status invalid)",
			courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseInvalid, errMsg)

		return nil, errx
	}

	if courseDB.CourseType != dbmodel.CourseTypeVideo {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course videos client (course id: %s) err (course type invalid)",
			courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseTypeInvalid, errMsg)

		return nil, errx
	}

	courseVideosDB, errx := dbmodel.FindCourseVideos(ctx, courseId, dbmodel.CourseVideoStatusNormal)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course videos client (course id: %s) err (db find course videos %v)",
			courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listCourseVideosClientRespData := &data.ListCourseVideosClientRespData{
		Videos: make([]*data.CourseVideoClientData, 0),
	}

	for _, courseVideoDB := range courseVideosDB {
		courseVideoData := &data.CourseVideoClientData{
			VideoId: courseVideoDB.Id,

			VideoUrl: courseVideoDB.VideoUrl,

			Duration: courseVideoDB.Duration,
			Format:   courseVideoDB.Format,
			Language: courseVideoDB.Language,
			Size:     courseVideoDB.Size,

			UploadAt: courseVideoDB.UploadAt.Format(time.RFC3339),
		}

		listCourseVideosClientRespData.Videos = append(listCourseVideosClientRespData.Videos, courseVideoData)
	}

	return listCourseVideosClientRespData, nil
}
