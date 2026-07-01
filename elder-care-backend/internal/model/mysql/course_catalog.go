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
	CourseCatalogNameLen = 1024
)

const (
	CourseCatalogStatusNormal   = 0
	CourseCatalogStatusDisabled = 1
)

var (
	CourseCatalogStatusesMap = map[int]bool{
		CourseCatalogStatusNormal:   true,
		CourseCatalogStatusDisabled: true,
	}
)

type CourseCatalog struct {
	Id string

	CourseId string
	ParentId string

	Name string

	Weight int
	Status int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CreateCourseCatalog(ctx context.Context, tx *gorm.DB, courseId, parentId, name string, weight, status int) (*CourseCatalog, *terror.Terror) {
	courseCatalogDB := &CourseCatalog{
		Id: tutil.NewOid().String(),

		CourseId: courseId,
		ParentId: parentId,

		Name: name,

		Weight: weight,
		Status: status,
	}

	retGorm := tx.Create(courseCatalogDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Create course catalog (course id: %s, parent id: %s, name: %s, weight: %d, status: %d) err (db create %v)",
			courseId, parentId, name, weight, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return courseCatalogDB, nil
}

func FindCourseCatalog(ctx context.Context, catalogId string) (*CourseCatalog, *terror.Terror) {
	courseCatalogsDB := make([]*CourseCatalog, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", catalogId).Limit(1).Find(&courseCatalogsDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course catalog (catalog id: %s) err (db find %v)",
			catalogId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(courseCatalogsDB) == 0 {
		return nil, nil
	}

	return courseCatalogsDB[0], nil
}

func FindCourseCatalogs(ctx context.Context, courseId string, status int) ([]*CourseCatalog, *terror.Terror) {
	courseCatalogsDB := make([]*CourseCatalog, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("course_id = ?", courseId)

	if status != -1 {
		retGorm = retGorm.Where("status = ?", status)
	}

	retGorm = retGorm.Order("weight ASC, created_at ASC").Find(&courseCatalogsDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find course catalogs (course id: %s, status: %d) err (db find %v)",
			courseId, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return courseCatalogsDB, nil
}

func UpdateCourseCatalog(ctx context.Context, tx *gorm.DB, catalogId, parentId, name string, weight, status int) *terror.Terror {
	params := map[string]interface{}{
		"parent_id": parentId,
		"name":      name,
		"weight":    weight,
		"status":    status,

		"updated_at": time.Now(),
	}

	retGorm := tx.Model(&CourseCatalog{}).Where("id = ?", catalogId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update course catalog (catalog id: %s, parent id: %s, name: %s, weight: %d, status: %d) err (db update %v)",
			catalogId, parentId, name, weight, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteCourseCatalog(ctx context.Context, tx *gorm.DB, catalogId string) *terror.Terror {
	retGorm := tx.Where("id = ?", catalogId).Delete(&CourseCatalog{})
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Delete course catalog (catalog id: %s) err (db delete %v)",
			catalogId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteCourseCatalogs(ctx context.Context, tx *gorm.DB, courseId string) *terror.Terror {
	retGorm := tx.Where("course_id = ?", courseId).Delete(&CourseCatalog{})
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Delete course catalogs (course id: %s) err (db delete %v)",
			courseId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}
