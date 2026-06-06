package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type PermissionApiResource struct {
	Id string

	PermissionId  string
	ApiResourceId string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
