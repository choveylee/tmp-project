package handler

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/choveylee/tlog"
	"github.com/gin-gonic/gin"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/data"
	dbmodel "dev.choveylee.top/elder-care-backend/internal/model/mysql"
	"dev.choveylee.top/elder-care-backend/internal/service"
)

func HandleListUsersAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	roleId := strings.TrimSpace(c.Query("role_id"))

	status := -1

	srcStatus := strings.TrimSpace(c.Query("status"))
	if srcStatus != "" {
		desStatus, err := strconv.Atoi(srcStatus)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list users admin (status: %s) err (strconv atoi %v)",
				srcStatus, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if _, ok := dbmodel.UserStatusesMap[desStatus]; !ok {
			errMsg := tlog.E(ctx).Msgf("Handle list users admin (status: %d) err (status invalid)",
				desStatus)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		status = desStatus
	}

	pageNum := 1

	srcPageNum := strings.TrimSpace(c.Query("page_num"))
	if srcPageNum != "" {
		desPageNum, err := strconv.Atoi(srcPageNum)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list users admin (page num: %s) err (strconv atoi %v)",
				srcPageNum, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageNum <= 0 || desPageNum > constant.MaxPageNum {
			errMsg := tlog.E(ctx).Msgf("Handle list users admin (page num: %d) err (page num invalid)",
				desPageNum)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageNum = desPageNum
	}

	pageSize := 10

	srcPageSize := strings.TrimSpace(c.Query("page_size"))
	if srcPageSize != "" {
		desPageSize, err := strconv.Atoi(srcPageSize)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list users admin (page size: %s) err (strconv atoi %v)",
				srcPageSize, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageSize <= 0 || desPageSize > constant.MaxPageSize {
			errMsg := tlog.E(ctx).Msgf("Handle list users admin (page size: %d) err (page size invalid)",
				desPageSize)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageSize = desPageSize
	}

	listUsersRespData, errx := service.ListUsersAdmin(ctx, userId, roleId, status, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list users admin (user id: %s, role id: %s, status: %d, page num: %d, page size: %d) err (list users admin %v)",
			userId, roleId, status, pageNum, pageSize, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, listUsersRespData)
}

func HandleCreateUserAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	createUserRequest := &data.CreateUserAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, createUserRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle create user admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	roleId := strings.TrimSpace(createUserRequest.RoleId)
	if roleId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create user admin err (role id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	name := strings.TrimSpace(createUserRequest.Name)
	if name == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create user admin err (name invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(name) > dbmodel.UserNameLen {
		errMsg := tlog.E(ctx).Msgf("Handle create user admin (name: %s) err (name len limit)",
			name)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	mobile := strings.TrimSpace(createUserRequest.Mobile)
	if mobile == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create user admin err (mobile invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(mobile) > dbmodel.UserMobileLen {
		errMsg := tlog.E(ctx).Msgf("Handle create user admin (mobile: %s) err (mobile len limit)",
			mobile)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	retMatches := constant.ChnMobileReg.FindAllString(mobile, -1)
	if len(retMatches) == 0 {
		errMsg := tlog.E(ctx).Msgf("Handle create user admin (mobile: %s) err (mobile invalid)",
			mobile)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	password := strings.TrimSpace(createUserRequest.Password)
	if password == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create user admin err (password invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(password) > dbmodel.UserPasswordLen {
		errMsg := tlog.E(ctx).Msgf("Handle create user admin err (password len limit)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	status := createUserRequest.Status
	if _, ok := dbmodel.UserStatusesMap[status]; !ok {
		errMsg := tlog.E(ctx).Msgf("Handle create user admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	createUserRespData, errx := service.CreateUserAdmin(ctx, userId, roleId, name, mobile, password, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle create user admin (user id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (create user admin %v)",
			userId, roleId, name, mobile, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, createUserRespData)
}

func HandleUpdateUserAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	tUserId := strings.TrimSpace(c.Param("id"))

	updateUserRequest := &data.UpdateUserAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, updateUserRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle update user admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	roleId := strings.TrimSpace(updateUserRequest.RoleId)
	if roleId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update user admin err (role id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	name := strings.TrimSpace(updateUserRequest.Name)
	if name == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update user admin err (name invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(name) > dbmodel.UserNameLen {
		errMsg := tlog.E(ctx).Msgf("Handle update user admin (name: %s) err (name len limit)",
			name)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	mobile := strings.TrimSpace(updateUserRequest.Mobile)
	if mobile == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update user admin err (mobile invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(mobile) > dbmodel.UserMobileLen {
		errMsg := tlog.E(ctx).Msgf("Handle update user admin (mobile: %s) err (mobile len limit)",
			mobile)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	retMatches := constant.ChnMobileReg.FindAllString(mobile, -1)
	if len(retMatches) == 0 {
		errMsg := tlog.E(ctx).Msgf("Handle update user admin (mobile: %s) err (mobile invalid)",
			mobile)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	status := updateUserRequest.Status
	if _, ok := dbmodel.UserStatusesMap[status]; !ok {
		errMsg := tlog.E(ctx).Msgf("Handle update user admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.UpdateUserAdmin(ctx, userId, tUserId, roleId, name, mobile, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle update user admin (user id: %s, tuser id: %s, role id: %s, name: %s, mobile: %s, status: %d) err (update user admin %v)",
			userId, tUserId, roleId, name, mobile, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}

func HandleUpdateUserPasswordAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	tUserId := strings.TrimSpace(c.Param("id"))

	updateUserPasswordRequest := &data.UpdateUserPasswordAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, updateUserPasswordRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle update user password admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	password := strings.TrimSpace(updateUserPasswordRequest.Password)
	if password == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update user password admin err (password invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(password) > dbmodel.UserPasswordLen {
		errMsg := tlog.E(ctx).Msgf("Handle update user password admin err (password len limit)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.UpdateUserPasswordAdmin(ctx, userId, tUserId, password)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle update user password admin (user id: %s, tuser id: %s) err (update user password admin %v)",
			userId, tUserId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}

func HandleUpdateOwnPasswordAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	updatePasswordRequest := &data.UpdateOwnPasswordAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, updatePasswordRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle update own password admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	oldPassword := strings.TrimSpace(updatePasswordRequest.OldPassword)
	if oldPassword == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update own password admin err (old password invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	newPassword := strings.TrimSpace(updatePasswordRequest.NewPassword)
	if newPassword == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update own password admin rr (new password invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(newPassword) > dbmodel.UserPasswordLen {
		errMsg := tlog.E(ctx).Msgf("Handle update own password admin err (password len limit)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.UpdateOwnPasswordAdmin(ctx, userId, oldPassword, newPassword)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle update own password admin (user id: %s) err (update own password admin %v)",
			userId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}
