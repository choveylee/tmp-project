package service

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"gorm.io/gorm"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/data"
	dbmodel "dev.choveylee.top/elder-care-backend/internal/model/mysql"
)

func ListCourseCategoriesAdmin(ctx context.Context, userId string, status int, pageNum, pageSize int) (*data.ListCourseCategoriesAdminRespData, *terror.Terror) {
	total, courseCategoriesDB, errx := dbmodel.FindCourseCategories(ctx, status, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course categories admin (user id: %s, status: %d, page num: %d, page size: %d) err (db find course categories %v)",
			userId, status, pageNum, pageSize, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listCourseCategoriesRespData := &data.ListCourseCategoriesAdminRespData{
		Categories: make([]*data.CourseCategoryAdminData, 0),

		Total: total,
	}

	for _, courseCategoryDB := range courseCategoriesDB {
		courseCategoryData := &data.CourseCategoryAdminData{
			CategoryId: courseCategoryDB.Id,

			Name: courseCategoryDB.Name,

			Weight: courseCategoryDB.Weight,

			Status: courseCategoryDB.Status,

			CreatedAt: courseCategoryDB.CreatedAt.Format(time.RFC3339),
			UpdatedAt: courseCategoryDB.UpdatedAt.Format(time.RFC3339),
		}

		listCourseCategoriesRespData.Categories = append(listCourseCategoriesRespData.Categories, courseCategoryData)
	}

	return listCourseCategoriesRespData, nil
}

func CreateCourseCategoryAdmin(ctx context.Context, userId string, name string, weight, status int) (*data.CreateCourseCategoryAdminRespData, *terror.Terror) {
	courseCategoryDB, errx := dbmodel.FindCourseCategoryByName(ctx, name)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create course category admin (user id: %s, name: %s, weight: %d, status: %d) err (db find course category by name %v)",
			userId, name, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseCategoryDB != nil {
		errMsg := tlog.E(ctx).Msgf("Create course category admin (user id: %s, name: %s, weight: %d, status: %d) err (name exist)",
			userId, name, weight, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("name"), constant.ErrorCodeCourseCategoryNameExist, errMsg)

		return nil, errx
	}

	courseCategoryDB, errx = dbmodel.CreateCourseCategory(ctx, name, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create course category admin (user id: %s, name: %s, weight: %d, status: %d) err (db create course category %v)",
			userId, name, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	createCourseCategoryRespData := &data.CreateCourseCategoryAdminRespData{
		CategoryId: courseCategoryDB.Id,
	}

	return createCourseCategoryRespData, nil
}

func GetCourseCategoryAdmin(ctx context.Context, userId string, categoryId string) (*data.GetCourseCategoryAdminRespData, *terror.Terror) {
	courseCategoryDB, errx := dbmodel.FindCourseCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get course category admin (user id: %s, category id: %s) err (db find course category %v)",
			userId, categoryId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseCategoryDB == nil {
		errMsg := tlog.E(ctx).Msgf("Get course category admin (user id: %s, category id: %s) err (course category not exist)",
			userId, categoryId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("category id"), constant.ErrorCodeCourseCategoryNotExist, errMsg)

		return nil, errx
	}

	getCourseCategoryRespData := &data.GetCourseCategoryAdminRespData{
		CategoryId: courseCategoryDB.Id,

		Name: courseCategoryDB.Name,

		Weight: courseCategoryDB.Weight,

		Status: courseCategoryDB.Status,

		CreatedAt: courseCategoryDB.CreatedAt.Format(time.RFC3339),
		UpdatedAt: courseCategoryDB.UpdatedAt.Format(time.RFC3339),
	}

	return getCourseCategoryRespData, nil
}

func UpdateCourseCategoryAdmin(ctx context.Context, userId string, categoryId string, name string, weight, status int) *terror.Terror {
	courseCategoryDB, errx := dbmodel.FindCourseCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update course category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (db find course category %v)",
			userId, categoryId, name, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if courseCategoryDB == nil {
		errMsg := tlog.E(ctx).Msgf("Update course category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (course category not exist)",
			userId, categoryId, name, weight, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("category id"), constant.ErrorCodeCourseCategoryNotExist, errMsg)

		return errx
	}

	if courseCategoryDB.Name != name {
		courseCategoryDB, errx := dbmodel.FindCourseCategoryByName(ctx, name)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Update course category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (db find course category by name %v)",
				userId, categoryId, name, weight, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		if courseCategoryDB != nil {
			errMsg := tlog.E(ctx).Msgf("Update course category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (name exist)",
				userId, categoryId, name, weight, status)

			errx := terror.NewTerror(ctx, terror.ErrParamInvalid("name"), constant.ErrorCodeCourseCategoryNameExist, errMsg)

			return errx
		}
	}

	errx = dbmodel.UpdateCourseCategory(ctx, categoryId, name, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update course category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (db update course category %v)",
			userId, categoryId, name, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}

func DeleteCourseCategoryAdmin(ctx context.Context, userId string, categoryId string) *terror.Terror {
	courseCategoryDB, errx := dbmodel.FindCourseCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course category admin (user id: %s, category id: %s) err (db find course category %v)",
			userId, categoryId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if courseCategoryDB == nil {
		errMsg := tlog.E(ctx).Msgf("Delete course category admin (user id: %s, category id: %s) err (course category not exist)",
			userId, categoryId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("category id"), constant.ErrorCodeCourseCategoryNotExist, errMsg)

		return errx
	}

	total, errx := dbmodel.FindCourseCountByCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course category admin (user id: %s, category id: %s) err (db find course count by category %v)",
			userId, categoryId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if total > 0 {
		errMsg := tlog.E(ctx).Msgf("Delete course category admin (user id: %s, category id: %s) err (course category in use)",
			userId, categoryId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("category id"), constant.ErrorCodeCourseCategoryInUse, errMsg)

		return errx
	}

	errx = dbmodel.DeleteCourseCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course category admin (user id: %s, category id: %s) err (db delete course category %v)",
			userId, categoryId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}

func ListCoursesAdmin(ctx context.Context, userId string, categoryId string, courseType int, status int, pageNum, pageSize int) (*data.ListCoursesAdminRespData, *terror.Terror) {
	total, coursesDB, errx := dbmodel.FindCourses(ctx, categoryId, courseType, status, dbmodel.CourseSortByPublish, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List courses admin (user id: %s, category id: %s, course type: %d, status: %d, page num: %d, page size: %d) err (db find courses %v)",
			userId, categoryId, courseType, status, pageNum, pageSize, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listCoursesRespData := &data.ListCoursesAdminRespData{
		Courses: make([]*data.CourseAdminData, 0),

		Total: total,
	}

	categoryIdsMap := make(map[string]bool)

	for _, courseDB := range coursesDB {
		courseData := &data.CourseAdminData{
			CourseId: courseDB.Id,

			CategoryId: courseDB.CategoryId,

			CourseType: courseDB.CourseType,

			Author: courseDB.Author,
			Source: courseDB.Source,

			Title:    courseDB.Title,
			Abstract: courseDB.Abstract,

			CoverUrl: courseDB.CoverUrl,
			LinkUrl:  courseDB.LinkUrl,

			PublishAt: courseDB.PublishAt.Format(time.RFC3339),

			Status: courseDB.Status,

			CreatedAt: courseDB.CreatedAt.Format(time.RFC3339),
			UpdatedAt: courseDB.UpdatedAt.Format(time.RFC3339),
		}

		categoryIdsMap[courseDB.CategoryId] = true

		listCoursesRespData.Courses = append(listCoursesRespData.Courses, courseData)
	}

	if len(categoryIdsMap) > 0 {
		categoryIds := make([]string, 0)

		for categoryId := range categoryIdsMap {
			categoryIds = append(categoryIds, categoryId)
		}

		courseCategoriesDB, errx := dbmodel.FindCourseCategoriesById(ctx, categoryIds)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("List courses admin (user id: %s, category id: %s, course type: %d, status: %d, page num: %d, page size: %d, category ids: %v) err (db find course categories %v)",
				userId, categoryId, courseType, status, pageNum, pageSize, categoryIds, errx)
			errx.AttachErrMsg(errMsg)

			return nil, errx
		}

		courseCategoriesMap := make(map[string]*dbmodel.CourseCategory)

		for _, courseCategoryDB := range courseCategoriesDB {
			courseCategoriesMap[courseCategoryDB.Id] = courseCategoryDB
		}

		for _, courseData := range listCoursesRespData.Courses {
			courseCategoryDB, ok := courseCategoriesMap[courseData.CategoryId]
			if ok {
				courseData.CategoryName = courseCategoryDB.Name
			}
		}
	}

	return listCoursesRespData, nil
}

func CreateCourseAdmin(ctx context.Context, userId string, categoryId string, courseType int, author, source string, title string, tags []string, abstract string, coverUrl, linkUrl string, detail, summary, objective, outline, references string, publishAt *time.Time, status int) (*data.CreateCourseAdminRespData, *terror.Terror) {
	courseCategoryDB, errx := dbmodel.FindCourseCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create course admin (user id: %s, category id: %s, course type: %d, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db find course category %v)",
			userId, categoryId, courseType, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseCategoryDB == nil {
		errMsg := tlog.E(ctx).Msgf("Create course admin (user id: %s, category id: %s, course type: %d, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (course category not exist)",
			userId, categoryId, courseType, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("category id"), constant.ErrorCodeCourseCategoryNotExist, errMsg)

		return nil, errx
	}

	courseId := ""

	err := dbmodel.DB(ctx).Transaction(func(tx *gorm.DB) error {
		courseDB, errx := dbmodel.CreateCourse(ctx, tx, categoryId, courseType, author, source, title, abstract, coverUrl, linkUrl, publishAt, status)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Create course admin (user id: %s, category id: %s, course type: %d, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db create course %v)",
				userId, categoryId, courseType, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		courseId = courseDB.Id

		_, errx = dbmodel.CreateCourseDetail(ctx, tx, courseId, detail, summary, objective, outline, references)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Create course admin (user id: %s, category id: %s, course type: %d, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db create course detail %v)",
				userId, categoryId, courseType, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		_, errx = dbmodel.CreateCourseTags(ctx, tx, courseId, tags)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Create course admin (user id: %s, category id: %s, course type: %d, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db create course tags %v)",
				userId, categoryId, courseType, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		return nil
	})
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Create course admin (user id: %s, category id: %s, course type: %d, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db transaction %v)",
			userId, categoryId, courseType, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	createCourseRespData := &data.CreateCourseAdminRespData{
		CourseId: courseId,
	}

	return createCourseRespData, nil
}

func GetCourseAdmin(ctx context.Context, userId string, courseId string) (*data.GetCourseAdminRespData, *terror.Terror) {
	courseDB, errx := dbmodel.FindCourse(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get course admin (user id: %s, course id: %s) err (db find course %v)",
			userId, courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseDB == nil {
		errMsg := tlog.E(ctx).Msgf("Get course admin (user id: %s, course id: %s) err (course not exist)",
			userId, courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseNotExist, errMsg)

		return nil, errx
	}

	courseDetailDB, errx := dbmodel.FindCourseDetail(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get course admin (user id: %s, course id: %s) err (db find course detail %v)",
			userId, courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseDetailDB == nil {
		errMsg := tlog.E(ctx).Msgf("Get course admin (user id: %s, course id: %s) err (course detail not exist)",
			userId, courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseDetailNotExist, errMsg)

		return nil, errx
	}

	categoryId := courseDB.CategoryId

	courseCategoryDB, errx := dbmodel.FindCourseCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get course admin (user id: %s, course id: %s) err (db find course category %v)",
			userId, courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseCategoryDB == nil {
		errMsg := tlog.E(ctx).Msgf("Get course admin (user id: %s, course id: %s) err (course category not exist)",
			userId, courseId)

		errx := terror.NewTerror(ctx, terror.ErrSvcAbnormal("course id"), constant.ErrorCodeCourseCategoryNotExist, errMsg)

		return nil, errx
	}

	courseTagsDB, errx := dbmodel.FindCourseTags(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get course admin (user id: %s, course id: %s) err (db find course tags %v)",
			userId, courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	getCourseRespData := &data.GetCourseAdminRespData{
		CourseId: courseDB.Id,

		CategoryId:   courseDB.CategoryId,
		CategoryName: courseCategoryDB.Name,

		CourseType: courseDB.CourseType,

		Author: courseDB.Author,
		Source: courseDB.Source,

		Title: courseDB.Title,

		// Tags []string `json:"tags"`

		Abstract: courseDB.Abstract,

		CoverUrl: courseDB.CoverUrl,
		LinkUrl:  courseDB.LinkUrl,

		Detail: courseDetailDB.Detail,

		Summary:    courseDetailDB.Summary,
		Objective:  courseDetailDB.Objective,
		Outline:    courseDetailDB.Outline,
		References: courseDetailDB.References,

		// Videos []*CourseVideoAdminData `json:"videos"`

		PublishAt: courseDB.PublishAt.Format(time.RFC3339),

		CreatedAt: courseDB.CreatedAt.Format(time.RFC3339),
		UpdatedAt: courseDB.UpdatedAt.Format(time.RFC3339),
	}

	for _, courseTagDB := range courseTagsDB {
		getCourseRespData.Tags = append(getCourseRespData.Tags, courseTagDB.Name)
	}

	return getCourseRespData, nil
}

func UpdateCourseAdmin(ctx context.Context, userId string, courseId string, categoryId string, author, source string, title string, tags []string, abstract string, coverUrl, linkUrl string, detail, summary, objective, outline, references string, publishAt *time.Time, status int) *terror.Terror {

}

func DeleteCourseAdmin(ctx context.Context, userId string, courseId string) *terror.Terror {

}

func ListCourseVideosAdmin(ctx context.Context, userId string, courseId string) (*data.ListCourseVideosAdminRespData, *terror.Terror) {
	courseDB, errx := dbmodel.FindCourse(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course videos admin (user id: %s, course id: %s) err (db find course %v)",
			userId, courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseDB == nil {
		errMsg := tlog.E(ctx).Msgf("List course videos admin (user id: %s, course id: %s) err (course not exist)",
			userId, courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseNotExist, errMsg)

		return nil, errx
	}

	if courseDB.CourseType != dbmodel.CourseTypeVideo {
		errMsg := tlog.E(ctx).Msgf("List course videos admin (user id: %s, course id: %s) err (course type invalid)",
			userId, courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course type"), constant.ErrorCodeCourseTypeInvalid, errMsg)

		return nil, errx
	}

	courseVideosDB, errx := dbmodel.FindCourseVideos(ctx, courseId, -1)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course videos admin (user id: %s, course id: %s) err (db find course videos %v)",
			userId, courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listCourseVideosRespData := &data.ListCourseVideosAdminRespData{}

	for _, courseVideoDB := range courseVideosDB {
		courseVideoData := &data.CourseVideoAdminData{
			VideoId: courseVideoDB.Id,

			VideoUrl: courseVideoDB.VideoUrl,

			Format:   courseVideoDB.Format,
			Language: courseVideoDB.Language,
			Size:     courseVideoDB.Size,
			Duration: courseVideoDB.Duration,

			UploadAt: courseVideoDB.UploadAt.Format(time.RFC3339),

			CreatedAt: courseVideoDB.CreatedAt.Format(time.RFC3339),
			UpdatedAt: courseVideoDB.UpdatedAt.Format(time.RFC3339),
		}

		listCourseVideosRespData.Videos = append(listCourseVideosRespData.Videos, courseVideoData)
	}

	return listCourseVideosRespData, nil
}

func CreateCourseVideoAdmin(ctx context.Context, userId string, courseId string, videoUrl string, format, language, size, duration string, uploadAt time.Time, weight, status int) (*data.CreateCourseVideoAdminRespData, *terror.Terror) {
	courseDB, errx := dbmodel.FindCourse(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create course video admin (user id: %s, course id: %s, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %s, weight: %d, status: %d) err (db find course %v)",
			userId, courseId, videoUrl, format, language, size, duration, uploadAt, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseDB == nil {
		errMsg := tlog.E(ctx).Msgf("Create course video admin (user id: %s, course id: %s, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %s, weight: %d, status: %d) err (course not exist)",
			userId, courseId, videoUrl, format, language, size, duration, uploadAt, weight, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseNotExist, errMsg)

		return nil, errx
	}

	courseVideoDB, errx := dbmodel.CreateCourseVideo(ctx, courseId, videoUrl, format, language, size, duration, uploadAt, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create course video admin (user id: %s, course id: %s, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %s, weight: %d, status: %d) err (db create course video %v)",
			userId, courseId, videoUrl, format, language, size, duration, uploadAt, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	createCourseVideoRespData := &data.CreateCourseVideoAdminRespData{
		VideoId: courseVideoDB.Id,
	}

	return createCourseVideoRespData, nil
}

func UpdateCourseVideoAdmin(ctx context.Context, userId string, videoId string, videoUrl string, format, language, size, duration string, uploadAt time.Time, weight, status int) *terror.Terror {
	courseVideoDB, errx := dbmodel.FindCourseVideo(ctx, videoId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update course video admin (user id: %s, video id: %s, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %s, weight: %d, status: %d) err (db find course video %v)",
			userId, videoId, videoUrl, format, language, size, duration, uploadAt, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if courseVideoDB == nil {
		errMsg := tlog.E(ctx).Msgf("Update course video admin (user id: %s, video id: %s, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %s, weight: %d, status: %d) err (course video not exist)",
			userId, videoId, videoUrl, format, language, size, duration, uploadAt, weight, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("video id"), constant.ErrorCodeCourseVideoNotExist, errMsg)

		return errx
	}

	errx = dbmodel.UpdateCourseVideo(ctx, videoId, videoUrl, format, language, size, duration, uploadAt, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update course video admin (user id: %s, video id: %s, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %s, weight: %d, status: %d) err (db update course video %v)",
			userId, videoId, videoUrl, format, language, size, duration, uploadAt, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}

func DeleteCourseVideoAdmin(ctx context.Context, userId string, videoId string) *terror.Terror {
	courseVideoDB, errx := dbmodel.FindCourseVideo(ctx, videoId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course video admin (user id: %s, video id: %s) err (db find course video %v)",
			userId, videoId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if courseVideoDB == nil {
		errMsg := tlog.E(ctx).Msgf("Delete course video admin (user id: %s, video id: %s) err (course video not exist)",
			userId, videoId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("video id"), constant.ErrorCodeCourseVideoNotExist, errMsg)

		return errx
	}

	errx = dbmodel.DeleteCourseVideo(ctx, videoId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course video admin (user id: %s, video id: %s) err (db delete course video %v)",
			userId, videoId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}
