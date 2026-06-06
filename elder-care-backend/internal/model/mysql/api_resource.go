package dbmodel

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"gorm.io/gorm"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

const (
	ApiResourceStatusNormal = 1
)

type ApiResource struct {
	Id string

	Code string
	Name string

	Method string
	Path   string

	ResourceCode string
	ActionCode   string

	AuthRequired int

	Status  int
	Remarks string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func FindApiResourceByPath(ctx context.Context, method, path string) (*ApiResource, *terror.Terror) {
	apiResourcesDB := make([]*ApiResource, 0)

	retGorm := serverClient.DB(ctx, runMode).
		Where("method = ? AND path = ? AND status = ?", method, path, ApiResourceStatusNormal).
		Limit(1).
		Find(&apiResourcesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find api resource by path (method: %s, path: %s) err (db find %v)",
			method, path, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(apiResourcesDB) == 0 {
		return nil, nil
	}

	return apiResourcesDB[0], nil
}

func FindApiResourceByRole(ctx context.Context, roleId, method, path string) (*ApiResource, *terror.Terror) {
	subQuery := serverClient.DB(ctx, runMode).Model(&RolePermission{}).Select("permission_id").Where("role_id = ?", roleId)

	subQuery2 := serverClient.DB(ctx, runMode).Model(&Permission{}).Select("id").Where("id IN (?) AND status = ?", subQuery, PermissionStatusNormal)

	subQuery3 := serverClient.DB(ctx, runMode).Model(&PermissionApiResource{}).Select("api_resource_id").Where("permission_id IN (?)", subQuery2)

	apiResourcesDB := make([]*ApiResource, 0)

	retGorm := serverClient.DB(ctx, runMode).
		Where("method = ? AND path = ? AND status = ? AND id IN (?)", method, path, ApiResourceStatusNormal, subQuery3).
		Limit(1).
		Find(&apiResourcesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find api resource by role (role id: %s, method: %s, path: %s) err (db find %v)",
			roleId, method, path, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(apiResourcesDB) == 0 {
		return nil, nil
	}

	return apiResourcesDB[0], nil
}
