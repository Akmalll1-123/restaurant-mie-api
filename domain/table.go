package domain

import (
	"time"

	"gorm.io/gorm"
)

type Table struct {
	ID uint `gorm:"primaryKey"`

	TableNo string `gorm:"unique;not null"`

	Capacity int

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
