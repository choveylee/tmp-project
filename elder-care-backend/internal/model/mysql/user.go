package dbmodel

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"github.com/choveylee/tutil"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

const (
	UserNameLen     = 255
	UserMobileLen   = 20
	UserPasswordLen = 255
)

const (
	UserStatusDisable = 0
	UserStatusNormal  = 1
)

var (
	UserStatusesMap = map[int]int{
		UserStatusDisable: 0,
		UserStatusNormal:  1,
	}
)

type User struct {
	Id string

	RoleId string

	Name     string
	Mobile   string
	Password string

	Status int

	LoginAt *time.Time

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt
}

func CreateUser(ctx context.Context, roleId, name, mobile, password string, status int) (*User, *terror.Terror) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Create user (role id: %s, name: %s, mobile: %s, status: %d) err (generate from password %v)",
			roleId, name, mobile, status, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return nil, errx
	}

	userDB := &User{
		Id: tutil.NewOid().String(),

		RoleId: roleId,

		Name:     name,
		Mobile:   mobile,
		Password: string(passwordHash),

		Status: status,
	}

	retGorm := serverClient.DB(ctx, runMode).Create(userDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Create user (role id: %s, name: %s, mobile: %s, status: %d) err (db create %v)",
			roleId, name, mobile, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return userDB, nil
}

func FindUser(ctx context.Context, userId string) (*User, *terror.Terror) {
	usersDB := make([]*User, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", userId).Limit(1).Find(&usersDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find user (user id: %s) err (db find %v)",
			userId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(usersDB) == 0 {
		return nil, nil
	}

	return usersDB[0], nil
}

func FindUserByMobile(ctx context.Context, mobile string) (*User, *terror.Terror) {
	usersDB := make([]*User, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("mobile = ?", mobile).Limit(1).Find(&usersDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find user by mobile (mobile: %s) err (db find %v)",
			mobile, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(usersDB) == 0 {
		return nil, nil
	}

	return usersDB[0], nil
}

func FindUsers(ctx context.Context, roleId string, status int, pageNum, pageSize int) (int64, []*User, *terror.Terror) {
	query := serverClient.DB(ctx, runMode).Model(&User{})

	if roleId != "" {
		query = query.Where("role_id = ?", roleId)
	}

	if status != -1 {
		query = query.Where("status = ?", status)
	}

	total := int64(0)

	retGorm := query.Count(&total)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find users (role id: %s, status: %d, page num: %d, page size: %d) err (db count %v)",
			roleId, status, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return -1, nil, errx
	}

	if total == 0 {
		return 0, nil, nil
	}

	usersDB := make([]*User, 0)

	retGorm = query

	if pageNum != -1 && pageSize != -1 {
		retGorm = retGorm.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	retGorm = retGorm.Order("created_at DESC").Find(&usersDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find users (role id: %s, status: %d, page num: %d, page size: %d) err (db find %v)",
			roleId, status, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return -1, nil, errx
	}

	return total, usersDB, nil
}

func FindUsersById(ctx context.Context, userIds []string) ([]*User, *terror.Terror) {
	usersDB := make([]*User, 0)

	retGorm := serverClient.DB(ctx, runMode)

	if len(userIds) != 0 {
		retGorm = retGorm.Where("id IN (?)", userIds)
	}

	retGorm = retGorm.Order("created_at DESC").Find(&usersDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find users by id (user ids: %v) err (db find %v)",
			userIds, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return usersDB, nil
}

func UpdateUser(ctx context.Context, userId, roleId, name, mobile string, status int) *terror.Terror {
	params := map[string]interface{}{
		"role_id": roleId,

		"name":   name,
		"mobile": mobile,

		"status": status,

		"updated_at": time.Now(),
	}

	retGorm := serverClient.DB(ctx, runMode).Model(&User{}).Where("id = ?", userId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update user (user id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (db updates %v)",
			userId, roleId, name, mobile, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func UpdateUserPassword(ctx context.Context, userId, password string) *terror.Terror {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Update user password (user id: %s) err (generate from password %v)",
			userId, err)

		errx := terror.NewTerror(ctx, err, constant.ErrorCodeUnknownServerAbnormal, errMsg)

		return errx
	}

	params := map[string]interface{}{
		"password": string(passwordHash),

		"updated_at": time.Now(),
	}

	retGorm := serverClient.DB(ctx, runMode).Model(&User{}).Where("id = ?", userId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update user password (user id: %s) err (db updates %v)",
			userId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func UpdateUserLoginAt(ctx context.Context, tx *gorm.DB, userId string, loginAt time.Time) *terror.Terror {
	params := map[string]interface{}{
		"login_at": loginAt,

		"updated_at": time.Now(),
	}

	retGorm := tx.Model(&User{}).Where("id = ?", userId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update user login at (user id: %s, login at: %s) err (db updates %v)",
			userId, loginAt.Format(time.RFC3339), retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}
