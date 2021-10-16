package model

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        uint64         `gorm:"primaryKey:autoIncrement" json:"id"`
	CreatedAt *time.Time     `gorm:"->;<-" json:"-"`
	UpdatedAt *time.Time     `gorm:"->;<-" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"->;<-" json:"-"`
}
