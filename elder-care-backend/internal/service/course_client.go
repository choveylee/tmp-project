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

func ListCourseCategoriesClient(ctx context.Context, moduleCode string) (*data.ListCourseCategoriesClientRespData, *terror.Terror) {
	courseModuleDB, errx := dbmodel.FindCourseModuleByCode(ctx, moduleCode)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course categories client (module code: %s) err (db find course module %v)",
			moduleCode, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseModuleDB == nil {
		errMsg := tlog.E(ctx).Msgf("List course categories client (module code: %s) err (course module not exist)",
			moduleCode)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("module code"), constant.ErrorCodeCourseModuleNotExist, errMsg)

		return nil, errx
	}

	moduleId := courseModuleDB.Id

	_, courseCategoriesDB, errx := dbmodel.FindCourseCategories(ctx, moduleId, dbmodel.CourseCategoryStatusNormal, -1, -1)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course categories client (module code: %s, module id: %s) err (db find course categories %v)",
			moduleCode, moduleId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listCourseCategoriesRespData := &data.ListCourseCategoriesClientRespData{
		Categories: make([]*data.CourseCategoryClientData, 0),
	}

	for _, courseCategoryDB := range courseCategoriesDB {
		courseCategoryData := &data.CourseCategoryClientData{
			CategoryId: courseCategoryDB.Id,

			ModuleId: courseCategoryDB.ModuleId,

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

func ListCourseCatalogsClient(ctx context.Context, courseId string) (*data.ListCourseCatalogsClientRespData, *terror.Terror) {
	courseDB, errx := dbmodel.FindCourse(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course catalogs client (course id: %s) err (db find course %v)",
			courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseDB == nil {
		return nil, nil
	}

	if courseDB.Status != dbmodel.CourseStatusNormal {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course catalogs client (course id: %s) err (course status invalid)",
			courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseInvalid, errMsg)

		return nil, errx
	}

	if courseDB.CourseType != dbmodel.CourseTypeVideo {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course catalogs client (course id: %s) err (course type invalid)",
			courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseTypeInvalid, errMsg)

		return nil, errx
	}

	courseCatalogsDB, errx := dbmodel.FindCourseCatalogs(ctx, courseId, dbmodel.CourseCatalogStatusNormal)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course catalogs client (course id: %s) err (db find course catalogs %v)",
			courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	catalogIds := make([]string, 0, len(courseCatalogsDB))

	for _, courseCatalogDB := range courseCatalogsDB {
		catalogIds = append(catalogIds, courseCatalogDB.Id)
	}

	courseVideosDBMap := make(map[string]*dbmodel.CourseVideo)

	if len(catalogIds) > 0 {
		courseVideosDB, errx := dbmodel.FindCourseVideosByCatalog(ctx, catalogIds)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("List course catalogs client (course id: %s, catalog ids: %v) err (db find course videos %v)",
				courseId, catalogIds, errx)
			errx.AttachErrMsg(errMsg)

			return nil, errx
		}

		for _, courseVideoDB := range courseVideosDB {
			courseVideosDBMap[courseVideoDB.CatalogId] = courseVideoDB
		}
	}

	listCourseCatalogsClientRespData := &data.ListCourseCatalogsClientRespData{
		Catalogs: make([]*data.CourseCatalogClientData, 0),
	}

	for _, courseCatalogDB := range courseCatalogsDB {
		courseCatalogData := &data.CourseCatalogClientData{
			CatalogId: courseCatalogDB.Id,

			ParentId: courseCatalogDB.ParentId,

			Name: courseCatalogDB.Name,
		}

		courseVideoDB, ok := courseVideosDBMap[courseCatalogDB.Id]
		if ok {
			courseCatalogData.Video = &data.CourseCatalogVideoClientData{
				VideoId: courseVideoDB.Id,

				VideoUrl: courseVideoDB.VideoUrl,

				Duration: courseVideoDB.Duration,
				Format:   courseVideoDB.Format,
				Language: courseVideoDB.Language,
				Size:     courseVideoDB.Size,

				UploadAt: courseVideoDB.UploadAt.Format(time.RFC3339),
			}
		}

		listCourseCatalogsClientRespData.Catalogs = append(listCourseCatalogsClientRespData.Catalogs, courseCatalogData)
	}

	return listCourseCatalogsClientRespData, nil
}
