package service

import (
	"context"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"

	"dev.choveylee.top/elder-care-backend/internal/data"
	dbmodel "dev.choveylee.top/elder-care-backend/internal/model/mysql"
)

func ListBannersClient(ctx context.Context) (*data.ListBannersClientRespData, *terror.Terror) {
	_, bannersDB, errx := dbmodel.FindBanners(ctx, dbmodel.BannerStatusNormal, -1, -1)
	if errx != nil {
		errMsg := tlog.E(ctx).Err(errx).Msgf("List banners client err (db find banners %v)",
			errx)
		errx.AttachErrMsg(errMsg)

		return nil, errx
	}

	listBannersRespData := &data.ListBannersClientRespData{
		Banners: make([]*data.BannerClientData, 0),
	}

	for _, bannerDB := range bannersDB {
		bannerData := &data.BannerClientData{
			BannerId: bannerDB.Id,

			Title: bannerDB.Title,

			Abstract: bannerDB.Abstract,

			ImageUrl: bannerDB.ImageUrl,
			LinkUrl:  bannerDB.LinkUrl,
		}

		listBannersRespData.Banners = append(listBannersRespData.Banners, bannerData)
	}

	return listBannersRespData, nil
}
