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

func HandleListArticleCategoriesAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	status := -1

	srcStatus := strings.TrimSpace(c.Query("status"))
	if srcStatus != "" {
		desStatus, err := strconv.Atoi(srcStatus)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list article categories admin (status: %s) err (stconv atoi %v)",
				srcStatus, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		_, ok := dbmodel.ArticleCategoryStatusesMap[desStatus]
		if !ok {
			errMsg := tlog.E(ctx).Msgf("Handle list article categories admin (status: %d) err (status invalid)",
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
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list article categories admin (page num: %s) err (strconv atoi %v)",
				srcPageNum, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageNum <= 0 || desPageNum > constant.MaxPageNum {
			errMsg := tlog.E(ctx).Msgf("Handle list article categories admin (page num: %d) err (page num invalid)",
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
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list article categories admin (page size: %s) err (strconv atoi %v)",
				srcPageSize, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageSize <= 0 || desPageSize > constant.MaxPageSize {
			errMsg := tlog.E(ctx).Msgf("Handle list article categories admin (page size: %d) err (page size invalid)",
				desPageSize)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageSize = desPageSize
	}

	getArticleCategoriesRespData, errx := service.ListArticleCategoriesAdmin(ctx, userId, status, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list article categories admin (user id: %s, status: %d, page num: %d, page size: %d) err (list article categories admin %v)",
			userId, status, pageNum, pageSize, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getArticleCategoriesRespData)
}

func HandleCreateArticleCategoryAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	createArticleCategoryRequest := &data.CreateArticleCategoryAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, createArticleCategoryRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle create article category admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	name := strings.TrimSpace(createArticleCategoryRequest.Name)
	if name == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create article category admin err (name invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(name) > dbmodel.ArticleCategoryNameLen {
		errMsg := tlog.E(ctx).Msgf("Handle create article category admin (name: %s) err (name len limit)",
			name)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	weight := createArticleCategoryRequest.Weight

	status := createArticleCategoryRequest.Status

	_, ok := dbmodel.ArticleCategoryStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle create article category admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	createArticleCategoryRespData, errx := service.CreateArticleCategoryAdmin(ctx, userId, name, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle create article category admin (user id: %s, name: %s, weight: %d, status: %d) err (create article category admin %v)",
			userId, name, weight, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, createArticleCategoryRespData)
}

func HandleGetArticleCategoryAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	categoryId := strings.TrimSpace(c.Param("id"))
	if categoryId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle get article category admin err (category id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	getArticleCategoryRespData, errx := service.GetArticleCategoryAdmin(ctx, userId, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle get article category admin (user id: %s, category id: %s) err (get article category admin %v)",
			userId, categoryId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getArticleCategoryRespData)
}

func HandleUpdateArticleCategoryAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	categoryId := strings.TrimSpace(c.Param("id"))
	if categoryId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update article category admin err (category id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	updateArticleCategoryRequest := &data.UpdateArticleCategoryAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, updateArticleCategoryRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle update article category admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	name := strings.TrimSpace(updateArticleCategoryRequest.Name)
	if name == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update article category admin err (name invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(name) > dbmodel.ArticleCategoryNameLen {
		errMsg := tlog.E(ctx).Msgf("Handle update article category admin (title: %s) err (name len limit)",
			name)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	weight := updateArticleCategoryRequest.Weight

	status := updateArticleCategoryRequest.Status

	_, ok := dbmodel.ArticleCategoryStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle update article category admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.UpdateArticleCategoryAdmin(ctx, userId, categoryId, name, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle update article category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (update article category admin %v)",
			userId, categoryId, name, weight, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}

func HandleDeleteArticleCategoryAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	categoryId := strings.TrimSpace(c.Param("id"))

	errx := service.DeleteArticleCategoryAdmin(ctx, userId, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle delete article category admin (user id: %s, category id: %s) err (delete article category admin %v)",
			userId, categoryId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}

func HandleListArticlesAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	categoryId := strings.TrimSpace(c.Query("category_id"))

	status := -1

	srcStatus := strings.TrimSpace(c.Query("status"))
	if srcStatus != "" {
		desStatus, err := strconv.Atoi(srcStatus)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list articles admin (status: %s) err (stconv atoi %v)",
				srcStatus, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		_, ok := dbmodel.ArticleStatusesMap[desStatus]
		if !ok {
			errMsg := tlog.E(ctx).Msgf("Handle list articles admin (status: %d) err (status invalid)",
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
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list articles admin (page num: %s) err (strconv atoi %v)",
				srcPageNum, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageNum <= 0 || desPageNum > constant.MaxPageNum {
			errMsg := tlog.E(ctx).Msgf("Handle list articles admin (page num: %d) err (page num invalid)",
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
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list articles admin (page size: %s) err (strconv atoi %v)",
				srcPageSize, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageSize <= 0 || desPageSize > constant.MaxPageSize {
			errMsg := tlog.E(ctx).Msgf("Handle list articles admin (page size: %d) err (page size invalid)",
				desPageSize)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageSize = desPageSize
	}

	getArticlesRespData, errx := service.ListArticlesAdmin(ctx, userId, categoryId, status, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list articles admin (user id: %s, category id: %s, status: %d, page num: %d, page size: %d) err (list articles admin %v)",
			userId, categoryId, status, pageNum, pageSize, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getArticlesRespData)
}

func HandleCreateArticleAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	createArticleRequest := &data.CreateArticleAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, createArticleRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle create article admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	categoryId := strings.TrimSpace(createArticleRequest.CategoryId)
	if categoryId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create article admin err (category id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	title := strings.TrimSpace(createArticleRequest.Title)
	if title == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create article admin err (title invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(title) > dbmodel.ArticleTitleLen {
		errMsg := tlog.E(ctx).Msgf("Handle create article admin (title: %s) err (title len limit)",
			title)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	abstract := strings.TrimSpace(createArticleRequest.Abstract)

	if len(abstract) > dbmodel.ArticleAbstractLen {
		errMsg := tlog.E(ctx).Msgf("Handle create article admin (abstract: %s) err (abstract len limit)",
			abstract)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	content := strings.TrimSpace(createArticleRequest.Content)
	if content == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create article admin err (content invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(content) > dbmodel.ArticleContentLen {
		errMsg := tlog.E(ctx).Msgf("Handle create article admin (content: %s) err (content len limit)",
			content)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	coverUrl := strings.TrimSpace(createArticleRequest.CoverUrl)

	if len(coverUrl) > dbmodel.LinkUrlLen {
		errMsg := tlog.E(ctx).Msgf("Handle create article admin (cover url: %s) err (cover url len limit)",
			coverUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	linkUrl := strings.TrimSpace(createArticleRequest.LinkUrl)

	if len(linkUrl) > dbmodel.LinkUrlLen {
		errMsg := tlog.E(ctx).Msgf("Handle create article admin (link url: %s) err (link url len limit)",
			linkUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	publishAt := time.Now()

	srcPublishAt := createArticleRequest.PublishAt
	if srcPublishAt != "" {
		desPublishAt, err := time.ParseInLocation(time.RFC3339, srcPublishAt, time.Local)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle create article admin (publish at: %s) err (time parse %v)",
				srcPublishAt, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		publishAt = desPublishAt
	}

	status := createArticleRequest.Status

	_, ok := dbmodel.ArticleCategoryStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle create article admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	createArticleRespData, errx := service.CreateArticleAdmin(ctx, userId, categoryId, title, abstract, content, coverUrl, linkUrl, publishAt, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle create article admin (user id: %s, category id: %s, title: %s, abstract: %s, content: %s, cover url: %s, link url: %s, publish at: %v, status: %d) err (create article admin %v)",
			userId, categoryId, title, abstract, content, coverUrl, linkUrl, publishAt, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, createArticleRespData)
}

func HandleGetArticleAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	articleId := strings.TrimSpace(c.Param("id"))
	if articleId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle get article admin err (article id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	getArticleRespData, errx := service.GetArticleAdmin(ctx, userId, articleId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle get article admin (user id: %s, article id: %s) err (get article admin %v)",
			userId, articleId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getArticleRespData)
}

func HandleUpdateArticleAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	articleId := strings.TrimSpace(c.Param("id"))
	if articleId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update article admin err (article id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	updateArticleRequest := &data.UpdateArticleAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, updateArticleRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle update article admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	title := strings.TrimSpace(updateArticleRequest.Title)
	if title == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update article admin err (title invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(title) > dbmodel.ArticleTitleLen {
		errMsg := tlog.E(ctx).Msgf("Handle update article admin (title: %s) err (title len limit)",
			title)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	abstract := strings.TrimSpace(updateArticleRequest.Abstract)

	if len(abstract) > dbmodel.ArticleAbstractLen {
		errMsg := tlog.E(ctx).Msgf("Handle update article admin (abstract: %s) err (abstract len limit)",
			abstract)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	content := strings.TrimSpace(updateArticleRequest.Content)
	if content == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update article admin err (content invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(content) > dbmodel.ArticleContentLen {
		errMsg := tlog.E(ctx).Msgf("Handle update article admin (content: %s) err (content len limit)",
			content)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	coverUrl := strings.TrimSpace(updateArticleRequest.CoverUrl)

	if len(coverUrl) > dbmodel.LinkUrlLen {
		errMsg := tlog.E(ctx).Msgf("Handle update article admin (cover url: %s) err (cover url len limit)",
			coverUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	linkUrl := strings.TrimSpace(updateArticleRequest.LinkUrl)

	if len(linkUrl) > dbmodel.LinkUrlLen {
		errMsg := tlog.E(ctx).Msgf("Handle update article admin (link url: %s) err (link url len limit)",
			linkUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	publishAt := time.Now()

	srcPublishAt := updateArticleRequest.PublishAt
	if srcPublishAt != "" {
		desPublishAt, err := time.ParseInLocation(time.RFC3339, srcPublishAt, time.Local)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle update article admin (publish at: %s) err (time parse %v)",
				srcPublishAt, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		publishAt = desPublishAt
	}

	status := updateArticleRequest.Status

	_, ok := dbmodel.ArticleCategoryStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle update article admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.UpdateArticleAdmin(ctx, userId, articleId, title, abstract, content, coverUrl, linkUrl, publishAt, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle update article admin (user id: %s, article id: %s, title: %s, abstract: %s, content: %s, cover url: %s, link url: %s, publish at: %v, status: %d) err (update article admin %v)",
			userId, articleId, title, abstract, content, coverUrl, linkUrl, publishAt, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}

func HandleDeleteArticleAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	articleId := strings.TrimSpace(c.Param("id"))
	if articleId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle delete article admin err (article id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.DeleteArticleAdmin(ctx, userId, articleId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle delete article admin (user id: %s, article id: %s) err (delete article admin %v)",
			userId, articleId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}
