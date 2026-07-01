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

		if len(tags) > 0 {
			_, errx = dbmodel.CreateCourseTags(ctx, tx, courseId, tags)
			if errx != nil {
				errMsg := tlog.E(ctx).Err(errx).Msgf("Create course admin (user id: %s, category id: %s, course type: %d, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db create course tags %v)",
					userId, categoryId, courseType, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)
				errx.AttachErrMsg(errMsg)

				return errx
			}
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
	courseDB, errx := dbmodel.FindCourse(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update course admin (user id: %s, course id: %s, category id: %s, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db find course %v)",
			userId, courseId, categoryId, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if courseDB == nil {
		errMsg := tlog.E(ctx).Msgf("Update course admin (user id: %s, course id: %s, category id: %s, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (course not exist)",
			userId, courseId, categoryId, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseNotExist, errMsg)

		return errx
	}

	if categoryId != courseDB.CategoryId {
		courseCategoryDB, errx := dbmodel.FindCourseCategory(ctx, categoryId)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Update course admin (user id: %s, course id: %s, category id: %s, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db find course category %v)",
				userId, courseId, categoryId, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		if courseCategoryDB == nil {
			errMsg := tlog.E(ctx).Msgf("Update course admin (user id: %s, course id: %s, category id: %s, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (course category not exist)",
				userId, courseId, categoryId, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status)

			errx := terror.NewTerror(ctx, terror.ErrParamInvalid("category id"), constant.ErrorCodeCourseCategoryNotExist, errMsg)

			return errx
		}
	}

	err := dbmodel.DB(ctx).Transaction(func(tx *gorm.DB) error {
		errx := dbmodel.UpdateCourse(ctx, tx, courseId, author, source, title, abstract, coverUrl, linkUrl, publishAt, status)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Update course admin (user id: %s, course id: %s, category id: %s, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db update course %v)",
				userId, courseId, categoryId, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		errx = dbmodel.UpdateCourseDetail(ctx, tx, courseId, detail, summary, objective, outline, references)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Update course admin (user id: %s, course id: %s, category id: %s, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db update course detail %v)",
				userId, courseId, categoryId, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		errx = dbmodel.DeleteCourseTags(ctx, tx, courseId)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Update course admin (user id: %s, course id: %s, category id: %s, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db delete course tags %v)",
				userId, courseId, categoryId, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		if len(tags) > 0 {
			_, errx = dbmodel.CreateCourseTags(ctx, tx, courseId, tags)
			if errx != nil {
				errMsg := tlog.E(ctx).Err(errx).Msgf("Update course admin (user id: %s, course id: %s, category id: %s, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db create course tags %v)",
					userId, courseId, categoryId, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)
				errx.AttachErrMsg(errMsg)

				return errx
			}
		}

		return nil
	})
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Update course admin (user id: %s, course id: %s, category id: %s, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (db transaction %v)",
			userId, courseId, categoryId, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, err)

		errx := terror.NewTerror(ctx, terror.ErrSvcAbnormal("db transaction"), constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteCourseAdmin(ctx context.Context, userId string, courseId string) *terror.Terror {
	courseDB, errx := dbmodel.FindCourse(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course admin (user id: %s, course id: %s) err (db find course %v)",
			userId, courseId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if courseDB == nil {
		errMsg := tlog.E(ctx).Msgf("Delete course admin (user id: %s, course id: %s) err (course not exist)",
			userId, courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseNotExist, errMsg)

		return errx
	}

	// TODO: 如果收藏了，是否需要处理？

	err := dbmodel.DB(ctx).Transaction(func(tx *gorm.DB) error {
		errx := dbmodel.DeleteCourse(ctx, tx, courseId)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course admin (user id: %s, course id: %s) err (db delete course %v)",
				userId, courseId, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		errx = dbmodel.DeleteCourseDetail(ctx, tx, courseId)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course admin (user id: %s, course id: %s) err (db delete course detail %v)",
				userId, courseId, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		errx = dbmodel.DeleteCourseTags(ctx, tx, courseId)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course admin (user id: %s, course id: %s) err (db delete course tags %v)",
				userId, courseId, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		courseCatalogsDB, errx := dbmodel.FindCourseCatalogs(ctx, courseId, -1)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course admin (user id: %s, course id: %s) err (db find course catalogs %v)",
				userId, courseId, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		catalogIds := make([]string, 0, len(courseCatalogsDB))

		for _, courseCatalogDB := range courseCatalogsDB {
			catalogIds = append(catalogIds, courseCatalogDB.Id)
		}

		errx = dbmodel.DeleteCourseVideosByCatalog(ctx, tx, catalogIds)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course admin (user id: %s, course id: %s) err (db delete course videos %v)",
				userId, courseId, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		errx = dbmodel.DeleteCourseCatalogs(ctx, tx, courseId)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course admin (user id: %s, course id: %s) err (db delete course catalogs %v)",
				userId, courseId, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		return nil
	})
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Delete course admin (user id: %s, course id: %s) err (db transaction %v)",
			userId, courseId, err)

		errx := terror.NewTerror(ctx, terror.ErrSvcAbnormal("db transaction"), constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func ListCourseCatalogsAdmin(ctx context.Context, userId string, courseId string) (*data.ListCourseCatalogsAdminRespData, *terror.Terror) {
	courseDB, errx := dbmodel.FindCourse(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course catalogs admin (user id: %s, course id: %s) err (db find course %v)",
			userId, courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseDB == nil {
		errMsg := tlog.E(ctx).Msgf("List course catalogs admin (user id: %s, course id: %s) err (course not exist)",
			userId, courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseNotExist, errMsg)

		return nil, errx
	}

	if courseDB.CourseType != dbmodel.CourseTypeVideo {
		errMsg := tlog.E(ctx).Msgf("List course catalogs admin (user id: %s, course id: %s) err (course type invalid)",
			userId, courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course type"), constant.ErrorCodeCourseTypeInvalid, errMsg)

		return nil, errx
	}

	courseCatalogsDB, errx := dbmodel.FindCourseCatalogs(ctx, courseId, -1)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List course catalogs admin (user id: %s, course id: %s) err (db find course catalogs %v)",
			userId, courseId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	catalogIds := make([]string, 0, len(courseCatalogsDB))

	for _, courseCatalogDB := range courseCatalogsDB {
		catalogIds = append(catalogIds, courseCatalogDB.Id)
	}

	courseVideosDBMap := make(map[string]*dbmodel.CourseVideo)

	if len(catalogIds) > 0 {
		courseVideosDB, errx := dbmodel.FindCourseVideosByCatalog(ctx, catalogIds, -1)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("List course catalogs admin (user id: %s, course id: %s, catalog ids: %v) err (db find course videos %v)",
				userId, courseId, catalogIds, errx)
			errx.AttachErrMsg(errMsg)

			return nil, errx
		}

		for _, courseVideoDB := range courseVideosDB {
			courseVideosDBMap[courseVideoDB.CatalogId] = courseVideoDB
		}
	}

	listCourseCatalogsRespData := &data.ListCourseCatalogsAdminRespData{
		Catalogs: make([]*data.CourseCatalogAdminData, 0),
	}

	for _, courseCatalogDB := range courseCatalogsDB {
		courseCatalogData := &data.CourseCatalogAdminData{
			CatalogId: courseCatalogDB.Id,

			ParentId: courseCatalogDB.ParentId,

			Name: courseCatalogDB.Name,

			Weight: courseCatalogDB.Weight,
			Status: courseCatalogDB.Status,

			CreatedAt: courseCatalogDB.CreatedAt.Format(time.RFC3339),
			UpdatedAt: courseCatalogDB.UpdatedAt.Format(time.RFC3339),
		}

		courseVideoDB, ok := courseVideosDBMap[courseCatalogDB.Id]
		if ok {
			courseCatalogData.Video = &data.CourseCatalogVideoAdminData{
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
		}

		listCourseCatalogsRespData.Catalogs = append(listCourseCatalogsRespData.Catalogs, courseCatalogData)
	}

	return listCourseCatalogsRespData, nil
}

func CreateCourseCatalogAdmin(ctx context.Context, userId string, courseId string, parentId string, name string, weight, status int, videoUrl string, format, language, size, duration string, uploadAt time.Time, videoWeight, videoStatus int) (*data.CreateCourseCatalogAdminRespData, *terror.Terror) {
	courseDB, errx := dbmodel.FindCourse(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create course catalog admin (user id: %s, course id: %s, parent id: %s, name: %s, weight: %d, status: %d, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %s, video weight: %d, video status: %d) err (db find course %v)",
			userId, courseId, parentId, name, weight, status, videoUrl, format, language, size, duration, uploadAt, videoWeight, videoStatus, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if courseDB == nil {
		errMsg := tlog.E(ctx).Msgf("Create course catalog admin (user id: %s, course id: %s, parent id: %s, name: %s, weight: %d, status: %d, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %s, video weight: %d, video status: %d) err (course not exist)",
			userId, courseId, parentId, name, weight, status, videoUrl, format, language, size, duration, uploadAt, videoWeight, videoStatus)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course id"), constant.ErrorCodeCourseNotExist, errMsg)

		return nil, errx
	}

	if courseDB.CourseType != dbmodel.CourseTypeVideo {
		errMsg := tlog.E(ctx).Msgf("Create course catalog admin (user id: %s, course id: %s) err (course type invalid)",
			userId, courseId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("course type"), constant.ErrorCodeCourseTypeInvalid, errMsg)

		return nil, errx
	}

	if parentId != "" {
		parentCatalogDB, errx := dbmodel.FindCourseCatalog(ctx, parentId)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Create course catalog admin (user id: %s, course id: %s, parent id: %s) err (db find parent catalog %v)",
				userId, courseId, parentId, errx)
			errx.AttachErrMsg(errMsg)

			return nil, errx
		}

		if parentCatalogDB == nil || parentCatalogDB.CourseId != courseId {
			errMsg := tlog.E(ctx).Msgf("Create course catalog admin (user id: %s, course id: %s, parent id: %s) err (parent catalog not exist)",
				userId, courseId, parentId)

			errx := terror.NewTerror(ctx, terror.ErrParamInvalid("parent id"), constant.ErrorCodeCourseCatalogNotExist, errMsg)

			return nil, errx
		}
	}

	catalogId := ""

	err := dbmodel.DB(ctx).Transaction(func(tx *gorm.DB) error {
		courseCatalogDB, errx := dbmodel.CreateCourseCatalog(ctx, tx, courseId, parentId, name, weight, status)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Create course catalog admin (user id: %s, course id: %s, parent id: %s, name: %s, weight: %d, status: %d) err (db create course catalog %v)",
				userId, courseId, parentId, name, weight, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		catalogId = courseCatalogDB.Id

		_, errx = dbmodel.CreateCourseVideo(ctx, tx, catalogId, videoUrl, format, language, size, duration, uploadAt, videoWeight, videoStatus)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Create course catalog admin (user id: %s, catalog id: %s, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %s, video weight: %d, video status: %d) err (db create course video %v)",
				userId, catalogId, videoUrl, format, language, size, duration, uploadAt, videoWeight, videoStatus, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		return nil
	})
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Create course catalog admin (user id: %s, course id: %s, parent id: %s, name: %s, weight: %d, status: %d) err (db transaction %v)",
			userId, courseId, parentId, name, weight, status, err)

		errx := terror.NewTerror(ctx, terror.ErrSvcAbnormal("db transaction"), constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	createCourseCatalogRespData := &data.CreateCourseCatalogAdminRespData{
		CatalogId: catalogId,
	}

	return createCourseCatalogRespData, nil
}

func UpdateCourseCatalogAdmin(ctx context.Context, userId string, catalogId string, parentId string, name string, weight, status int, videoUrl string, format, language, size, duration string, uploadAt time.Time, videoWeight, videoStatus int) *terror.Terror {
	courseCatalogDB, errx := dbmodel.FindCourseCatalog(ctx, catalogId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update course catalog admin (user id: %s, catalog id: %s) err (db find course catalog %v)",
			userId, catalogId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if courseCatalogDB == nil {
		errMsg := tlog.E(ctx).Msgf("Update course catalog admin (user id: %s, catalog id: %s) err (course catalog not exist)",
			userId, catalogId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("catalog id"), constant.ErrorCodeCourseCatalogNotExist, errMsg)

		return errx
	}

	if parentId == catalogId {
		errMsg := tlog.E(ctx).Msgf("Update course catalog admin (user id: %s, catalog id: %s, parent id: %s) err (parent id invalid)",
			userId, catalogId, parentId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("parent id"), constant.ErrorCodeRequestParamInvalid, errMsg)

		return errx
	}

	if parentId != "" && parentId != courseCatalogDB.ParentId {
		parentCatalogDB, errx := dbmodel.FindCourseCatalog(ctx, parentId)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Update course catalog admin (user id: %s, catalog id: %s, parent id: %s) err (db find parent catalog %v)",
				userId, catalogId, parentId, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		if parentCatalogDB == nil || parentCatalogDB.CourseId != courseCatalogDB.CourseId {
			errMsg := tlog.E(ctx).Msgf("Update course catalog admin (user id: %s, catalog id: %s, parent id: %s) err (parent catalog not exist)",
				userId, catalogId, parentId)

			errx := terror.NewTerror(ctx, terror.ErrParamInvalid("parent id"), constant.ErrorCodeCourseCatalogNotExist, errMsg)

			return errx
		}
	}

	courseVideoDB, errx := dbmodel.FindCourseVideoByCatalog(ctx, catalogId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update course catalog admin (user id: %s, catalog id: %s) err (db find course video by catalog %v)",
			userId, catalogId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	err := dbmodel.DB(ctx).Transaction(func(tx *gorm.DB) error {
		errx := dbmodel.UpdateCourseCatalog(ctx, tx, catalogId, parentId, name, weight, status)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Update course catalog admin (user id: %s, catalog id: %s, parent id: %s, name: %s, weight: %d, status: %d) err (db update course catalog %v)",
				userId, catalogId, parentId, name, weight, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		if courseVideoDB == nil {
			_, errx = dbmodel.CreateCourseVideo(ctx, tx, catalogId, videoUrl, format, language, size, duration, uploadAt, videoWeight, videoStatus)
			if errx != nil {
				errMsg := tlog.E(ctx).Err(errx).Msgf("Update course catalog admin (user id: %s, catalog id: %s) err (db create course video %v)",
					userId, catalogId, errx)
				errx.AttachErrMsg(errMsg)

				return errx
			}

			return nil
		}

		videoId := courseVideoDB.Id

		errx = dbmodel.UpdateCourseVideo(ctx, tx, videoId, videoUrl, format, language, size, duration, uploadAt, videoWeight, videoStatus)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Update course catalog admin (user id: %s, catalog id: %s, video id: %s) err (db update course video %v)",
				userId, catalogId, courseVideoDB.Id, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		return nil
	})
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Update course catalog admin (user id: %s, catalog id: %s) err (db transaction %v)",
			userId, catalogId, err)

		errx := terror.NewTerror(ctx, terror.ErrSvcAbnormal("db transaction"), constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteCourseCatalogAdmin(ctx context.Context, userId string, catalogId string) *terror.Terror {
	courseCatalogDB, errx := dbmodel.FindCourseCatalog(ctx, catalogId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course catalog admin (user id: %s, catalog id: %s) err (db find course catalog %v)",
			userId, catalogId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if courseCatalogDB == nil {
		errMsg := tlog.E(ctx).Msgf("Delete course catalog admin (user id: %s, catalog id: %s) err (course catalog not exist)",
			userId, catalogId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("catalog id"), constant.ErrorCodeCourseCatalogNotExist, errMsg)

		return errx
	}

	courseVideoDB, errx := dbmodel.FindCourseVideoByCatalog(ctx, catalogId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course catalog admin (user id: %s, catalog id: %s) err (db find course video by catalog %v)",
			userId, catalogId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	err := dbmodel.DB(ctx).Transaction(func(tx *gorm.DB) error {
		if courseVideoDB != nil {
			errx := dbmodel.DeleteCourseVideo(ctx, tx, courseVideoDB.Id)
			if errx != nil {
				errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course catalog admin (user id: %s, catalog id: %s, video id: %s) err (db delete course video %v)",
					userId, catalogId, courseVideoDB.Id, errx)
				errx.AttachErrMsg(errMsg)

				return errx
			}
		}

		errx := dbmodel.DeleteCourseCatalog(ctx, tx, catalogId)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Delete course catalog admin (user id: %s, catalog id: %s) err (db delete course catalog %v)",
				userId, catalogId, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		return nil
	})
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Delete course catalog admin (user id: %s, catalog id: %s) err (db transaction %v)",
			userId, catalogId, err)

		errx := terror.NewTerror(ctx, terror.ErrSvcAbnormal("db transaction"), constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}
