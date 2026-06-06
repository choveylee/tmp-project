package handler

import (
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

	listArticlesRespData, errx := service.ListArticlesClient(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list articles client (category id: %s) err (list articles client %v)",
			categoryId, errx)

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
