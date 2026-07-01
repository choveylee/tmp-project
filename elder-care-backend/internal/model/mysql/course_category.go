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
	CourseCategoryNameLen = 32
)

const (
	CourseCategoryStatusNormal   = 1
	CourseCategoryStatusDisabled = 0
)

var (
	CourseCategoryStatusesMap = map[int]int{
		CourseCategoryStatusNormal:   1,
		CourseCategoryStatusDisabled: 0,
	}
)

type CourseCategory struct {
	Id string

	ModuleId string

	Name string

	Weight int
	Status int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CreateCourseCategory(ctx context.Context, moduleId, name string, weight, status int) (*CourseCategory, *terror.Terror) {
	courseCategoryDB := &CourseCategory{
		Id: tutil.NewOid().String(),

		ModuleId: moduleId,

		Name: name,

		Weight: weight,

		Status: status,
	}

	retGorm := serverClient.DB(ctx, runMode).Create(courseCategoryDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Crete course category (module id: %s, nme: %s, weight: %d, status: %d) err (db create %v)",
			moduleId, name, weight, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return courseCategoryDB, nil
}

func FindCourseCategory(ctx context.Context, categoryId string) (*CourseCategory, *terror.Terror) {
	courseCategoriesDB := make([]*CourseCategory, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", categoryId).Limit(1).Find(&courseCategoriesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course category (category id: %s) err (db find %v)",
			categoryId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(courseCategoriesDB) == 0 {
		return nil, nil
	}

	return courseCategoriesDB[0], nil
}

func FindCourseCategoryByName(ctx context.Context, moduleId, name string) (*CourseCategory, *terror.Terror) {
	courseCategoriesDB := make([]*CourseCategory, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("module_id = ? AND name = ?", moduleId, name).Limit(1).Find(&courseCategoriesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course category by name (module id: %s, name: %s) err (db find %v)",
			moduleId, name, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(courseCategoriesDB) == 0 {
		return nil, nil
	}

	return courseCategoriesDB[0], nil
}

func FindCourseCategories(ctx context.Context, moduleId string, status int, pageNum, pageSize int) (int64, []*CourseCategory, *terror.Terror) {
	query := serverClient.DB(ctx, runMode)

	if moduleId != "" {
		query = query.Where("module_id = ?", moduleId)
	}

	if status != -1 {
		query = query.Where("status = ?", status)
	}

	total := int64(0)

	retGorm := query.Model(&CourseCategory{}).Count(&total)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course categories (module id: %s, status: %d, page num: %d, page size: %d) err (db count %v)",
			moduleId, status, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return 0, nil, errx
	}

	if total == 0 {
		return 0, nil, nil
	}

	courseCategoriesDB := make([]*CourseCategory, 0)

	retGorm = query

	if pageNum != -1 && pageSize != -1 {
		retGorm = retGorm.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	retGorm = retGorm.Order("weight ASC, created_at DESC").Find(&courseCategoriesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course categories (module id: %s, status: %d, page num: %d, page size: %d) err (db find %v)",
			moduleId, status, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return 0, nil, errx
	}

	return total, courseCategoriesDB, nil
}

func FindCourseCategoriesById(ctx context.Context, categoryIds []string) ([]*CourseCategory, *terror.Terror) {
	courseCategoriesDB := make([]*CourseCategory, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id IN (?)", categoryIds).Find(&courseCategoriesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course categories by id (category ids: %v) err (db find %v)",
			categoryIds, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return courseCategoriesDB, nil
}

func UpdateCourseCategory(ctx context.Context, categoryId string, moduleId string, name string, weight, status int) *terror.Terror {
	params := map[string]interface{}{
		"module_id": moduleId,
		"name":      name,

		"weight": weight,

		"status": status,

		"updated_at": time.Now(),
	}

	retGorm := serverClient.DB(ctx, runMode).Model(&CourseCategory{}).Where("id = ?", categoryId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update course category (category id: %s, module id: %s, name: %s, weight: %d, status: %d) err (db update %v)",
			categoryId, moduleId, name, weight, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteCourseCategory(ctx context.Context, categoryId string) *terror.Terror {
	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", categoryId).Delete(&CourseCategory{})
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Delete course category (category id: %s) err (db delete %v)",
			categoryId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}
