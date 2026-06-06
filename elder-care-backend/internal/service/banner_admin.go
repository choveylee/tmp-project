package service

import (
	"context"

	"github.com/choveylee/terror"

	"dev.choveylee.top/elder-care-backend/internal/data"
)

func ListBannersAdmin(ctx context.Context, userId string, status int, pageNum, pageSize int) (*data.ListBannersAdminRespData, *terror.Terror) {
	return nil, nil
}

func CreateBannerAdmin(ctx context.Context, userId string, title, abstract string, imageUrl, linkUrl string, weight, status int) (*data.CreateBannerAdminRespData, *terror.Terror) {
	return nil, nil
}

func GetBannerAdmin(ctx context.Context, userId string, bannerId string) (*data.GetBannerAdminRespData, *terror.Terror) {
	return nil, nil
}

func UpdateBannerAdmin(ctx context.Context, userId string, bannerId string, title, abstract string, imageUrl, linkUrl string, weight, status int) *terror.Terror {
	return nil
}

func DeleteBannerAdmin(ctx context.Context, userId string, bannerId string) *terror.Terror {
	return nil
}
