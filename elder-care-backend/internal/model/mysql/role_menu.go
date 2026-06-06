package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type RoleMenu struct {
	Id string

	RoleId string
	MenuId string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
