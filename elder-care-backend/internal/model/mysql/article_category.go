package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

const (
	ArticleCategoryNameLen = 64
)

const (
	ArticleCategoryStatusNormal   = 1
	ArticleCategoryStatusDisabled = 0
)

var (
	ArticleCategoryStatusesMap = map[int]int{
		ArticleCategoryStatusNormal:   1,
		ArticleCategoryStatusDisabled: 0,
	}
)

type ArticleCategory struct {
	Id string

	Name string

	Weight int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
