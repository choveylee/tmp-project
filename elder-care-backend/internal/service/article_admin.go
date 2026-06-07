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

func ListArticleCategoriesAdmin(ctx context.Context, userId string, status int, pageNum, pageSize int) (*data.ListArticleCategoriesAdminRespData, *terror.Terror) {
	total, articleCategoriesDB, errx := dbmodel.FindArticleCategories(ctx, status, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List article categories admin (user id: %s, status: %d, page num: %d, page size: %d) err (db find article categories %v)",
			userId, status, pageNum, pageSize, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listArticleCategoriesRespData := &data.ListArticleCategoriesAdminRespData{
		Categories: make([]*data.ArticleCategoryAdminData, 0),

		Total: total,
	}

	for _, articleCategoryDB := range articleCategoriesDB {
		articleCategoryData := &data.ArticleCategoryAdminData{
			CategoryId: articleCategoryDB.Id,

			Name: articleCategoryDB.Name,

			Weight: articleCategoryDB.Weight,

			Status: articleCategoryDB.Status,

			CreatedAt: articleCategoryDB.CreatedAt.Format(time.RFC3339),
			UpdatedAt: articleCategoryDB.UpdatedAt.Format(time.RFC3339),
		}

		listArticleCategoriesRespData.Categories = append(listArticleCategoriesRespData.Categories, articleCategoryData)
	}

	return listArticleCategoriesRespData, nil
}

func CreateArticleCategoryAdmin(ctx context.Context, userId string, name string, weight, status int) (*data.CreateArticleCategoryAdminRespData, *terror.Terror) {
	articleCategoryDB, errx := dbmodel.FindArticleCategoryByName(ctx, name)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create article category admin (user id: %s, name: %s, weight: %d, status: %d) err (db find article category by name %v)",
			userId, name, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if articleCategoryDB != nil {
		errMsg := tlog.E(ctx).Msgf("Create article category admin (user id: %s, name: %s, weight: %d, status: %d) err (name exist)",
			userId, name, weight, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("name"), constant.ErrorCodeArticleCategoryNameExist, errMsg)

		return nil, errx
	}

	articleCategoryDB, errx := dbmodel.CreateArticleCategory(ctx, name, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create article category admin (user id: %s, name: %s, weight: %d, status: %d) err (db create article category %v)",
			userId, name, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	createArticleCategoryRespData := &data.CreateArticleCategoryAdminRespData{
		CategoryId: articleCategoryDB.Id,
	}

	return createArticleCategoryRespData, nil
}

func GetArticleCategoryAdmin(ctx context.Context, userId string, categoryId string) (*data.GetArticleCategoryAdminRespData, *terror.Terror) {
	articleCategoryDB, errx := dbmodel.FindArticleCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get article category admin (user id: %s, category id: %s) err (db find article category %v)",
			userId, categoryId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if articleCategoryDB == nil {
		errMsg := tlog.E(ctx).Msgf("Get article category admin (user id: %s, category id: %s) err (article category not exist)",
			userId, categoryId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("category id"), constant.ErrorCodeArticleCategoryNotExist, errMsg)

		return nil, errx
	}

	getArticleCategoryRespData := &data.GetArticleCategoryAdminRespData{
		CategoryId: articleCategoryDB.Id,

		Name: articleCategoryDB.Name,

		Weight: articleCategoryDB.Weight,

		Status: articleCategoryDB.Status,

		CreatedAt: articleCategoryDB.CreatedAt.Format(time.RFC3339),
		UpdatedAt: articleCategoryDB.UpdatedAt.Format(time.RFC3339),
	}

	return getArticleCategoryRespData, nil
}

func UpdateArticleCategoryAdmin(ctx context.Context, userId string, categoryId string, name string, weight, status int) *terror.Terror {
	articleCategoryDB, errx := dbmodel.FindArticleCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update article category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (db find article category %v)",
			userId, categoryId, name, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if articleCategoryDB == nil {
		errMsg := tlog.E(ctx).Msgf("Update article category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (article category not exist)",
			userId, categoryId, name, weight, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("category id"), constant.ErrorCodeArticleCategoryNotExist, errMsg)

		return errx
	}

	if articleCategoryDB.Name != name {
		articleCategoryDB, errx := dbmodel.FindArticleCategoryByName(ctx, name)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Update article category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (db find article category by name %v)",
				userId, categoryId, name, weight, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		if articleCategoryDB != nil {
			errMsg := tlog.E(ctx).Msgf("Update article category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (name exist)",
				userId, categoryId, name, weight, status)

			errx := terror.NewTerror(ctx, terror.ErrParamInvalid("name"), constant.ErrorCodeArticleCategoryNameExist, errMsg)

			return errx
		}
	}

	errx = dbmodel.UpdateArticleCategory(ctx, categoryId, name, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update article category admin (user id: %s, category id: %s, name: %s, weight: %d, status: %d) err (db update article category %v)",
			userId, categoryId, name, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}

func DeleteArticleCategoryAdmin(ctx context.Context, userId string, categoryId string) *terror.Terror {
	articleCategoryDB, errx := dbmodel.FindArticleCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete article category admin (user id: %s, category id: %s) err (db find article category %v)",
			userId, categoryId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if articleCategoryDB == nil {
		errMsg := tlog.E(ctx).Msgf("Delete article category admin (user id: %s, category id: %s) err (article category not exist)",
			userId, categoryId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("category id"), constant.ErrorCodeArticleCategoryNotExist, errMsg)

		return errx
	}

	total, errx := dbmodel.FindArticleCountByCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete article category admin (user id: %s, category id: %s) err (db find article count by category %v)",
			userId, categoryId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if total > 0 {
		errMsg := tlog.E(ctx).Msgf("Delete article category admin (user id: %s, category id: %s) err (article category in use)",
			userId, categoryId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("category id"), constant.ErrorCodeArticleCategoryInUse, errMsg)

		return errx
	}

	errx = dbmodel.DeleteArticleCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete article category admin (user id: %s, category id: %s) err (db delete article category %v)",
			userId, categoryId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}

func ListArticlesAdmin(ctx context.Context, userId string, categoryId string, status int, pageNum, pageSize int) (*data.ListArticlesAdminRespData, *terror.Terror) {
	total, articlesDB, errx := dbmodel.FindArticles(ctx, categoryId, status, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List articles admin (user id: %s, category id: %s, status: %d, page num: %d, page size: %d) err (db find articles %v)",
			userId, categoryId, status, pageNum, pageSize, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listArticlesRespData := &data.ListArticlesAdminRespData{
		Articles: make([]*data.ArticleAdminData, 0),

		Total: total,
	}

	for _, articleDB := range articlesDB {
		articleData := &data.ArticleAdminData{
			ArticleId: articleDB.Id,

			Title: articleDB.Title,

			Abstract: articleDB.Abstract,
			Content:  articleDB.Content,

			CoverUrl: articleDB.CoverUrl,
			LinkUrl:  articleDB.LinkUrl,

			PublishAt: articleDB.PublishAt.Format(time.RFC3339),

			Status: articleDB.Status,

			CreatedAt: articleDB.CreatedAt.Format(time.RFC3339),
			UpdatedAt: articleDB.UpdatedAt.Format(time.RFC3339),
		}

		listArticlesRespData.Articles = append(listArticlesRespData.Articles, articleData)
	}

	return listArticlesRespData, nil
}

func CreateArticleAdmin(ctx context.Context, userId string, categoryId string, title string, abstract, content string, coverUrl, linkUrl string, publishAt time.Time, status int) (*data.CreateArticleAdminRespData, *terror.Terror) {
	articleCategoryDB, errx := dbmodel.FindArticleCategory(ctx, categoryId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create article admin (user id: %s, category id: %s, title: %s, abstract: %s, content: %s, cover url: %s, link url: %s, publish at: %s, status: %d) err (db find article category %v)",
			userId, categoryId, title, abstract, content, coverUrl, linkUrl, publishAt, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if articleCategoryDB == nil {
		errMsg := tlog.E(ctx).Msgf("Create article admin (user id: %s, category id: %s, title: %s, abstract: %s, content: %s, cover url: %s, link url: %s, publish at: %s, status: %d) err (article category not exist)",
			userId, categoryId, title, abstract, content, coverUrl, linkUrl, publishAt, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("category id"), constant.ErrorCodeArticleCategoryNotExist, errMsg)

		return nil, errx
	}

	articleDB, errx := dbmodel.CreateArticle(ctx, categoryId, title, abstract, content, coverUrl, linkUrl, publishAt, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create article admin (user id: %s, category id: %s, title: %s, abstract: %s, content: %s, cover url: %s, link url: %s, publish at: %s, status: %d) err (db create article %v)",
			userId, categoryId, title, abstract, content, coverUrl, linkUrl, publishAt, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	createArticleRespData := &data.CreateArticleAdminRespData{
		ArticleId: articleDB.Id,
	}

	return createArticleRespData, nil
}

func GetArticleAdmin(ctx context.Context, userId string, articleId string) (*data.GetArticleAdminRespData, *terror.Terror) {
	articleDB, errx := dbmodel.FindArticle(ctx, articleId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get article admin (user id: %s, article id: %s) err (db find article %v)",
			userId, articleId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if articleDB == nil {
		errMsg := tlog.E(ctx).Msgf("Get article admin (user id: %s, article id: %s) err (article not exist)",
			userId, articleId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("article id"), constant.ErrorCodeArticleNotExist, errMsg)

		return nil, errx
	}

	getArticleRespData := &data.GetArticleAdminRespData{
		ArticleId: articleDB.Id,

		Title: articleDB.Title,

		Abstract: articleDB.Abstract,
		Content:  articleDB.Content,

		CoverUrl: articleDB.CoverUrl,
		LinkUrl:  articleDB.LinkUrl,

		PublishAt: articleDB.PublishAt.Format(time.RFC3339),

		Status: articleDB.Status,

		CreatedAt: articleDB.CreatedAt.Format(time.RFC3339),
		UpdatedAt: articleDB.UpdatedAt.Format(time.RFC3339),
	}

	return getArticleRespData, nil
}

func UpdateArticleAdmin(ctx context.Context, userId string, articleId string, title string, abstract, content string, coverUrl, linkUrl string, publishAt time.Time, status int) *terror.Terror {
	articleDB, errx := dbmodel.FindArticle(ctx, articleId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update article admin (user id: %s, article id: %s, title: %s, abstract: %s, content: %s, cover url: %s, link url: %s, publish at: %s, status: %d) err (db find article %v)",
			userId, articleId, title, abstract, content, coverUrl, linkUrl, publishAt, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if articleDB == nil {
		errMsg := tlog.E(ctx).Msgf("Update article admin (user id: %s, article id: %s, title: %s, abstract: %s, content: %s, cover url: %s, link url: %s, publish at: %s, status: %d) err (article not exist)",
			userId, articleId, title, abstract, content, coverUrl, linkUrl, publishAt, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("article id"), constant.ErrorCodeArticleNotExist, errMsg)

		return errx
	}

	errx = dbmodel.UpdateArticle(ctx, articleId, title, abstract, content, coverUrl, linkUrl, publishAt, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update article admin (user id: %s, article id: %s, title: %s, abstract: %s, content: %s, cover url: %s, link url: %s, publish at: %s, status: %d) err (db update article %v)",
			userId, articleId, title, abstract, content, coverUrl, linkUrl, publishAt, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}

func DeleteArticleAdmin(ctx context.Context, userId string, articleId string) *terror.Terror {
	articleDB, errx := dbmodel.FindArticle(ctx, articleId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete article admin (user id: %s, article id: %s) err (db find article %v)",
			userId, articleId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if articleDB == nil {
		errMsg := tlog.E(ctx).Msgf("Delete article admin (user id: %s, article id: %s) err (article not exist)",
			userId, articleId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("article id"), constant.ErrorCodeArticleNotExist, errMsg)

		return errx
	}

	errx = dbmodel.DeleteArticle(ctx, articleId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete article admin (user id: %s, article id: %s) err (db delete article %v)",
			userId, articleId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}
