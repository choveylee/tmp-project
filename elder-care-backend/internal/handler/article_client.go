package handler

import (
	"strconv"
	"strings"

	"github.com/choveylee/tlog"
	"github.com/gin-gonic/gin"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/service"
)

func HandleListArticleCategoriesClient(c *gin.Context) {
	ctx := c.Request.Context()

	listArticleCategoriesRespData, errx := service.ListArticleCategoriesClient(ctx)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list article categories client err (list article categories client %v)",
			errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, listArticleCategoriesRespData)
}

func HandleListArticlesClient(c *gin.Context) {
	ctx := c.Request.Context()

	categoryId := strings.TrimSpace(c.Query("category_id"))
	if categoryId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle list articles client err (category id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	pageNum := 1

	srcPageNum := strings.TrimSpace(c.Query("page_num"))
	if srcPageNum != "" {
		desPageNum, err := strconv.Atoi(srcPageNum)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list articles client (page num: %s) err (strconv atoi %v)",
				srcPageNum, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageNum <= 0 || desPageNum > constant.MaxPageNum {
			errMsg := tlog.E(ctx).Msgf("Handle list articles client (page num: %d) err (page num invalid)",
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
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list articles client (page size: %s) err (strconv atoi %v)",
				srcPageSize, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageSize <= 0 || desPageSize > constant.MaxPageSize {
			errMsg := tlog.E(ctx).Msgf("Handle list articles client (page size: %d) err (page size invalid)",
				desPageSize)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageSize = desPageSize
	}

	listArticlesRespData, errx := service.ListArticlesClient(ctx, categoryId, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list articles client (category id: %s, page num: %d, page size: %d) err (list articles client %v)",
			categoryId, pageNum, pageSize, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, listArticlesRespData)
}

func HandleGetArticleClient(c *gin.Context) {
	ctx := c.Request.Context()

	articleId := strings.TrimSpace(c.Param("id"))
	if articleId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle get article client err (article id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	getArticleRespData, errx := service.GetArticleClient(ctx, articleId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle get article client (article id: %s) err (get article client %v)",
			articleId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getArticleRespData)
}
