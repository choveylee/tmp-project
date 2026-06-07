package service

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"golang.org/x/crypto/bcrypt"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/data"
	dbmodel "dev.choveylee.top/elder-care-backend/internal/model/mysql"
)

func ListUsersAdmin(ctx context.Context, userId string, roleId string, status int, pageNum, pageSize int) (*data.ListUsersAdminRespData, *terror.Terror) {
	total, usersDB, errx := dbmodel.FindUsers(ctx, roleId, status, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List users admin (user id: %s, role id: %s, status: %d, page num: %d, page size: %d) err (db find users %v)",
			userId, roleId, status, pageNum, pageSize, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listUsersRespData := &data.ListUsersAdminRespData{
		List: make([]*data.UserAdminData, 0),

		Total: total,
	}

	if len(usersDB) == 0 {
		return listUsersRespData, nil
	}

	roleIdsMap := make(map[string]bool)

	for _, userDB := range usersDB {
		roleIdsMap[userDB.RoleId] = true
	}

	roleIds := make([]string, 0)

	for roleId := range roleIdsMap {
		roleIds = append(roleIds, roleId)
	}

	rolesDB, errx := dbmodel.FindRolesById(ctx, roleIds)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List users admin (user id: %s, role id:%s, role ids: %v) err (db find roles by id %v)",
			userId, roleId, roleIds, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	rolesDBMap := make(map[string]*dbmodel.Role)

	for _, roleDB := range rolesDB {
		rolesDBMap[roleDB.Id] = roleDB
	}

	for _, userDB := range usersDB {
		userData := &data.UserAdminData{
			UserId: userDB.Id,

			RoleId: userDB.RoleId,

			Name:   userDB.Name,
			Mobile: userDB.Mobile,

			Status: userDB.Status,

			CreatedAt: userDB.CreatedAt.Format(time.RFC3339),
			UpdatedAt: userDB.UpdatedAt.Format(time.RFC3339),
		}

		if userDB.LoginAt != nil {
			userData.LoginAt = userDB.LoginAt.Format(time.RFC3339)
		}

		if roleDB, ok := rolesDBMap[userDB.RoleId]; ok {
			userData.RoleName = roleDB.Name
		}

		listUsersRespData.List = append(listUsersRespData.List, userData)
	}

	return listUsersRespData, nil
}

func CreateUserAdmin(ctx context.Context, userId, roleId, name, mobile, password string, status int) (*data.CreateUserAdminRespData, *terror.Terror) {
	roleDB, errx := dbmodel.FindRole(ctx, roleId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create user admin (user id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (db find role %v)",
			userId, roleId, name, mobile, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if roleDB == nil {
		errMsg := tlog.E(ctx).Msgf("Create user admin (user id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (role not exist)",
			userId, roleId, name, mobile, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("role id"), constant.ErrorCodeRoleNotExist, errMsg)

		return nil, errx
	}

	userDB, errx := dbmodel.FindUserByMobile(ctx, mobile)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create user admin (user id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (db find user by mobile %v)",
			userId, roleId, name, mobile, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if userDB != nil {
		errMsg := tlog.E(ctx).Msgf("Create user admin (user id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (mobile exist)",
			userId, roleId, name, mobile, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("mobile"), constant.ErrorCodeUserMobileExist, errMsg)

		return nil, errx
	}

	userDB, errx = dbmodel.CreateUser(ctx, roleId, name, mobile, password, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create user admin (user id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (db create user %v)",
			userId, roleId, name, mobile, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	createUserRespData := &data.CreateUserAdminRespData{
		UserId: userDB.Id,
	}

	return createUserRespData, nil
}

func UpdateUserAdmin(ctx context.Context, userId, tUserId, roleId, name, mobile string, status int) *terror.Terror {
	userDB, errx := dbmodel.FindUser(ctx, tUserId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update user admin (user id: %s, tuser id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (db find user %v)",
			userId, tUserId, roleId, name, mobile, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if userDB == nil {
		errMsg := tlog.E(ctx).Msgf("Update user admin (user id: %s, tuser id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (user not exist)",
			userId, tUserId, roleId, name, mobile, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("user id"), constant.ErrorCodeUserNotExist, errMsg)

		return errx
	}

	if userDB.RoleId != roleId {
		roleDB, errx := dbmodel.FindRole(ctx, roleId)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Update user admin (user id: %s, tuser id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (db find role %v)",
				userId, tUserId, roleId, name, mobile, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		if roleDB == nil {
			errMsg := tlog.E(ctx).Msgf("Update user admin (user id: %s, tuser id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (role not exist)",
				userId, tUserId, roleId, name, mobile, status)

			errx := terror.NewTerror(ctx, terror.ErrParamInvalid("role id"), constant.ErrorCodeRoleNotExist, errMsg)

			return errx
		}
	}

	if userDB.Mobile != mobile {
		userDB, errx := dbmodel.FindUserByMobile(ctx, mobile)
		if errx != nil {
			errMsg := tlog.E(ctx).Err(errx).Msgf("Update user admin (user id: %s, tuser id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (db find user by mobile %v)",
				userId, tUserId, roleId, name, mobile, status, errx)
			errx.AttachErrMsg(errMsg)

			return errx
		}

		if userDB != nil {
			errMsg := tlog.E(ctx).Msgf("Update user admin (user id: %s, tuser id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (mobile exist)",
				userId, tUserId, roleId, name, mobile, status)

			errx := terror.NewTerror(ctx, terror.ErrParamInvalid("mobile"), constant.ErrorCodeUserMobileExist, errMsg)

			return errx
		}
	}

	errx = dbmodel.UpdateUser(ctx, tUserId, roleId, name, mobile, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update user admin (user id: %s, tuser id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (db update user %v)",
			userId, tUserId, roleId, name, mobile, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}

func UpdateUserPasswordAdmin(ctx context.Context, userId, tUserId, password string) *terror.Terror {
	userDB, errx := dbmodel.FindUser(ctx, tUserId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update user password admin (user id: %s, tuser id: %s) err (db find user %v)",
			userId, tUserId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if userDB == nil {
		errMsg := tlog.E(ctx).Msgf("Update user password admin (user id: %s, tuser id: %s) err (user not exist)",
			userId, tUserId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("user id"), constant.ErrorCodeUserNotExist, errMsg)

		return errx
	}

	errx = dbmodel.UpdateUserPassword(ctx, tUserId, password)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update user password admin (user id: %s, tuser id: %s) err (db update user password %v)",
			userId, tUserId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}

func UpdateOwnPasswordAdmin(ctx context.Context, userId, oldPassword, newPassword string) *terror.Terror {
	userDB, errx := dbmodel.FindUser(ctx, userId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update own password admin (user id: %s) err (db find user %v)",
			userId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if userDB == nil {
		errMsg := tlog.E(ctx).Msgf("Update own password admin (user id: %s) err (user not exist)",
			userId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("user id"), constant.ErrorCodeUserNotExist, errMsg)

		return errx
	}

	if userDB.Status != dbmodel.UserStatusNormal {
		errMsg := tlog.E(ctx).Msgf("Update own password admin (user id: %s, status: %d) err (user status invalid)",
			userId, userDB.Status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("user status"), constant.ErrorCodeUserStatusInvalid, errMsg)

		return errx
	}

	err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(oldPassword))
	if err != nil {
		errMsg := tlog.E(ctx).Msgf("Update own password admin (user id: %s) err (old password not match)",
			userId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("old password"), constant.ErrorCodePasswordNotMatch, errMsg)

		return errx
	}

	errx = dbmodel.UpdateUserPassword(ctx, userId, newPassword)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update own password admin (user id: %s) err (db update user password %v)",
			userId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}
