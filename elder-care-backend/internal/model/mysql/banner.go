package dbmodel

import (
	"context"
	"time"

	"github.com/choveylee/terror"
	"github.com/choveylee/tlog"
	"github.com/choveylee/tutil"
	"gorm.io/gorm"

	constant "dev.choveylee.top/elder-care-backend/internal/const"
)

const (
	BannerTitleLen    = 1024
	BannerAbstractLen = 65535
)

const (
	BannerStatusNormal   = 1
	BannerStatusDisabled = 0
)

var (
	BannerStatusesMap = map[int]int{
		BannerStatusNormal:   1,
		BannerStatusDisabled: 0,
	}
)

type Banner struct {
	Id string

	Title    string
	Abstract string

	ImageUrl string
	LinkUrl  string

	Weight int

	Status int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CreateBanner(ctx context.Context, title, abstract string, imageUrl, linkUrl string, weight, status int) (*Banner, *terror.Terror) {
	bannerDB := &Banner{
		Id: tutil.NewOid().String(),

		Title:    title,
		Abstract: abstract,

		ImageUrl: imageUrl,
		LinkUrl:  linkUrl,

		Weight: weight,
		Status: status,
	}

	retGorm := serverClient.DB(ctx, runMode).Create(bannerDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Create banner (title: %s, abstract: %s, image url: %s, link url: %s, weight: %d, status: %d) err (db create %v)",
			title, abstract, imageUrl, linkUrl, weight, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	return bannerDB, nil
}

func FindBanner(ctx context.Context, bannerId string) (*Banner, *terror.Terror) {
	bannersDB := make([]*Banner, 0)

	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", bannerId).Limit(1).Find(&bannersDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find banner (banner id: %s) err (db find %v)",
			bannerId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return nil, errx
	}

	if len(bannersDB) == 0 {
		return nil, nil
	}

	return bannersDB[0], nil
}

func FindBanners(ctx context.Context, status int, pageNum, pageSize int) (int64, []*Banner, *terror.Terror) {
	query := serverClient.DB(ctx, runMode)

	if status != -1 {
		query = query.Where("status = ?", status)
	}

	total := int64(0)

	retGorm := query.Model(&Banner{}).Count(&total)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find banners (status: %d, page num: %d, page size: %d) err (db count %v)",
			status, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return -1, nil, errx
	}

	if total == 0 {
		return 0, nil, nil
	}

	bannersDB := make([]*Banner, 0)

	retGorm = query

	if pageNum != -1 && pageSize != -1 {
		retGorm = retGorm.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	retGorm = query.Order("weight ASC, created_at DESC").Find(&bannersDB)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Find banners (status: %d, page num: %d, page size: %d) err (db find %v)",
			status, pageNum, pageSize, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return -1, nil, errx
	}

	return total, bannersDB, nil
}

func UpdateBanner(ctx context.Context, bannerId string, title, abstract string, imageUrl, linkUrl string, weight, status int) *terror.Terror {
	params := map[string]interface{}{
		"title":    title,
		"abstract": abstract,

		"image_url": imageUrl,
		"link_url":  linkUrl,

		"weight": weight,
		"status": status,

		"updated_at": time.Now(),
	}

	retGorm := serverClient.DB(ctx, runMode).Model(&Banner{}).Where("id = ?", bannerId).Updates(params)
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Update banner (banner id: %s, title: %s, abstract: %s, image url: %s, link url: %s, weight: %d, status: %d) err (db updates %v)",
			bannerId, title, abstract, imageUrl, linkUrl, weight, status, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}

func DeleteBanner(ctx context.Context, bannerId string) *terror.Terror {
	retGorm := serverClient.DB(ctx, runMode).Where("id = ?", bannerId).Delete(&Banner{})
	if retGorm.Error != nil {
		errMsg := tlog.E(ctx).Err(retGorm.Error).Msgf("Delete banner (banner id: %s) err (db delete %v)",
			bannerId, retGorm.Error)

		errx := terror.NewTerror(ctx, retGorm.Error, constant.ErrorCodeMysqlServerAbnormal, errMsg)

		return errx
	}

	return nil
}
