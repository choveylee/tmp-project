package handler

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/choveylee/tlog"
	"github.com/gin-gonic/gin"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/data"
	dbmodel "dev.choveylee.top/elder-care-backend/internal/model/mysql"
	"dev.choveylee.top/elder-care-backend/internal/service"
)

func HandleListCourseCategoriesAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	status := -1

	srcStatus := strings.TrimSpace(c.Query("status"))
	if srcStatus != "" {
		desStatus, err := strconv.Atoi(srcStatus)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list course categories admin (status: %s) err (stconv atoi %v)",
				srcStatus, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		_, ok := dbmodel.CourseCategoryStatusesMap[desStatus]
		if !ok {
			errMsg := tlog.E(ctx).Msgf("Handle list course categories admin (status: %d) err (status invalid)",
				desStatus)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		status = desStatus
	}

	pageNum := 1

	srcPageNum := strings.TrimSpace(c.Query("page_num"))
	if srcPageNum != "" {
		desPageNum, err := strconv.Atoi(srcPageNum)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list course categories admin (page num: %s) err (strconv atoi %v)",
				srcPageNum, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageNum <= 0 || desPageNum > constant.MaxPageNum {
			errMsg := tlog.E(ctx).Msgf("Handle list course categories admin (page num: %d) err (page num invalid)",
				desPageNum)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageNum = desPageNum
	}

	pageSize := 10

	srcPageSize := strings.TrimSpace(c.Query("page_size"))
	if srcPageSize != "" {
		desPageSize, err := strconv.Atoi(srcPageSize)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list course categories admin (page size: %s) err (strconv atoi %v)",
				srcPageSize, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageSize <= 0 || desPageSize > constant.MaxPageSize {
			errMsg := tlog.E(ctx).Msgf("Handle list course categories admin (page size: %d) err (page size invalid)",
				desPageSize)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageSize = desPageSize
	}

	getCourseCategoriesRespData, errx := service.ListCourseCategoriesAdmin(ctx, userId, status, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list course categories admin (user id: %s, status: %d, page num: %d, page size: %d) err (list course categories admin %v)",
			userId, status, pageNum, pageSize, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getCourseCategoriesRespData)
}

func HandleCreateCourseCategoryAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	createCourseCategoryRequest := &data.CreateCourseCategoryAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, createCourseCategoryRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle create course category admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	name := strings.TrimSpace(createCourseCategoryRequest.Name)
	if name == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create course category admin err (name invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(name) > dbmodel.CourseCategoryNameLen {
		errMsg := tlog.E(ctx).Msgf("Handle create course category admin (name: %s) err (name len limit)",
			name)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	weight := createCourseCategoryRequest.Weight

	status := createCourseCategoryRequest.Status

	_, ok := dbmodel.CourseCategoryStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle create course category admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	createCourseCategoryRespData, errx := service.CreateCourseCategoryAdmin(ctx, userId, name, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle create course category admin (user id: %s, name: %s, weight: %d, status: %d) err (create course category admin %v)",
			userId, name, weight, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, createCourseCategoryRespData)
}

func HandleGetCourseCategoryAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	categoryId := strings.TrimSpace(c.Param("id"))
	if categoryId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle get course category admin err (category id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	getCourseCategoryRespData, errx := service.GetCourseCategoryAdmin(ctx, userId, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle get course category admin (user id: %s, category id: %s) err (get course category admin %v)",
			userId, categoryId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getCourseCategoryRespData)
}

func HandleUpdateCourseCategoryAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	categoryId := strings.TrimSpace(c.Param("id"))
	if categoryId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update course category admin err (category id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	updateCourseCategoryRequest := &data.UpdateCourseCategoryAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, updateCourseCategoryRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle update course category admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	name := strings.TrimSpace(updateCourseCategoryRequest.Name)
	if name == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update course category admin err (name invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(name) > dbmodel.CourseCategoryNameLen {
		errMsg := tlog.E(ctx).Msgf("Handle update course category admin (title: %s) err (name len limit)",
			name)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	weight := updateCourseCategoryRequest.Weight

	status := updateCourseCategoryRequest.Status

	_, ok := dbmodel.CourseCategoryStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle update course category admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.UpdateCourseCategoryAdmin(ctx, userId, categoryId, name, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle update course category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (update course category admin %v)",
			userId, categoryId, name, weight, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}

func HandleDeleteCourseCategoryAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	categoryId := strings.TrimSpace(c.Param("id"))
	if categoryId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle delete course category admin err (category id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.DeleteCourseCategoryAdmin(ctx, userId, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle delete course category admin (user id: %s, category id: %s) err (delete course category admin %v)",
			userId, categoryId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}

func HandleListCoursesAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	categoryId := strings.TrimSpace(c.Query("category_id"))

	courseType := -1

	srcCourseType := strings.TrimSpace(c.Query("course_type"))
	if srcCourseType != "" {
		desCourseType, err := strconv.Atoi(srcCourseType)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list courses admin (course type: %s) err (strconv atoi %v)",
				srcCourseType, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		_, ok := dbmodel.CourseTypesMap[desCourseType]
		if !ok {
			errMsg := tlog.E(ctx).Msgf("Handle list courses admin (course type: %d) err (course type invalid)",
				desCourseType)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		courseType = desCourseType
	}

	status := -1

	srcStatus := strings.TrimSpace(c.Query("status"))
	if srcStatus != "" {
		desStatus, err := strconv.Atoi(srcStatus)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list courses admin (status: %s) err (stconv atoi %v)",
				srcStatus, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		_, ok := dbmodel.CourseCategoryStatusesMap[desStatus]
		if !ok {
			errMsg := tlog.E(ctx).Msgf("Handle list courses admin (status: %d) err (status invalid)",
				desStatus)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		status = desStatus
	}

	pageNum := 1

	srcPageNum := strings.TrimSpace(c.Query("page_num"))
	if srcPageNum != "" {
		desPageNum, err := strconv.Atoi(srcPageNum)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list courses admin (page num: %s) err (strconv atoi %v)",
				srcPageNum, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageNum <= 0 || desPageNum > constant.MaxPageNum {
			errMsg := tlog.E(ctx).Msgf("Handle list courses admin (page num: %d) err (page num invalid)",
				desPageNum)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageNum = desPageNum
	}

	pageSize := 10

	srcPageSize := strings.TrimSpace(c.Query("page_size"))
	if srcPageSize != "" {
		desPageSize, err := strconv.Atoi(srcPageSize)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list courses admin (page size: %s) err (strconv atoi %v)",
				srcPageSize, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageSize <= 0 || desPageSize > constant.MaxPageSize {
			errMsg := tlog.E(ctx).Msgf("Handle list courses admin (page size: %d) err (page size invalid)",
				desPageSize)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageSize = desPageSize
	}

	getCoursesRespData, errx := service.ListCoursesAdmin(ctx, userId, categoryId, courseType, status, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list courses admin (user id: %s, category id: %s, course type: %d, status: %d, page num: %d, page size: %d) err (list course categories admin %v)",
			userId, categoryId, courseType, status, pageNum, pageSize, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getCoursesRespData)
}

func HandleCreateCourseAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	createCourseRequest := &data.CreateCourseAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, createCourseRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle create course admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	categoryId := strings.TrimSpace(createCourseRequest.CategoryId)

	courseType := createCourseRequest.CourseType

	_, ok := dbmodel.CourseTypesMap[courseType]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (course type: %d) err (course type invalid)",
			courseType)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	author := strings.TrimSpace(createCourseRequest.Author)

	if len(author) > dbmodel.CourseAuthorLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (author: %s) err (author length invalid)",
			author)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	source := strings.TrimSpace(createCourseRequest.Source)

	if len(source) > dbmodel.CourseSourceLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (source: %s) err (source length invalid)",
			source)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	title := strings.TrimSpace(createCourseRequest.Title)
	if title == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin err (title invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(title) > dbmodel.CourseTitleLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (title: %s) err (title length invalid)",
			title)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	tagsMap := make(map[string]bool)

	tags := createCourseRequest.Tags

	for index, tag := range tags {
		desTag := strings.TrimSpace(tag)
		tags[index] = desTag

		if desTag == "" {
			errMsg := tlog.E(ctx).Msgf("Handle create course admin err (tag invalid)")

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if len(desTag) > dbmodel.CourseTagNameLenLimit {
			errMsg := tlog.E(ctx).Msgf("Handle create course admin (tag: %s) err (tag length invalid)",
				tag)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		_, ok := tagsMap[desTag]
		if ok {
			errMsg := tlog.E(ctx).Msgf("Handle create course admin (tag: %s) err (tag duplicate)",
				desTag)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		tagsMap[desTag] = true
	}

	abstract := strings.TrimSpace(createCourseRequest.Abstract)

	if len(abstract) > dbmodel.CourseAbstractLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (abstract: %s) err (abstract length invalid)",
			abstract)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	coverUrl := strings.TrimSpace(createCourseRequest.CoverUrl)

	if len(coverUrl) > dbmodel.CourseCoverUrlLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (cover url: %s) err (cover url length invalid)",
			coverUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	linkUrl := strings.TrimSpace(createCourseRequest.LinkUrl)

	if len(linkUrl) > dbmodel.CourseLinkUrlLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (link url: %s) err (link url length invalid)",
			linkUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	detail := strings.TrimSpace(createCourseRequest.Detail)

	if len(detail) > dbmodel.CourseDetailLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (detail: %s) err (detail length invalid)",
			detail)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	summary := strings.TrimSpace(createCourseRequest.Summary)

	if len(summary) > dbmodel.CourseSummaryLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (summary: %s) err (summary length invalid)",
			summary)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	objective := strings.TrimSpace(createCourseRequest.Objective)

	if len(objective) > dbmodel.CourseObjectiveLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (objective: %s) err (objective length invalid)",
			objective)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	outline := strings.TrimSpace(createCourseRequest.Outline)

	if len(outline) > dbmodel.CourseOutlineLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (outline: %s) err (outline length invalid)",
			outline)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	references := strings.TrimSpace(createCourseRequest.References)

	if len(references) > dbmodel.CourseReferencesLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (references: %s) err (references length invalid)",
			references)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	var publishAt *time.Time

	srcPublishAt := createCourseRequest.PublishAt
	if srcPublishAt != "" {
		desPublishAt, err := time.ParseInLocation(time.RFC3339, srcPublishAt, time.Local)
		if err != nil {
			errMsg := tlog.E(ctx).Msgf("Handle create course admin (publish at: %s) err (publish at invalid)",
				srcPublishAt)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		publishAt = &desPublishAt
	}

	status := createCourseRequest.Status
	_, ok = dbmodel.CourseStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle create course admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	createCourseRespData, errx := service.CreateCourseAdmin(ctx, userId, categoryId, courseType, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle create course admin (user id: %s, category id: %s, course type: %d, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (create course admin %v)",
			userId, categoryId, courseType, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, createCourseRespData)
}

func HandleGetCourseAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	courseId := strings.TrimSpace(c.Param("id"))
	if courseId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle get course admin err (course id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	getCourseRespData, errx := service.GetCourseAdmin(ctx, userId, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle get course admin (user id: %s, course id: %s) err (get course admin %v)",
			userId, courseId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getCourseRespData)
}

func HandleUpdateCourseAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	courseId := strings.TrimSpace(c.Param("id"))
	if courseId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin err (course id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	updateCourseRequest := &data.UpdateCourseAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, updateCourseRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle update course admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	categoryId := strings.TrimSpace(updateCourseRequest.CategoryId)

	author := strings.TrimSpace(updateCourseRequest.Author)

	if len(author) > dbmodel.CourseAuthorLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (author: %s) err (author length invalid)",
			author)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	source := strings.TrimSpace(updateCourseRequest.Source)

	if len(source) > dbmodel.CourseSourceLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (source: %s) err (source length invalid)",
			source)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	title := strings.TrimSpace(updateCourseRequest.Title)
	if title == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin err (title invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(title) > dbmodel.CourseTitleLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (title: %s) err (title length invalid)",
			title)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	tagsMap := make(map[string]bool)

	tags := updateCourseRequest.Tags

	for index, tag := range tags {
		desTag := strings.TrimSpace(tag)
		tags[index] = desTag

		if desTag == "" {
			errMsg := tlog.E(ctx).Msgf("Handle update course admin err (tag invalid)")

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if len(desTag) > dbmodel.CourseTagNameLenLimit {
			errMsg := tlog.E(ctx).Msgf("Handle update course admin (tag: %s) err (tag length invalid)",
				tag)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		_, ok := tagsMap[desTag]
		if ok {
			errMsg := tlog.E(ctx).Msgf("Handle update course admin (tag: %s) err (tag duplicate)",
				desTag)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		tagsMap[desTag] = true
	}

	abstract := strings.TrimSpace(updateCourseRequest.Abstract)

	if len(abstract) > dbmodel.CourseAbstractLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (abstract: %s) err (abstract length invalid)",
			abstract)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	coverUrl := strings.TrimSpace(updateCourseRequest.CoverUrl)

	if len(coverUrl) > dbmodel.CourseCoverUrlLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (cover url: %s) err (cover url length invalid)",
			coverUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	linkUrl := strings.TrimSpace(updateCourseRequest.LinkUrl)

	if len(linkUrl) > dbmodel.CourseLinkUrlLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (link url: %s) err (link url length invalid)",
			linkUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	detail := strings.TrimSpace(updateCourseRequest.Detail)

	if len(detail) > dbmodel.CourseDetailLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (detail: %s) err (detail length invalid)",
			detail)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	summary := strings.TrimSpace(updateCourseRequest.Summary)

	if len(summary) > dbmodel.CourseSummaryLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (summary: %s) err (summary length invalid)",
			summary)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	objective := strings.TrimSpace(updateCourseRequest.Objective)

	if len(objective) > dbmodel.CourseObjectiveLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (objective: %s) err (objective length invalid)",
			objective)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	outline := strings.TrimSpace(updateCourseRequest.Outline)

	if len(outline) > dbmodel.CourseOutlineLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (outline: %s) err (outline length invalid)",
			outline)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	references := strings.TrimSpace(updateCourseRequest.References)

	if len(references) > dbmodel.CourseReferencesLenLimit {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (references: %s) err (references length invalid)",
			references)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	var publishAt *time.Time

	srcPublishAt := updateCourseRequest.PublishAt
	if srcPublishAt != "" {
		desPublishAt, err := time.ParseInLocation(time.RFC3339, srcPublishAt, time.Local)
		if err != nil {
			errMsg := tlog.E(ctx).Msgf("Handle update course admin (publish at: %s) err (publish at invalid)",
				srcPublishAt)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		publishAt = &desPublishAt
	}

	status := updateCourseRequest.Status
	_, ok := dbmodel.CourseStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle update course admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.UpdateCourseAdmin(ctx, userId, courseId, categoryId, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle update course admin (user id: %s, course id: %s, category id: %s, author: %s, source: %s, title: %s, tags: %v, abstract: %s, cover url: %s, link url: %s, detail: %s, summary: %s, objective: %s, outline: %s, references: %s, publish at: %s, status: %d) err (update course admin %v)",
			userId, courseId, categoryId, author, source, title, tags, abstract, coverUrl, linkUrl, detail, summary, objective, outline, references, publishAt, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}

func HandleDeleteCourseAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	courseId := strings.TrimSpace(c.Param("id"))
	if courseId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle delete course admin err (course id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.DeleteCourseAdmin(ctx, userId, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle delete course admin (user id: %s, course id: %s) err (delete course admin %v)",
			userId, courseId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}

func HandleListCourseCatalogsAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	courseId := strings.TrimSpace(c.Param("id"))
	if courseId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle list course catalogs admin err (course id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	listCourseCatalogsRespData, errx := service.ListCourseCatalogsAdmin(ctx, userId, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list course catalogs admin (user id: %s, course id: %s) err (list course catalogs admin %v)",
			userId, courseId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, listCourseCatalogsRespData)
}

func HandleCreateCourseCatalogAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	courseId := strings.TrimSpace(c.Param("id"))
	if courseId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin err (course id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	createCourseCatalogRequest := &data.CreateCourseCatalogAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, createCourseCatalogRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle create course catalog admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	parentId := strings.TrimSpace(createCourseCatalogRequest.ParentId)

	name := strings.TrimSpace(createCourseCatalogRequest.Name)
	if name == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin err (name invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(name) > dbmodel.CourseCatalogNameLen {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin (name: %s) err (name len limit)",
			name)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	weight := createCourseCatalogRequest.Weight

	status := createCourseCatalogRequest.Status
	_, ok := dbmodel.CourseCatalogStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if createCourseCatalogRequest.Video == nil {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin err (video invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	video := createCourseCatalogRequest.Video

	videoUrl := strings.TrimSpace(video.VideoUrl)
	if videoUrl == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin err (video url invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(videoUrl) > dbmodel.CourseVideoUrlLen {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin (video url: %s) err (video url len limit)",
			videoUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	format := strings.TrimSpace(video.Format)
	if format == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin err (format invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(format) > dbmodel.CourseVideoFormatLen {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin (format: %s) err (format len limit)",
			format)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	language := strings.TrimSpace(video.Language)
	if language == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin err (language invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(language) > dbmodel.CourseVideoLanguageLen {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin (language: %s) err (language len limit)",
			language)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	size := strings.TrimSpace(video.Size)
	if size == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin err (size invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(size) > dbmodel.CourseVideoSizeLen {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin (size: %s) err (size len limit)",
			size)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	duration := strings.TrimSpace(video.Duration)
	if duration == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin err (duration invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(duration) > dbmodel.CourseVideoDurationLen {
		errMsg := tlog.E(ctx).Msgf("Handle create course catalog admin (duration: %s) err (duration len limit)",
			duration)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	srcUploadAt := strings.TrimSpace(video.UploadAt)

	uploadAt, err := time.ParseInLocation(time.RFC3339, srcUploadAt, time.Local)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle create course catalog admin (upload at: %s) err (upload at parse %v)",
			srcUploadAt, err)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	createCourseCatalogRespData, errx := service.CreateCourseCatalogAdmin(ctx, userId, courseId, parentId, name, weight, status, videoUrl, format, language, size, duration, uploadAt)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle create course catalog admin (user id: %s, course id: %s, parent id: %s, name: %s, weight: %d, status: %d, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %v) err (create course catalog admin %v)",
			userId, courseId, parentId, name, weight, status, videoUrl, format, language, size, duration, uploadAt, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, createCourseCatalogRespData)
}

func HandleUpdateCourseCatalogAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	catalogId := strings.TrimSpace(c.Param("id"))
	if catalogId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin err (catalog id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	updateCourseCatalogRequest := &data.UpdateCourseCatalogAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, updateCourseCatalogRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle update course catalog admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	parentId := strings.TrimSpace(updateCourseCatalogRequest.ParentId)

	name := strings.TrimSpace(updateCourseCatalogRequest.Name)
	if name == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin err (name invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(name) > dbmodel.CourseCatalogNameLen {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin (name: %s) err (name len limit)",
			name)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	weight := updateCourseCatalogRequest.Weight

	status := updateCourseCatalogRequest.Status
	_, ok := dbmodel.CourseCatalogStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if updateCourseCatalogRequest.Video == nil {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin err (video invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	video := updateCourseCatalogRequest.Video

	videoUrl := strings.TrimSpace(video.VideoUrl)
	if videoUrl == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin err (video url invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(videoUrl) > dbmodel.CourseVideoUrlLen {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin (video url: %s) err (video url len limit)",
			videoUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	format := strings.TrimSpace(video.Format)
	if format == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin err (format invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(format) > dbmodel.CourseVideoFormatLen {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin (format: %s) err (format len limit)",
			format)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	language := strings.TrimSpace(video.Language)
	if language == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin err (language invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(language) > dbmodel.CourseVideoLanguageLen {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin (language: %s) err (language len limit)",
			language)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	size := strings.TrimSpace(video.Size)
	if size == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin err (size invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(size) > dbmodel.CourseVideoSizeLen {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin (size: %s) err (size len limit)",
			size)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	duration := strings.TrimSpace(video.Duration)
	if duration == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin err (duration invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(duration) > dbmodel.CourseVideoDurationLen {
		errMsg := tlog.E(ctx).Msgf("Handle update course catalog admin (duration: %s) err (duration len limit)",
			duration)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	srcUploadAt := strings.TrimSpace(video.UploadAt)

	uploadAt, err := time.ParseInLocation(time.RFC3339, srcUploadAt, time.Local)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle update course catalog admin (upload at: %s) err (upload at parse %v)",
			srcUploadAt, err)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.UpdateCourseCatalogAdmin(ctx, userId, catalogId, parentId, name, weight, status, videoUrl, format, language, size, duration, uploadAt)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle update course catalog admin (user id: %s, catalog id: %s, parent id: %s, name: %s, weight: %d, status: %d, video url: %s, format: %s, language: %s, size: %s, duration: %s, upload at: %v) err (update course catalog admin %v)",
			userId, catalogId, parentId, name, weight, status, videoUrl, format, language, size, duration, uploadAt, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}

func HandleDeleteCourseCatalogAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	catalogId := strings.TrimSpace(c.Param("id"))
	if catalogId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle delete course catalog admin err (catalog id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.DeleteCourseCatalogAdmin(ctx, userId, catalogId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle delete course catalog admin (user id: %s, catalog id: %s) err (delete course catalog admin %v)",
			userId, catalogId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}
