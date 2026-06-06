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
	MenuTypeCatalog     = 0
	MenuTypeMenu        = 1
	MenuTypeHiddenRoute = 2
	MenuTypeButton      = 3
)

const (
	MenuStatusDisable = 0
	MenuStatusNormal  = 1
)

type Menu struct {
	Id string

	ParentId string

	Code string
	Name string

	MenuType int

	RouteName string
	IconName  string

	Weight int
	Status int

	Remarks string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func FindMenusByRole(ctx context.Context, roleId string) ([]*Menu, *terror.Terror) {
	subQuery := serverClient.DB(ctx, runMode).Model(&RoleMenu{}).Select("menu_id").Where("role_id = ?", roleId)

	menusDB := make([]*Menu, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id IN (?) AND status = ?", subQuery, MenuStatusNormal).Order("weight ASC, created_at DESC").Find(&menusDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find menus by role (role id: %s) err (db find %v)",
			roleId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return menusDB, nil
}
