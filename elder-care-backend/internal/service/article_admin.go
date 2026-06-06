package service

import (
	"context"
	"time"

	"github.com/choveylee/terror"

	"dev.choveylee.top/elder-care-backend/internal/data"
)

func ListArticleCategoriesAdmin(ctx context.Context, userId string, status int, pageNum, pageSize int) (*data.ListArticleCategoriesAdminRespData, *terror.Terror) {
	return nil, nil
}

func CreateArticleCategoryAdmin(ctx context.Context, userId string, name string, weight, status int) (*data.CreateArticleCategoryAdminRespData, *terror.Terror) {
	return nil, nil
}

func GetArticleCategoryAdmin(ctx context.Context, userId string, categoryId string) (*data.GetArticleCategoryAdminRespData, *terror.Terror) {
	return nil, nil
}

func UpdateArticleCategoryAdmin(ctx context.Context, userId string, categoryId string, name string, weight, status int) *terror.Terror {
	return nil
}

func DeleteArticleCategoryAdmin(ctx context.Context, userId string, categoryId string) *terror.Terror {
	return nil
}

func ListArticlesAdmin(ctx context.Context, userId string, categoryId string, status int, pageNum, pageSize int) (*data.ListArticlesAdminRespData, *terror.Terror) {
	return nil, nil
}

func CreateArticleAdmin(ctx context.Context, userId string, categoryId string, title string, abstract, content string, coverUrl, linkUrl string, publishAt time.Time, status int) (*data.CreateArticleAdminRespData, *terror.Terror) {
	return nil, nil
}

func GetArticleAdmin(ctx context.Context, userId string, articleId string) (*data.GetArticleAdminRespData, *terror.Terror) {
	return nil, nil
}

func UpdateArticleAdmin(ctx context.Context, userId string, articleId string, title string, abstract, content string, coverUrl, linkUrl string, publishAt time.Time, status int) *terror.Terror {
	return nil
}

func DeleteArticleAdmin(ctx context.Context, userId string, articleId string) *terror.Terror {
	return nil
}
