package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type RolePermission struct {
	Id string

	RoleId       string
	PermissionId string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
