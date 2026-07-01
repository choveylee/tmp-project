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
	ArticleCategoryNameLen = 32
)

const (
	ArticleCategoryStatusNormal   = 1
	ArticleCategoryStatusDisabled = 0
)

var (
	ArticleCategoryStatusesMap = map[int]int{
		ArticleCategoryStatusNormal:   1,
		ArticleCategoryStatusDisabled: 0,
	}
)

type ArticleCategory struct {
	Id string

	Name string

	Weight int

	Status int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CreateArticleCategory(ctx context.Context, name string, weight, status int) (*ArticleCategory, *terror.Terror) {
	articleCategoryDB := &ArticleCategory{
		Id: tutil.NewOid().String(),

		Name: name,

		Weight: weight,

		Status: status,
	}

	retGorm := serverClient.DB(ctx, runMode).Create(articleCategoryDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Crete article category (nme: %s, weight: %d, status: %d) err (db create %v)",
			name, weight, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return articleCategoryDB, nil
}

func FindArticleCategory(ctx context.Context, categoryId string) (*ArticleCategory, *terror.Terror) {
	articleCategoriesDB := make([]*ArticleCategory, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", categoryId).Limit(1).Find(&articleCategoriesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find article category (category id: %s) err (db find %v)",
			categoryId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(articleCategoriesDB) == 0 {
		return nil, nil
	}

	return articleCategoriesDB[0], nil
}

func FindArticleCategoryByName(ctx context.Context, name string) (*ArticleCategory, *terror.Terror) {
	articleCategoriesDB := make([]*ArticleCategory, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("name = ?", name).Limit(1).Find(&articleCategoriesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find article category by name (name: %s) err (db find %v)",
			name, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(articleCategoriesDB) == 0 {
		return nil, nil
	}

	return articleCategoriesDB[0], nil
}

func FindArticleCategories(ctx context.Context, status int, pageNum, pageSize int) (int64, []*ArticleCategory, *terror.Terror) {
	query := serverClient.DB(ctx, runMode)

	if status != -1 {
		query = query.Where("status = ?", status)
	}

	total := int64(0)

	retGorm := query.Model(&ArticleCategory{}).Count(&total)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find article categories (status: %d, page num: %d, page size: %d) err (db count %v)",
			status, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return 0, nil, errx
	}

	if total == 0 {
		return 0, nil, nil
	}

	articleCategoriesDB := make([]*ArticleCategory, 0)

	retGorm = query

	if pageNum != -1 && pageSize != -1 {
		retGorm = retGorm.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	retGorm = retGorm.Order("weight ASC, created_at DESC").Find(&articleCategoriesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find article categories (status: %d, page num: %d, page size: %d) err (db find %v)",
			status, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return 0, nil, errx
	}

	return total, articleCategoriesDB, nil
}

func UpdateArticleCategory(ctx context.Context, categoryId string, name string, weight, status int) *terror.Terror {
	params := map[string]interface{}{
		"name": name,

		"weight": weight,

		"status": status,

		"updated_at": time.Now(),
	}

	retGorm := serverClient.DB(ctx, runMode).Model(&ArticleCategory{}).Where("id = ?", categoryId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update article category (category id: %s, name: %s, weight: %d, status: %d) err (db update %v)",
			categoryId, name, weight, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteArticleCategory(ctx context.Context, categoryId string) *terror.Terror {
	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", categoryId).Delete(&ArticleCategory{})
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Delete article category (category id: %s) err (db delete %v)",
			categoryId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}
