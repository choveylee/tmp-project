package dbmodel

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"gorm.io/gorm"

	constant "dev.choveylee.top/elder-care-backend/internal/const"

const (
	PermissionTypeMenu   = 1
	PermissionTypeButton = 2
	PermissionTypeApi    = 3
	PermissionTypeData   = 4
)

const (
	PermissionStatusDisable = 0
	PermissionStatusNormal  = 1
)

type Permission struct {
	Id string

	Code string
	Name string

	PermissionType int
	GroupId        string

	ResourceCode string
	ActionCode   string

	Weight int
	Status int

	Remarks string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func FindPermissionsByRole(ctx context.Context, roleId string) ([]*Permission, *terror.Terror) {
	subQuery := serverClient.DB(ctx, runMode).Model(&RolePermission{}).Select("permission_id").Where("role_id = ?", roleId)

	permissionsDB := make([]*Permission, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id IN (?) AND status = ?", subQuery, PermissionStatusNormal).Order("weight ASC, created_at DESC").Find(&permissionsDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find permissions by role (role id: %s) err (db find %v)",
			roleId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return permissionsDB, nil
}
