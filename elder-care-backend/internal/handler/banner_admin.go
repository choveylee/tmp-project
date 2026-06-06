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

func HandleListBannersAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	status := -1

	srcStatus := strings.TrimSpace(c.Query("status"))
	if srcStatus != "" {
		desStatus, err := strconv.Atoi(srcStatus)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list banners admin (status: %s) err (stconv atoi %v)",
				srcStatus, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		_, ok := dbmodel.BannerStatusesMap[desStatus]
		if !ok {
			errMsg := tlog.E(ctx).Msgf("Handle list banners admin (status: %d) err (status invalid)",
				desStatus)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		status = desStatus
	}

	pageNum := -1

	srcPageNum := strings.TrimSpace(c.Query("page_num"))
	if srcPageNum != "" {
		desPageNum, err := strconv.Atoi(srcPageNum)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list banners admin (page num: %s) err (strconv atoi %v)",
				srcPageNum, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageNum <= 0 || desPageNum > constant.MaxPageNum {
			errMsg := tlog.E(ctx).Msgf("Handle list banners admin (page num: %d) err (page num invalid)",
				desPageNum)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageNum = desPageNum
	}

	pageSize := -1

	srcPageSize := strings.TrimSpace(c.Query("page_size"))
	if srcPageSize != "" {
		desPageSize, err := strconv.Atoi(srcPageSize)
		if err != nil {
			errMsg := tlog.E(ctx).Err(err).Msgf("Handle list banners admin (page size: %s) err (strconv atoi %v)",
				srcPageSize, err)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		if desPageSize <= 0 || desPageSize > constant.MaxPageSize {
			errMsg := tlog.E(ctx).Msgf("Handle list banners admin (page size: %d) err (page size invalid)",
				desPageSize)

			SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

			return
		}

		pageSize = desPageSize
	}

	getBannersRespData, errx := service.ListBannersAdmin(ctx, userId, status, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle list banners admin (user id: %s, status: %d, page num: %d, page size: %d) err (list banners admin %v)",
			userId, status, pageNum, pageSize, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getBannersRespData)
}

func HandleCreateBannerAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	createBannerRequest := &data.CreateBannerAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, createBannerRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle create banner admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	title := strings.TrimSpace(createBannerRequest.Title)
	if title == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create banner admin err (title invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(title) > dbmodel.BannerTitleLen {
		errMsg := tlog.E(ctx).Msgf("Handle create banner admin (title: %s) err (title len limit)",
			title)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	abstract := strings.TrimSpace(createBannerRequest.Abstract)

	if len(abstract) > dbmodel.BannerAbstractLen {
		errMsg := tlog.E(ctx).Msgf("Handle create banner admin (abstract: %s) err (abstract len limit)",
			abstract)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	imageUrl := strings.TrimSpace(createBannerRequest.ImageUrl)
	if imageUrl == "" {
		errMsg := tlog.E(ctx).Msgf("Handle create banner admin err (image url invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(imageUrl) > dbmodel.LinkUrlLen {
		errMsg := tlog.E(ctx).Msgf("Handle create banner admin (image url: %s) err (image url len limit)",
			imageUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	linkUrl := strings.TrimSpace(createBannerRequest.LinkUrl)

	if len(linkUrl) > dbmodel.LinkUrlLen {
		errMsg := tlog.E(ctx).Msgf("Handle create banner admin (link url: %s) err (link url len limit)",
			linkUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	weight := createBannerRequest.Weight

	status := createBannerRequest.Status

	_, ok := dbmodel.BannerStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle create banner admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	createBannerRespData, errx := service.CreateBannerAdmin(ctx, userId, title, abstract, imageUrl, linkUrl, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle create banner admin (user id: %s, title: %s, abstract: %s, image url: %s, link url: %s, weight: %d, status: %d) err (create banner admin %v)",
			userId, title, abstract, imageUrl, linkUrl, weight, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, createBannerRespData)
}

func HandleGetBannerAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	bannerId := strings.TrimSpace(c.Param("id"))
	if bannerId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle get banner admin err (banner id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	getBannerRespData, errx := service.GetBannerAdmin(ctx, userId, bannerId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle get banner admin (user id: %s, banner id: %s) err (get banner admin %v)",
			userId, bannerId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, getBannerRespData)
}

func HandleUpdateBannerAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	bannerId := strings.TrimSpace(c.Param("id"))
	if bannerId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update banner admin err (banner id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	updateBannerRequest := &data.UpdateBannerAdminRequest{}

	body, _ := io.ReadAll(c.Request.Body)

	err := json.Unmarshal(body, updateBannerRequest)
	if err != nil {
		errMsg := tlog.E(ctx).Err(err).Msgf("Handle update banner admin (body: %s) err (request body unmarshal %v)",
			string(body), err)

		SendFailResponse(c, constant.ErrorCodeRequestBodyInvalid, errMsg)

		return
	}

	title := strings.TrimSpace(updateBannerRequest.Title)
	if title == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update banner admin err (title invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(title) > dbmodel.BannerTitleLen {
		errMsg := tlog.E(ctx).Msgf("Handle update banner admin (title: %s) err (title len limit)",
			title)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	abstract := strings.TrimSpace(updateBannerRequest.Abstract)

	if len(abstract) > dbmodel.BannerAbstractLen {
		errMsg := tlog.E(ctx).Msgf("Handle update banner admin (abstract: %s) err (abstract len limit)",
			abstract)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	imageUrl := strings.TrimSpace(updateBannerRequest.ImageUrl)
	if imageUrl == "" {
		errMsg := tlog.E(ctx).Msgf("Handle update banner admin err (image url invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	if len(imageUrl) > dbmodel.LinkUrlLen {
		errMsg := tlog.E(ctx).Msgf("Handle update banner admin (image url: %s) err (image url len limit)",
			imageUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	linkUrl := strings.TrimSpace(updateBannerRequest.LinkUrl)

	if len(linkUrl) > dbmodel.LinkUrlLen {
		errMsg := tlog.E(ctx).Msgf("Handle update banner admin (link url: %s) err (link url len limit)",
			linkUrl)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	weight := updateBannerRequest.Weight

	status := updateBannerRequest.Status

	_, ok := dbmodel.BannerStatusesMap[status]
	if !ok {
		errMsg := tlog.E(ctx).Msgf("Handle update banner admin (status: %d) err (status invalid)",
			status)

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.UpdateBannerAdmin(ctx, userId, bannerId, title, abstract, imageUrl, linkUrl, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle update banner admin (user id: %s, banner id: %s, title: %s, abstract: %s, image url: %s, link url: %s, weight: %d, status: %d) err (update banner admin %v)",
			userId, bannerId, title, abstract, imageUrl, linkUrl, weight, status, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}

func HandleDeleteBannerAdmin(c *gin.Context) {
	ctx := c.Request.Context()

	userId := c.Request.Header.Get("user_id")

	bannerId := strings.TrimSpace(c.Param("id"))
	if bannerId == "" {
		errMsg := tlog.E(ctx).Msgf("Handle delete banner admin err (banner id invalid)")

		SendFailResponse(c, constant.ErrorCodeRequestParamInvalid, errMsg)

		return
	}

	errx := service.DeleteBannerAdmin(ctx, userId, bannerId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Handle delete banner admin (user id: %s, banner id: %s) err (delete banner admin %v)",
			userId, bannerId, errx)

		SendFailResponse(c, errx.ErrCode(), errMsg)

		return
	}

	SendPassResponse(c, nil)
}
