package service

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"

	"dev.choveylee.top/elder-care-backend/internal/data"
	dbmodel "dev.choveylee.top/elder-care-backend/internal/model/mysql"
)

func ListArticleCategoriesClient(ctx context.Context) (*data.ListArticleCategoriesClientRespData, *terror.Terror) {
	_, articleCategoriesDB, errx := dbmodel.FindArticleCategories(ctx, dbmodel.ArticleCategoryStatusNormal, -1, -1)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List article categories client err (db find article categories %v)",
			errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listArticleCategoriesRespData := &data.ListArticleCategoriesClientRespData{
		Categories: make([]*data.ArticleCategoryClientData, 0),
	}

	for _, articleCategoryDB := range articleCategoriesDB {
		articleCategoryData := &data.ArticleCategoryClientData{
			CategoryId: articleCategoryDB.Id,

			Name: articleCategoryDB.Name,
		}

		listArticleCategoriesRespData.Categories = append(listArticleCategoriesRespData.Categories, articleCategoryData)
	}

	return listArticleCategoriesRespData, nil
}

func ListArticlesClient(ctx context.Context, categoryId string, pageNum, pageSize int) (*data.ListArticlesClientRespData, *terror.Terror) {
	total, articlesDB, errx := dbmodel.FindArticles(ctx, categoryId, dbmodel.ArticleStatusNormal, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List articles client (category id: %s, page num: %d, page size: %d) err (db find articles %v)",
			categoryId, pageNum, pageSize, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listArticlesRespData := &data.ListArticlesClientRespData{
		Articles: make([]*data.ArticleClientData, 0),

		Total: total,
	}

	for _, articleDB := range articlesDB {
		articleData := &data.ArticleClientData{
			ArticleId: articleDB.Id,

			Title: articleDB.Title,

			Abstract: articleDB.Abstract,
			Content:  articleDB.Content,

			CoverUrl: articleDB.CoverUrl,
			LinkUrl:  articleDB.LinkUrl,

			PublishAt: articleDB.PublishAt.Format(time.RFC3339),
		}

		listArticlesRespData.Articles = append(listArticlesRespData.Articles, articleData)
	}

	return listArticlesRespData, nil
}

func GetArticleClient(ctx context.Context, articleId string) (*data.GetArticleClientRespData, *terror.Terror) {
	articleDB, errx := dbmodel.FindArticle(ctx, articleId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get article client (article id: %s) err (db find article %v)",
			articleId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if articleDB == nil {
		return nil, nil
	}

	getArticleRespData := &data.GetArticleClientRespData{
		ArticleId: articleDB.Id,

		Title: articleDB.Title,

		Abstract: articleDB.Abstract,
		Content:  articleDB.Content,

		CoverUrl: articleDB.CoverUrl,
		LinkUrl:  articleDB.LinkUrl,

		PublishAt: articleDB.PublishAt.Format(time.RFC3339),
	}

	return getArticleRespData, nil
}
