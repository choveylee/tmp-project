package handler

import (
	"strconv"
	"strings"

	"github.com/choveylee/tlog"
	"github.com/gin-gonic/gin"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	dbmodel "dev.choveylee.top/elder-care-backend/internal/model/mysql"
	"dev.choveylee.top/elder-care-backend/internal/service"
)

func HandleListCourseCategoriesClient(c *gin.Context) {
	ctx := c.Request.Context()

	moduleCode := strings.TrimSpace(c.Query("module_code"))
	if moduleCode == "" {
		errMsg := tlog.E(ctx).Msgf("Handle list course categories client err (module code invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	listCourseCategoriesRespData, errx := service.ListCourseCategoriesClient(ctx, moduleCode)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list course categories client (module code: %s) err (list course categories client %v)",
			moduleCode, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, listCourseCategoriesRespData)
}

func HandleListCoursesClient(c *gin.Context) {
	ctx := c.Request.Context()

	categoryId := strings.TrimSpace(c.Query("category_id"))
	if categoryId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle list courses client err (category id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	courseType := -1

	srcCourseType := strings.TrimSpace(c.Query("course_type"))
	if srcCourseType != "" {
		desCourseType, err := strconv.Atoi(srcCourseType)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list courses client (course type: %s) err (strconv atoi %v)",
				srcCourseType, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		_, ok := dbmodel.CourseTypesMap[desCourseType]
		if !ok {
			errMsg := tlog.E(ctx).Msgf("Handle list courses client (course type: %d) err (course type invalid)",
				desCourseType)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		courseType = desCourseType
	}

	sortBy := dbmodel.CourseSortByPublish

	srcSortBy := strings.TrimSpace(c.Query("sort_by"))
	if srcSortBy != "" {
		_, ok := dbmodel.CourseSortBysMap[srcSortBy]
		if !ok {
			errMsg := tlog.E(ctx).Msgf("Handle list courses client (sort by: %s) err (sort by invalid)",
				srcSortBy)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		sortBy = srcSortBy
	}

	pageNum := 1

	srcPageNum := strings.TrimSpace(c.Query("page_num"))
	if srcPageNum != "" {
		desPageNum, err := strconv.Atoi(srcPageNum)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list courses client (page num: %s) err (strconv atoi %v)",
				srcPageNum, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageNum <= 0 || desPageNum > constant.MaxPageNum {
			errMsg := tlog.E(ctx).Msgf("Handle list courses client (page num: %d) err (page num invalid)",
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
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list courses client (page size: %s) err (strconv atoi %v)",
				srcPageSize, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageSize <= 0 || desPageSize > constant.MaxPageSize {
			errMsg := tlog.E(ctx).Msgf("Handle list courses client (page size: %d) err (page size invalid)",
				desPageSize)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageSize = desPageSize
	}

	listCoursesRespData, errx := service.ListCoursesClient(ctx, categoryId, courseType, sortBy, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list courses client (category id: %s, course type: %d, sort by: %s, page num: %d, page size: %d) err (list courses client %v)",
			categoryId, courseType, sortBy, pageNum, pageSize, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, listCoursesRespData)
}

func HandleGetCourseClient(c *gin.Context) {
	ctx := c.Request.Context()

	courseId := strings.TrimSpace(c.Param("id"))
	if courseId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle get course client err (course id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	getCourseRespData, errx := service.GetCourseClient(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle get course client (course id: %s) err (get course client %v)",
			courseId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getCourseRespData)
}

func HandleListCourseCatalogsClient(c *gin.Context) {
	ctx := c.Request.Context()

	courseId := strings.TrimSpace(c.Param("id"))
	if courseId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle list course catalogs client err (course id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	listCourseCatalogsRespData, errx := service.ListCourseCatalogsClient(ctx, courseId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list course catalogs client (course id: %s) err (list course catalogs client %v)",
			courseId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, listCourseCatalogsRespData)
}
