package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

const (
	ArticleTitleLen    = 1024
	ArticleAbstractLen = 1024
	ArticleContentLen  = 65535
)

const (
	ArticleStatusNormal   = 1
	ArticleStatusDisabled = 0
)

var (
	ArticleStatusesMap = map[int]int{
		ArticleStatusNormal:   1,
		ArticleStatusDisabled: 0,
	}
)

type Article struct {
	Id string

	CategoryId string

	Title string

	Abstract string
	Content  string

	CoverUrl string
	LinkUrl  string

	PublishAt time.Time

	Status int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
