package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type CourseFavourite struct {
	Id string

	CourseId string
	UserId   string

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
