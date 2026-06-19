package domain

import (
	"time"

	"gorm.io/gorm"
)

const (
	RoleUser  = "USER"
	RoleKasir = "KASIR"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
