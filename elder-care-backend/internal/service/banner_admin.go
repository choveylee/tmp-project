package service

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
	"dev.choveylee.top/elder-care-backend/internal/data"
	dbmodel "dev.choveylee.top/elder-care-backend/internal/model/mysql"
)

func ListBannersAdmin(ctx context.Context, userId string, status int, pageNum, pageSize int) (*data.ListBannersAdminRespData, *terror.Terror) {
	total, bannersDB, errx := dbmodel.FindBanners(ctx, status, pageNum, pageSize)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List banners admin (user id: %s, status: %d, page num: %d, page size: %d) err (db find banners %v)",
			userId, status, pageNum, pageSize, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listBannersRespData := &data.ListBannersAdminRespData{
		Banners: make([]*data.BannerAdminData, 0),

		Total: total,
	}

	for _, bannerDB := range bannersDB {
		bannerData := &data.BannerAdminData{
			BannerId: bannerDB.Id,

			Title:    bannerDB.Title,
			Abstract: bannerDB.Abstract,

			ImageUrl: bannerDB.ImageUrl,
			LinkUrl:  bannerDB.LinkUrl,

			Weight: bannerDB.Weight,

			Status: bannerDB.Status,

			CreatedAt: bannerDB.CreatedAt.Format(time.RFC3339),
			UpdatedAt: bannerDB.UpdatedAt.Format(time.RFC3339),
		}

		listBannersRespData.Banners = append(listBannersRespData.Banners, bannerData)
	}

	return listBannersRespData, nil
}

func CreateBannerAdmin(ctx context.Context, userId string, title, abstract string, imageUrl, linkUrl string, weight, status int) (*data.CreateBannerAdminRespData, *terror.Terror) {
	bannerDB, errx := dbmodel.CreateBanner(ctx, title, abstract, imageUrl, linkUrl, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Create banner admin (user id: %s, title: %s, abstract: %s, image url: %s, link url: %s, weight: %d, status: %d) err (db create banner %v)",
			userId, title, abstract, imageUrl, linkUrl, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	createBannerRespData := &data.CreateBannerAdminRespData{
		BannerId: bannerDB.Id,
	}

	return createBannerRespData, nil
}

func GetBannerAdmin(ctx context.Context, userId string, bannerId string) (*data.GetBannerAdminRespData, *terror.Terror) {
	bannerDB, errx := dbmodel.FindBanner(ctx, bannerId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Get banner admin (user id: %s, banner id: %s) err (db find banner %v)",
			userId, bannerId, errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	if bannerDB == nil {
		errMsg := tlog.E(ctx).Msgf("Get banner admin (user id: %s, banner id: %s) err (banner id invalid)",
			userId, bannerId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("banner id"), constant.ErrorCodeBannerNotExist, errMsg)

		return nil, errx
	}

	getBannerRespData := &data.GetBannerAdminRespData{
		BannerId: bannerDB.Id,

		Title:    bannerDB.Title,
		Abstract: bannerDB.Abstract,

		ImageUrl: bannerDB.ImageUrl,
		LinkUrl:  bannerDB.LinkUrl,

		Weight: bannerDB.Weight,

		Status: bannerDB.Status,

		CreatedAt: bannerDB.CreatedAt.Format(time.RFC3339),
		UpdatedAt: bannerDB.UpdatedAt.Format(time.RFC3339),
	}

	return getBannerRespData, nil
}

func UpdateBannerAdmin(ctx context.Context, userId string, bannerId string, title, abstract string, imageUrl, linkUrl string, weight, status int) *terror.Terror {
	bannerDB, errx := dbmodel.FindBanner(ctx, bannerId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update banner admin (user id: %s, banner id: %s, title: %s, abstract: %s, image url: %s, link url: %s, weight: %d, status: %d) err (db find banner %v)",
			userId, bannerId, title, abstract, imageUrl, linkUrl, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if bannerDB == nil {
		errMsg := tlog.E(ctx).Msgf("Update banner admin (user id: %s, banner id: %s, title: %s, abstract: %s, image url: %s, link url: %s, weight: %d, status: %d) err (banner id invalid)",
			userId, bannerId, title, abstract, imageUrl, linkUrl, weight, status)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("banner id"), constant.ErrorCodeBannerNotExist, errMsg)

		return errx
	}

	errx = dbmodel.UpdateBanner(ctx, bannerId, title, abstract, imageUrl, linkUrl, weight, status)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Update banner admin (user id: %s, banner id: %s, title: %s, abstract: %s, image url: %s, link url: %s, weight: %d, status: %d) err (db update banner %v)",
			userId, bannerId, title, abstract, imageUrl, linkUrl, weight, status, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}

func DeleteBannerAdmin(ctx context.Context, userId string, bannerId string) *terror.Terror {
	bannerDB, errx := dbmodel.FindBanner(ctx, bannerId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete banner admin (user id: %s, banner id: %s) err (db find banner %v)",
			userId, bannerId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	if bannerDB == nil {
		errMsg := tlog.E(ctx).Msgf("Delete banner admin (user id: %s, banner id: %s) err (banner id invalid)",
			userId, bannerId)

		errx := terror.NewTerror(ctx, terror.ErrParamInvalid("banner id"), constant.ErrorCodeBannerNotExist, errMsg)

		return errx
	}

	errx = dbmodel.DeleteBanner(ctx, bannerId)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("Delete banner admin (user id: %s, banner id: %s) err (db delete banner %v)",
			userId, bannerId, errx)
		errx.AttachErrMsg(errMsg)

		return errx
	}

	return nil
}
