package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

const (
	BannerTitleLen    = 50
	BannerAbstractLen = 100
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
