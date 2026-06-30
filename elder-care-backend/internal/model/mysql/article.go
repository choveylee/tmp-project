package dbmodel

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"github.com/choveylee/tutil"
	"gorm.io/gorm"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

const (
	ArticleTitleLen    = 1024
	ArticleAbstractLen = 1024
	ArticleContentLen  = 65535
)

const (
	ArticleStatusNormal   = 1
	ArticleStatusDisabled = 0
)

var (
	ArticleStatusesMap = map[int]int{
		ArticleStatusNormal:   1,
		ArticleStatusDisabled: 0,
	}
)

type Article struct {
	Id string

	CategoryId string

	Title string

	Abstract string
	Content  string

	CoverUrl string
	LinkUrl  string

	PublishAt time.Time

	Status int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CreateArticle(ctx context.Context, categoryId string, title string, abstract, content string, coverUrl, linkUrl string, publishAt time.Time, status int) (*Article, *terror.Terror) {
	articleDB := &Article{
		Id: tutil.NewOid().String(),

		CategoryId: categoryId,

		Title: title,

		Abstract: abstract,
		Content:  content,

		CoverUrl: coverUrl,
		LinkUrl:  linkUrl,

		PublishAt: publishAt,

		Status: status,
	}

	retGorm := serverClient.DB(ctx, runMode).Create(articleDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Create article (category id: %s, title: %s, abstract: %s, content: %s, cover url: %s, link url:%s, publish at: %v, status: %d) err (db delete %v)",
			categoryId, title, abstract, content, coverUrl, linkUrl, publishAt, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return articleDB, nil
}

func FindArticle(ctx context.Context, articleId string) (*Article, *terror.Terror) {
	articlesDB := make([]*Article, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", articleId).Limit(1).Find(&articlesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find article (article id: %s) err (db find %v)",
			articleId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(articlesDB) == 0. {
		return nil, nil
	}

	return articlesDB[0], nil
}

func FindArticles(ctx context.Context, categoryId string, status int, pageNum, pageSize int) (int64, []*Article, *terror.Terror) {
	query := serverClient.DB(ctx, runMode)

	if categoryId != "" {
		query = query.Where("category_id = ?", categoryId)
	}

	if status != -1 {
		query = query.Where("status = ?", status)
	}

	total := int64(0)

	retGorm := query.Model(&Article{}).Count(&total)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find articles (category id: %s, status: %d, page num: %d, page size: %d) err (db count %v)",
			categoryId, status, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return -1, nil, errx
	}

	if total == 0 {
		return 0, nil, nil
	}

	articlesDB := make([]*Article, 0)

	retGorm = query

	if pageNum != -1 && pageSize != -1 {
		retGorm = retGorm.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	retGorm = query.Order("publish_at DESC, created_at DESC").Find(&articlesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find articles (category id: %s, status: %d, page num: %d, page size: %d) err (db find %v)",
			categoryId, status, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return -1, nil, errx
	}

	return total, articlesDB, nil
}

func FindArticleCountByCategory(ctx context.Context, categoryId string) (int64, *terror.Terror) {
	total := int64(0)

	retGorm := serverClient.DB(ctx, runMode).Model(&Article{}).Where("category_id = ?", categoryId).Count(&total)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find article count by category (category id: %s) err (db count %v)",
			categoryId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return -1, errx
	}

	return total, nil
}

func UpdateArticle(ctx context.Context, articleId string, title string, abstract, content string, coverUrl, linkUrl string, publishAt time.Time, status int) *terror.Terror {
	params := map[string]interface{}{
		"title": title,

		"abstract": abstract,
		"content":  content,

		"cover_url": coverUrl,
		"link_url":  linkUrl,

		"publish_at": publishAt,

		"status": status,

		"updated_at": time.Now(),
	}

	retGorm := serverClient.DB(ctx, runMode).Model(&Article{}).Where("id = ?", articleId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update article (article id: %s, title: %s, abstract: %s, content: %s, cover url: %s, link url:%s, publish at: %v, status: %d) err (db update %v)",
			articleId, title, abstract, content, coverUrl, linkUrl, publishAt, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteArticle(ctx context.Context, articleId string) *terror.Terror {
	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", articleId).Delete(&Article{})
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Delete article (article id: %s) err (db delete %v)",
			articleId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}
