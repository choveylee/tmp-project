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
	RoleCodeAdmin = "ADMIN"
)

type Role struct {
	Id string

	Code string
	Name string

	IsAdmin bool

	Weight int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func FindRole(ctx context.Context, roleId string) (*Role, *terror.Terror) {
	rolesDB := make([]*Role, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", roleId).Limit(1).Find(&rolesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find role (role id: %s) err (db find %v)",
			roleId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(rolesDB) == 0 {
		return nil, nil
	}

	return rolesDB[0], nil
}

func FindRoles(ctx context.Context) ([]*Role, *terror.Terror) {
	rolesDB := make([]*Role, 0)

	retGorm := serverClient.DB(ctx, runMode).Order("weight ASC, created_at DESC").Find(&rolesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find roles err (db find %v)",
			retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return rolesDB, nil
}

func FindRolesById(ctx context.Context, roleIds []string) ([]*Role, *terror.Terror) {
	rolesDB := make([]*Role, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id in (?)", roleIds).Order("weight ASC, created_at DESC").Find(&rolesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find roles by id (role ids: %v) err (db find %v)",
			roleIds, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return rolesDB, nil
}

func FindRolesByCode(ctx context.Context, codes []string) ([]*Role, *terror.Terror) {
	rolesDB := make([]*Role, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("code in (?)", codes).Order("weight ASC, created_at DESC").Find(&rolesDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find roles by code (codes: %v) err (db find %v)",
			codes, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return rolesDB, nil
}
