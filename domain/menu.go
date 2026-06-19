package domain

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID uint `gorm:"primaryKey"`

	Name string `gorm:"not null"`

	Category string `gorm:"not null"`

	Price float64 `gorm:"not null"`

	Stock int `gorm:"not null"`

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
