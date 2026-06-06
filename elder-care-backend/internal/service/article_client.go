package service

import (
	"context"

	"github.com/choveylee/terror"

	"dev.choveylee.top/elder-care-backend/internal/data"
)

func ListArticleCategoriesClient(ctx context.Context) (*data.ListArticleCategoriesClientRespData, *terror.Terror) {
	return nil, nil
}

func ListArticlesClient(ctx context.Context, categoryId string) (*data.ListArticlesClientRespData, *terror.Terror) {
	return nil, nil
}

func GetArticleClient(ctx context.Context, articleId string) (*data.ArticleClientData, *terror.Terror) {
	return nil, nil
}
